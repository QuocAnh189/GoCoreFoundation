package users

import (
	"context"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/pagination"
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
	resp, err := s.repo.List(ctx, req)
	if err != nil {
		return status.INTERNAL, nil, nil, err
	}
	return status.SUCCESS, resp.Users, resp.Pagination, nil
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
	resStatus, err := ValidateCreateUserRequest(req)
	if err != nil {
		return resStatus, nil, err
	}

	dto := BuildCreateUserDTO(req)
	result, err := s.repo.Create(ctx, dto)
	if err != nil {
		return status.INTERNAL, nil, err
	}
	return status.SUCCESS, result, nil
}

func (s *Service) UpdateUser(ctx context.Context, req *UpdateUserRequest) (status.Code, *User, error) {
	resStatus, err := ValidateUpdateUserRequest(req)
	if err != nil {
		return resStatus, nil, err
	}

	dto := BuildUpdateUserDTO(req)
	result, err := s.repo.Update(ctx, dto)

	if err != nil {
		return status.INTERNAL, nil, err
	}
	return status.SUCCESS, result, nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) (status.Code, error) {
	if id == "" {
		return status.USER_INVALID_ID, ErrInvalidUserID
	}

	err := s.repo.Delete(ctx, id)
	if err != nil {
		return status.INTERNAL, err
	}
	return status.SUCCESS, nil
}
