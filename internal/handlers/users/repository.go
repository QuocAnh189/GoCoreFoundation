package users

import (
	"context"

	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
)

type IRepository interface {
	Create(ctx context.Context, user *User) error
	CreateInBatches(ctx context.Context, users []*User, batchSize int) error
	FindById(ctx context.Context, id string) (*User, error)
	FindOne(ctx context.Context, opts ...db.FindOption) (*User, error)
	Find(ctx context.Context, opts ...db.FindOption) ([]*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string, opts ...db.FindOption) error
	Count(ctx context.Context, opts ...db.FindOption) (int64, error)
}

type UserRepository struct {
	db db.IDatabase
}

func NewUserRepository(db db.IDatabase) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *User) error {
	return r.db.Create(ctx, "users", user)
}

func (r *UserRepository) CreateInBatches(ctx context.Context, users []*User, batchSize int) error {
	return r.db.CreateInBatches(ctx, "users", users, batchSize)
}

func (r *UserRepository) FindById(ctx context.Context, id string) (*User, error) {
	var user User
	err := r.db.FindById(ctx, "users", id, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindOne(ctx context.Context, opts ...db.FindOption) (*User, error) {
	var user User
	err := r.db.FindOne(ctx, "users", &user, opts...)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Find(ctx context.Context, opts ...db.FindOption) ([]*User, error) {
	var users []*User
	err := r.db.Find(ctx, "users", &users, opts...)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, user *User) error {
	return r.db.Update(ctx, "users", user)
}

func (r *UserRepository) Delete(ctx context.Context, id string, opts ...db.FindOption) error {
	if id != "" {
		opts = append(opts, db.WithQuery(db.NewQuery("id = ?", id)))
	}
	return r.db.Delete(ctx, "users", &User{}, opts...)
}

func (r *UserRepository) Count(ctx context.Context, opts ...db.FindOption) (int64, error) {
	var total int64
	err := r.db.Count(ctx, "users", &User{}, &total, opts...)
	if err != nil {
		return 0, err
	}
	return total, nil
}
