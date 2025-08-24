package db

import "database/sql"

type Config struct {
}

type IDatabase interface {
	Close() error
}

type Database struct {
	db *sql.DB
}

func NewDatabase(config *Config) (*Database, error) {
	return &Database{}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
