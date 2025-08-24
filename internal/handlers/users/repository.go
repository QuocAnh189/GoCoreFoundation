package users

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
)

type IRepository interface {
}

type UserRepository struct {
	db db.IDatabase
}

func NewUserRepository(db db.IDatabase) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
