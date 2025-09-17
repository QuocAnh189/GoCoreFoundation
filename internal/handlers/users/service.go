package users

import (
	"context"

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

func (s *Service) ListUsers(ctx context.Context, req *ListUserRequest) ([]*User, *pagination.Pagination, error) {
	resp, err := s.repo.List(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return resp.Users, resp.Pagination, nil
}

func (s *Service) GetUserByID(ctx context.Context, id string) (*User, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, ErrUserNotFound
	}
	return result, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	err := ValidateCreateUserRequest(req)
	if err != nil {
		return nil, err
	}

	dto := BuildCreateUserDTO(req)
	return s.repo.Create(ctx, dto)
}

func (s *Service) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	err := ValidateUpdateUserRequest(req)
	if err != nil {
		return nil, err
	}

	dto := BuildUpdateUserDTO(req)
	return s.repo.Update(ctx, dto)
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidUserID
	}

	return s.repo.Delete(ctx, id)
}
