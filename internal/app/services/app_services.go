package appservices

import (
	"fmt"
	"log"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/device"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/health"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/login"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
	"github.com/QuocAnh189/GoCoreFoundation/internal/services/mail"
	"github.com/QuocAnh189/GoCoreFoundation/internal/services/sms"
	"github.com/QuocAnh189/GoCoreFoundation/internal/sessions"
	"github.com/QuocAnh189/GoCoreFoundation/pkg/mailer"
	"github.com/QuocAnh189/GoCoreFoundation/root/jwt"
	rootSession "github.com/QuocAnh189/GoCoreFoundation/root/session"
	"github.com/QuocAnh189/GoCoreFoundation/root/sessionprovider"
)

type ServiceContainer struct {
	// Root resouces
	SessionManager  *sessions.SessionManager
	SessionProvider sessionprovider.SessionProvider
	JwtHelper       jwt.JwtHelper

	HealthService *health.Service
	LoginService  *login.Service
	UserService   *users.Service
	DeviceService *device.Service
}

const (
	sessionTTL = 14 * 24 * time.Hour // 14 days
)

func SetUpAppServices(res *resource.AppResource) (*ServiceContainer, error) {
	log.Println("Initializing services")

	env := res.Env

	log.Println("> jwtHelper...")
	var jwtHelper jwt.JwtHelper
	if env.SharedKeyBytes != nil {
		helper, err := jwt.NewHmacJwtHelper(env.SharedKeyBytes)
		if err != nil {
			panic("failed to create jwt toolkit from env shared key")
		}
		jwtHelper = helper
	} else {
		return nil, fmt.Errorf("unable to determine jwt helper from env")
	}

	log.Println("> sessionManager...")
	sessionManager := sessions.NewSessionManager()

	// Build the session provider
	log.Println("> sessionProvider...")
	var sessionProvider sessionprovider.SessionProvider
	{
		defaultSessFactory := func() rootSession.SessionStorer {
			// Create the basic session that all new sessions are based on
			return sessions.NewSession()
		}
		if env.RootSessionDriver == "xwt" {
			sessionProvider = sessionprovider.NewXwtSessionProvider(
				sessionManager.Container(),
				jwtHelper,
				defaultSessFactory,
				sessionTTL,
			)
		} else {
			sessionProvider = sessionprovider.NewJwtSessionProvider(
				sessionManager.Container(),
				jwtHelper,
				defaultSessFactory,
				sessionTTL,
			)
		}
	}

	log.Println("> smsSvc...")
	var smsSvc *sms.Service
	{
		twilioEnv := res.Env.TwilioConfig
		var smsProvider sms.Provider
		if twilioEnv.AccountSID != nil && twilioEnv.AuthToken != nil && twilioEnv.FromPhoneNumber != nil && twilioEnv.MessagingServiceSID != nil {
			smsProvider = sms.NewTwilioSmsProvider(
				*twilioEnv.AccountSID,
				*twilioEnv.AuthToken,
				*twilioEnv.FromPhoneNumber,
				*twilioEnv.MessagingServiceSID,
			)
		} else {
			log.Println("Twilio SMS config not fully provided; using NoOp SMS provider")
		}

		smsSvc = sms.NewService(smsProvider)
	}

	log.Println("> mailerSvc...")
	var mailSvc *mail.Service
	{
		mailerEnv := res.Env.MailerConfig
		var mailClient *mailer.Client
		if mailerEnv.SMTPHost != nil && mailerEnv.SMTPPort != nil && mailerEnv.Username != nil && mailerEnv.Password != nil && mailerEnv.FromName != nil {
			mailClient = mailer.NewClient(
				*mailerEnv.SMTPHost,
				*mailerEnv.SMTPPort,
				*mailerEnv.Username,
				*mailerEnv.Password,
				*mailerEnv.FromName,
			)
		} else {
			log.Println("Mailer config not fully provided; mailer service will not be initialized")
		}

		mailSvc = mail.NewService(mailClient)
	}

	log.Println("> healthSvc...")
	var healthSvc = health.NewService(smsSvc, mailSvc)

	log.Println("> loginSvc...")
	loginRepo := login.NewRepository(res.Db)
	var loginSvc = login.NewService(loginRepo)

	log.Println("> userSvc...")
	userRepo := users.NewRepository(res.Db)
	var userSvc = users.NewService(userRepo)

	log.Println("> deviceSvc...")
	deviceRepo := device.NewRepository(res.Db)
	var deviceSvc = device.NewService(deviceRepo)

	svcs := ServiceContainer{
		SessionManager:  sessionManager,
		JwtHelper:       jwtHelper,
		SessionProvider: sessionProvider,

		UserService:   userSvc,
		LoginService:  loginSvc,
		HealthService: healthSvc,
		DeviceService: deviceSvc,
	}

	return &svcs, nil
}
