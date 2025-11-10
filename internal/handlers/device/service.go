package device

import (
	"context"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/uuid"
)

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetDeviceByDeviceUUID(ctx context.Context, deviceUUID string) (status.Code, *Device, error) {
	device, err := s.repo.GetDeviceByDeviceUUID(ctx, deviceUUID)
	if err != nil {
		return status.INTERNAL, nil, err
	}

	if device == nil {
		return status.NOT_FOUND, nil, nil
	}

	return status.SUCCESS, device, nil
}

func (s *Service) CreateDevice(ctx context.Context, req *CreateDeviceReq) (status.Code, error) {
	statusCode, err := ValidationCreateDeviceReq(req)
	if err != nil {
		return statusCode, err
	}

	dto := BuildCreateDeviceDTO(req)
	dto.ID, err = uuid.GenerateUUIDV7()
	if err != nil {
		return status.INTERNAL, err
	}

	err = s.repo.StoreDevice(ctx, nil, dto)
	if err != nil {
		return status.INTERNAL, err
	}

	return status.SUCCESS, nil
}

func (s *Service) UpdateDevice(ctx context.Context, req *UpdateDeviceReq) (status.Code, error) {
	statusCode, err := ValidationUpdateDeviceReq(req)
	if err != nil {
		return statusCode, err
	}

	dto := BuildUpdateDeviceDTO(req)
	err = s.repo.UpdateDevice(ctx, dto)
	if err != nil {
		return status.INTERNAL, err
	}

	return status.SUCCESS, nil
}
