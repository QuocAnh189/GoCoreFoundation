package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	DatabaseTimeout = time.Second * 5
)

// Config holds database connection parameters.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// Database wraps a sql.DB connection.
type Database struct {
	db *sql.DB
}

// NewDatabase initializes a new database connection.
func NewDatabase(config *Config) (*Database, error) {
	address := fmt.Sprintf("%s:%s", config.Host, config.Port)

	mysqlCfg := mysql.NewConfig()
	mysqlCfg.User = config.User
	mysqlCfg.Passwd = config.Password
	mysqlCfg.Addr = address
	mysqlCfg.DBName = config.Name
	mysqlCfg.AllowNativePasswords = true
	mysqlCfg.Net = "tcp"
	mysqlCfg.ParseTime = true

	db, err := sql.Open("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Set up connection pool
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 2)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	log.Print("Connected to database successfully")
	return &Database{db: db}, nil
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func (d *Database) WithTransaction(function func() error) error {
	tx, err := d.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	if err := function(); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("failed to rollback transaction: %v (original error: %v)", rbErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

func (d *Database) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, query, args...)
}

func (d *Database) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	return d.db.QueryRowContext(ctx, query, args...)
}

func (d *Database) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}

func (d *Database) PingContext(ctx context.Context) error {
	return d.db.PingContext(ctx)
}

func (d *Database) Close() error {
	return d.Close()
}
