package users

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
)

type UserService struct {
	repo IRepository
}

func NewService(repo IRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *User) error {
	if err := s.validateUser(ctx, user, true); err != nil {
		return errors.New("invalid user data: " + err.Error())
	}

	now := time.Now().UTC().Format("2006-01-02 15:04:05")
	user.CreateDt = &now
	user.ModifyDt = &now
	if user.CreateID == nil {
		createID := 1 // Replace with authenticated user ID
		user.CreateID = &createID
		user.ModifyID = &createID
	}

	count, err := s.repo.Count(ctx, db.WithQuery(db.NewQuery("email = ?", user.Email)))
	if err != nil {
		return errors.New("failed to check email uniqueness: " + err.Error())
	}
	if count > 0 {
		return errors.New(fmt.Sprintf("email already exists: email %s is already taken", user.Email))
	}

	count, err = s.repo.Count(ctx, db.WithQuery(db.NewQuery("phone = ?", user.Phone)))
	if err != nil {
		return errors.New("failed to check phone uniqueness: " + err.Error())
	}
	if count > 0 {
		return errors.New(fmt.Sprintf("phone already exists: phone %s is already taken", user.Phone))
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return errors.New("failed to create user: " + err.Error())
	}
	return nil
}

func (s *UserService) CreateInBatches(ctx context.Context, users []*User, batchSize int) error {
	for _, user := range users {
		if err := s.validateUser(ctx, user, true); err != nil {
			return errors.New(fmt.Sprintf("invalid user data for %s: %s", user.Email, err.Error()))
		}
		now := time.Now().UTC().Format("2006-01-02 15:04:05")
		user.CreateDt = &now
		user.ModifyDt = &now
		if user.CreateID == nil {
			createID := 1 // Replace with authenticated user ID
			user.CreateID = &createID
			user.ModifyID = &createID
		}
	}

	emails := make([]string, len(users))
	phones := make([]string, len(users))
	for i, user := range users {
		emails[i] = user.Email
		phones[i] = user.Phone
	}
	count, err := s.repo.Count(ctx, db.WithQuery(db.NewQuery("email IN ?", emails)))
	if err != nil {
		return errors.New("failed to check email uniqueness: " + err.Error())
	}
	if count > 0 {
		return errors.New("duplicate emails in batch: some emails are already taken")
	}
	count, err = s.repo.Count(ctx, db.WithQuery(db.NewQuery("phone IN ?", phones)))
	if err != nil {
		return errors.New("failed to check phone uniqueness: " + err.Error())
	}
	if count > 0 {
		return errors.New("duplicate phones in batch: some phones are already taken")
	}

	if err := s.repo.CreateInBatches(ctx, users, batchSize); err != nil {
		return errors.New("failed to create users in batch: " + err.Error())
	}
	return nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (*User, error) {
	if id == 0 {
		return nil, errors.New("invalid input: ID is required")
	}

	user, err := s.repo.FindById(ctx, fmt.Sprintf("%d", id))
	if err != nil {
		return nil, errors.New("failed to find user: " + err.Error())
	}
	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, opts ...db.FindOption) (*User, error) {
	user, err := s.repo.FindOne(ctx, opts...)
	if err != nil {
		return nil, errors.New("failed to find user: " + err.Error())
	}
	return user, nil
}

func (s *UserService) GetListUser(ctx context.Context, opts ...db.FindOption) ([]*User, error) {
	users, err := s.repo.Find(ctx, opts...)
	if err != nil {
		return nil, errors.New("failed to find users: " + err.Error())
	}
	return users, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *User) error {
	if user.ID == 0 {
		return errors.New("invalid input: ID is required")
	}

	if err := s.validateUser(ctx, user, false); err != nil {
		return errors.New("invalid user data: " + err.Error())
	}

	now := time.Now().UTC().Format("2006-01-02 15:04:05")
	user.ModifyDt = &now
	if user.ModifyID == nil {
		modifyID := 1 // Replace with authenticated user ID
		user.ModifyID = &modifyID
	}

	existingUser, err := s.repo.FindById(ctx, fmt.Sprintf("%d", user.ID))
	if err != nil {
		return errors.New("failed to check existing user: " + err.Error())
	}
	if existingUser.Email != user.Email {
		count, err := s.repo.Count(ctx, db.WithQuery(db.NewQuery("email = ? AND id != ?", user.Email, user.ID)))
		if err != nil {
			return errors.New("failed to check email uniqueness: " + err.Error())
		}
		if count > 0 {
			return errors.New(fmt.Sprintf("email already exists: email %s is already taken", user.Email))
		}
	}

	if existingUser.Phone != user.Phone {
		count, err := s.repo.Count(ctx, db.WithQuery(db.NewQuery("phone = ? AND id != ?", user.Phone, user.ID)))
		if err != nil {
			return errors.New("failed to check phone uniqueness: " + err.Error())
		}
		if count > 0 {
			return errors.New(fmt.Sprintf("phone already exists: phone %s is already taken", user.Phone))
		}
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return errors.New("failed to update user: " + err.Error())
	}
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int64, opts ...db.FindOption) error {
	if id == 0 && len(opts) == 0 {
		return errors.New("invalid input: ID or conditions required")
	}

	if id != 0 {
		opts = append(opts, db.WithQuery(db.NewQuery("id = ?", id)))
	}
	if err := s.repo.Delete(ctx, fmt.Sprintf("%d", id), opts...); err != nil {
		return errors.New("failed to delete user: " + err.Error())
	}
	return nil
}
