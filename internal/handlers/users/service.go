package users

import (
	"context"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/pagination"
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

	dto := BuildCreateUserDTO(req)

	dto.ID, err = uuid.GenerateUUIDV7()
	if err != nil {
		return status.INTERNAL, nil, err
	}

	result, err := s.repo.CreateUserWithAssociations(ctx, dto)
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

func (s *Service) DeleteUser(ctx context.Context, id string) (status.Code, error) {
	if id == "" {
		return status.USER_INVALID_ID, ErrInvalidUserID
	}

	err := s.repo.DeleteUserWithAssociations(ctx, id)
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
