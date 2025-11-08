package login

import "github.com/QuocAnh189/GoCoreFoundation/internal/db"

type IRepository interface {
	// Define repository methods here
}

type Repository struct {
	db db.IDatabase
}

func NewRepository(db db.IDatabase) *Repository {
	return &Repository{
		db: db,
	}
}
