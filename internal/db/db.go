package db

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/configs"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/colors"
	"github.com/go-sql-driver/mysql"
)

const (
	DatabaseTimeout = time.Second * 5
)

// Database wraps a sql.DB connection.
type Database struct {
	db *sql.DB
}

// NewDatabase initializes a new database connection.
func NewDatabase(config *configs.DBConfig) (*Database, error) {
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

func (d *Database) WithTransaction(function func(tx *sql.Tx) error) error {
	log.Println("Starting transaction")
	tx, err := d.db.Begin()
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	if err := function(tx); err != nil {
		log.Printf("Rolling back transaction due to error: %v", err)
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("Failed to rollback transaction: %v", rbErr)
			return fmt.Errorf("failed to rollback transaction: %v (original error: %v)", rbErr, err)
		}
		log.Println("Transaction rolled back successfully")
		return err
	}

	log.Println("Committing transaction")
	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	log.Println("Transaction committed successfully")
	return nil
}

func (d *Database) Query(ctx context.Context, tx *sql.Tx, query string, args ...any) (*sql.Rows, error) {
	_, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	d.logInputSQL(colors.FGGreen, query, args...)
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.QueryContext(ctx, query, args...)
	} else {
		rows, err = d.db.QueryContext(ctx, query, args...)
	}
	if err != nil {
		d.logQueryError(err)
	}
	d.logQueryRowsResult(rows)
	return rows, err
}

func (d *Database) QueryRow(ctx context.Context, tx *sql.Tx, query string, args ...any) *sql.Row {
	_, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	d.logInputSQL(colors.FGGreen, query, args...)
	if tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	}
	return d.db.QueryRowContext(ctx, query, args...)
}

func (d *Database) Exec(ctx context.Context, tx *sql.Tx, query string, args ...any) (sql.Result, error) {
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
	var result sql.Result
	var err error
	if tx != nil {
		result, err = tx.ExecContext(ctx, query, args...)
	} else {
		result, err = d.db.ExecContext(ctx, query, args...)
	}
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
