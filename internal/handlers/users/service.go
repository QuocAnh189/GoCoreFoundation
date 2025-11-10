package users

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/device"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/pagination"
)

type Service struct {
	repo       IRepository
	deviceRepo device.IRepository
}

func NewService(repo IRepository, deviceRepo device.IRepository) *Service {
	return &Service{
		repo:       repo,
		deviceRepo: deviceRepo,
	}
}

func (s *Service) ListUsers(ctx context.Context, req *ListUserRequest) (status.Code, []*User, *pagination.Pagination, error) {
	result, err := s.repo.List(ctx, req)
	if err != nil {
		return status.INTERNAL, nil, nil, err
	}
	return status.SUCCESS, result.Users, result.Pagination, nil
}

func (s *Service) GetUserByLoginName(ctx context.Context, loginName string) (status.Code, *User, error) {
	result, err := s.repo.GetUserByLoginName(ctx, loginName)
	if err != nil {
		return status.INTERNAL, nil, err
	}

	return status.SUCCESS, result, nil
}

func (s *Service) GetUserByID(ctx context.Context, id string) (status.Code, *User, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return status.INTERNAL, nil, err
	}
	if result == nil {
		return status.USER_NOT_FOUND, nil, ErrUserNotFound
	}
	return status.SUCCESS, result, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (status.Code, *User, error) {
	result, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return status.INTERNAL, nil, err
	}
	if result == nil {
		return status.USER_NOT_FOUND, nil, ErrUserNotFound
	}
	return status.SUCCESS, result, nil
}

func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (status.Code, *User, error) {
	statusCode, err := ValidateCreateUserRequest(req)
	if err != nil {
		return statusCode, nil, err
	}

	for _, loginName := range []string{req.Email, req.Phone} {
		existingUser, err := s.repo.GetUserByLoginName(ctx, loginName)
		if err != nil {
			return status.INTERNAL, nil, err
		}
		if existingUser != nil {
			switch loginName {
			case req.Email:
				return status.USER_EMAIL_ALREADY_EXISTS, nil, ErrEmailAlreadyExists
			case req.Phone:
				return status.USER_PHONE_ALREADY_EXISTS, nil, ErrPhoneAlreadyExists
			}
		}
	}

	createUserDTO := BuildCreateUserDTO(req)
	handler := func(tx *sql.Tx) error {
		// Create the user
		_, err := s.repo.Create(ctx, tx, createUserDTO)
		if err != nil {
			return fmt.Errorf("failed to create user in transaction: %v", err)
		}

		// Store aliases
		for _, aka := range []string{createUserDTO.Email, createUserDTO.Phone} {
			if aka == "" {
				continue // Skip empty aliases
			}

			createAliasDTO := BuildAliasDTO(createUserDTO.ID, createUserDTO.Email)
			if err := s.repo.StoreUserAlias(ctx, tx, createAliasDTO); err != nil {
				return fmt.Errorf("failed to store user alias in transaction: %v", err)
			}
		}

		// Store login
		hashedPassword, err := utils.DefaultHasher.Hash(createUserDTO.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %v", err)
		}
		createLoginDTO := BuildLoginDTO(createUserDTO.ID, hashedPassword)
		if err := s.repo.StoreLogin(ctx, tx, createLoginDTO); err != nil {
			return fmt.Errorf("failed to store user login in transaction: %v", err)
		}

		// Store device
		createDeviceDto := device.BuildCreateDeviceDTO(&device.CreateDeviceReq{
			UID:        &createUserDTO.ID,
			DeviceUUID: req.DeviceUUID,
			DeviceName: req.DeviceName,
		})
		err = s.deviceRepo.StoreDevice(ctx, tx, createDeviceDto)
		if err != nil {
			return fmt.Errorf("failed to store device info in transaction: %v", err)
		}

		return nil
	}

	result, err := s.repo.CreateUserWithAssociations(ctx, handler, createUserDTO.ID)
	if err != nil {
		return status.INTERNAL, nil, err
	}

	return status.SUCCESS, result, nil
}

func (s *Service) UpdateUser(ctx context.Context, req *UpdateUserRequest) (status.Code, *User, error) {
	statusCode, err := ValidateUpdateUserRequest(req)
	if err != nil {
		return statusCode, nil, err
	}

	dto := BuildUpdateUserDTO(req)
	_, err = s.repo.Update(ctx, dto)
	if err != nil {
		return status.INTERNAL, nil, err
	}

	result, err := s.repo.FindByID(ctx, dto.ID)
	if err != nil {
		return status.INTERNAL, nil, err
	}

	return status.SUCCESS, result, nil
}

func (s *Service) DeleteUser(ctx context.Context, uid string) (status.Code, error) {
	if uid == "" {
		return status.USER_INVALID_ID, ErrInvalidUserID
	}

	handler := func(tx *sql.Tx) error {
		// Delete users
		err := s.repo.Delete(ctx, uid)
		if err != nil {
			return fmt.Errorf("failed to create user in transaction: %v", err)
		}

		// Delete user aliases
		err = s.repo.DeleteUserAlias(ctx, uid)
		if err != nil {
			return fmt.Errorf("failed to delete user aliases in transaction: %v", err)
		}

		// Delete user logins
		err = s.repo.DeleteLogin(ctx, uid)
		if err != nil {
			return fmt.Errorf("failed to delete user logins in transaction: %v", err)
		}

		// Delete user devices
		err = s.deviceRepo.DeleteDeviceByUID(ctx, tx, uid)
		if err != nil {
			return fmt.Errorf("failed to delete user devices in transaction: %v", err)
		}

		return nil
	}

	err := s.repo.DeleteUserWithAssociations(ctx, handler)
	if err != nil {
		return status.INTERNAL, err
	}
	return status.SUCCESS, nil
}

func (s *Service) ForceDeleteUser(ctx context.Context, id string) (status.Code, error) {
	if id == "" {
		return status.USER_INVALID_ID, ErrInvalidUserID
	}

	err := s.repo.ForceDeleteUserWithAssociations(ctx, id)
	if err != nil {
		return status.INTERNAL, err
	}
	return status.SUCCESS, nil
}
