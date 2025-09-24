package db

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/colors"
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
	mysqlCfg.TLS = &tls.Config{MinVersion: tls.VersionTLS12, MaxVersion: tls.VersionTLS12}
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
	_, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	d.logInputSQL(colors.FGGreen, query, args...)
	rows, err := d.db.Query(query, args...)
	if err != nil {
		d.logQueryError(err)
	}
	d.logQueryRowsResult(rows)

	return rows, err
}

func (d *Database) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	_, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	d.logInputSQL(colors.FGGreen, query, args...)
	row := d.db.QueryRow(query, args...)
	d.logQueryRowResult(row)

	return row
}

func (d *Database) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	_, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	var color colors.Color

	normalized := strings.TrimSpace(strings.ToUpper(query))
	switch {
	case strings.HasPrefix(normalized, "SELECT"):
		color = colors.FGGreen
	case strings.HasPrefix(normalized, "INSERT"):
		color = colors.FGCyan
	case strings.HasPrefix(normalized, "UPDATE"):
		color = colors.FGYellow
	case strings.HasPrefix(normalized, "DELETE"):
		color = colors.FGRed
	default:
		color = colors.FGOrange
	}

	d.logInputSQL(color, query, args...)
	result, err := d.db.Exec(query, args...)
	if err != nil {
		d.logQueryError(err)
	}
	if result != nil {
		d.logExecResult(result)
	}

	return result, err
}

func (d *Database) PingContext(ctx context.Context) error {
	_, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	d.logInputSQL(colors.FGPurple, "PING")
	return d.db.PingContext(ctx)
}

func (d *Database) Close() error {
	return d.db.Close()
}
