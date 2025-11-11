package otp

import "github.com/QuocAnh189/GoCoreFoundation/internal/db"

type IRepository interface {
}

type Repository struct {
	db db.IDatabase
}

func NewRepository(db db.IDatabase) IRepository {
	return &Repository{
		db: db,
	}
}
