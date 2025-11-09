package middleware

import (
	"fmt"
	"net/http"
	"net/url"

	"context"

	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/device"
	"github.com/QuocAnh189/GoCoreFoundation/internal/sessions"
	appctx "github.com/QuocAnh189/GoCoreFoundation/internal/utils/context"
)

type requestDeviceDetails struct {
	deviceUUID  string
	deviceName  string
	deviceToken string
}

// deviceMiddleware handles device management for authenticated sessions
// A device entry should ONLY be created in the following cases:
// 1. User registration
// 2. Successful user login on a new device
type deviceMiddleware struct {
	sessionSvc *sessions.SessionManager
	deviceSvc  device.Service
}

func DeviceMiddleware(sessionSvc *sessions.SessionManager, deviceSvc device.Service) func(http.Handler) http.Handler {
	mw := &deviceMiddleware{
		sessionSvc: sessionSvc,
		deviceSvc:  deviceSvc,
	}
	return mw.Handler
}

// Handler returns the middleware handler function
func (m *deviceMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := appctx.GetLogger(r.Context())

		sess := sessions.GetRequestSession(r)
		if sess == nil {
			next.ServeHTTP(w, r)
			return
		}

		if m.isSessionAnonymous(sess) {
			logger.Info("sessionDeviceMiddleware: session is anonymous, skipping device middleware...")
			next.ServeHTTP(w, r)
			return
		}

		headerDevice := m.extractDeviceDetails(r)
		if headerDevice.deviceUUID == "" {
			logger.Info("sessionDeviceMiddleware: missing device uuid, skipping device middleware...")
			next.ServeHTTP(w, r)
			return
		}

		println("sessionDeviceMiddleware: processing device uuid:", headerDevice.deviceName)

		userDevice, err := m.ensureDevice(r.Context(), headerDevice)
		if err != nil {
			logger.Info("sessionDeviceMiddleware: ensure device error: %v", err)
			next.ServeHTTP(w, r)
			return
		}

		m.persistDeviceToSession(sess, userDevice)

		next.ServeHTTP(w, r)
	})
}

func (m *deviceMiddleware) extractDeviceDetails(r *http.Request) requestDeviceDetails {
	logger := appctx.GetLogger(r.Context())
	rawDeviceName := r.Header.Get("Device-Name")

	decodedDeviceName, err := url.QueryUnescape(rawDeviceName)
	if err != nil {
		logger.Info("Could not decode device name: %v", err)
		decodedDeviceName = rawDeviceName
	}

	return requestDeviceDetails{
		deviceUUID:  r.Header.Get("Device-UUID"),
		deviceName:  decodedDeviceName,
		deviceToken: r.Header.Get("Device-Push-Token"),
	}
}

func (m *deviceMiddleware) isSessionAnonymous(sess *sessions.AppSession) bool {
	isSecureRaw, ok := sess.Get("is_secure")
	if !ok {
		return true
	}

	isSecure, ok := isSecureRaw.(bool)
	if !ok {
		return true
	}

	return !isSecure
}

func (m *deviceMiddleware) ensureDevice(ctx context.Context, details requestDeviceDetails) (*device.Device, error) {
	logger := appctx.GetLogger(ctx)

	_, userDevice, err := m.deviceSvc.GetDeviceByDeviceUUID(ctx, details.deviceUUID)
	if err != nil {
		return nil, fmt.Errorf("get device by uuid: %w", err)
	}

	if userDevice == nil || userDevice.ID == "" {
		logger.Info("* creating device")
		return m.createDevice(ctx, details)
	}

	logger.Info("* updating device")
	return m.updateDevice(ctx, userDevice.ID, details)
}

func (m *deviceMiddleware) createDevice(ctx context.Context, details requestDeviceDetails) (*device.Device, error) {
	createDeviceReq := device.CreateDeviceReq{
		DeviceUUID:      details.deviceUUID,
		DeviceName:      details.deviceName,
		DevicePushToken: &details.deviceToken,
	}

	_, err := m.deviceSvc.CreateDevice(ctx, &createDeviceReq)
	if err != nil {
		return nil, fmt.Errorf("create device: %w", err)
	}

	_, deviceCreated, err := m.deviceSvc.GetDeviceByDeviceUUID(ctx, details.deviceUUID)
	if err != nil {
		return nil, fmt.Errorf("get device by id: %w", err)
	}

	return deviceCreated, nil
}

func (m *deviceMiddleware) updateDevice(ctx context.Context, deviceID string, details requestDeviceDetails) (*device.Device, error) {
	updateDeviceReq := device.UpdateDeviceReq{
		DeviceUUID:      &deviceID,
		DeviceName:      &details.deviceName,
		DevicePushToken: &details.deviceToken,
	}

	_, err := m.deviceSvc.UpdateDevice(ctx, &updateDeviceReq)
	if err != nil {
		return nil, fmt.Errorf("modify device: %w", err)
	}

	_, deviceUpdated, err := m.deviceSvc.GetDeviceByDeviceUUID(ctx, details.deviceUUID)
	if err != nil {
		return nil, fmt.Errorf("get device by id: %w", err)
	}

	return deviceUpdated, nil
}

func (m *deviceMiddleware) persistDeviceToSession(sess *sessions.AppSession, userDevice *device.Device) {
	if userDevice == nil || userDevice.ID == "" {
		return
	}

	sess.Put("device", map[string]any{
		"device_id":         userDevice.ID,
		"device_uuid":       userDevice.DeviceUuid,
		"device_name":       userDevice.DeviceName,
		"device_push_token": userDevice.DevicePushToken,
	})
}
