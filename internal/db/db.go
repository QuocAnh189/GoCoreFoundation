package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
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

// Query represents a SQL query with arguments.
type Query struct {
	Query string
	Args  []any
}

// NewQuery creates a new Query instance.
func NewQuery(query string, args ...any) Query {
	return Query{
		Query: query,
		Args:  args,
	}
}

// IDatabase defines the database interface.
type IDatabase interface {
	GetDB() *sql.DB
	AutoMigrate(tableName string, models ...any) error
	WithTransaction(function func() error) error
	Create(ctx context.Context, tableName string, doc any) error
	CreateInBatches(ctx context.Context, tableName string, docs any, batchSize int) error
	Update(ctx context.Context, tableName string, doc any) error
	Delete(ctx context.Context, tableName string, value any, opts ...FindOption) error
	FindById(ctx context.Context, tableName string, id string, result any) error
	FindOne(ctx context.Context, tableName string, result any, opts ...FindOption) error
	Find(ctx context.Context, tableName string, result any, opts ...FindOption) error
	Count(ctx context.Context, tableName string, model any, total *int64, opts ...FindOption) error
	Close() error
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
	mysqlCfg.ParseTime = false

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

func (d *Database) AutoMigrate(tableName string, models ...any) error {
	return nil
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

func (d *Database) Create(ctx context.Context, tableName string, doc any) error {
	ctx, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	query, args, err := buildInsertQuery(doc)
	if err != nil {
		return fmt.Errorf("failed to build insert query: %v", err)
	}
	query = fmt.Sprintf("INSERT INTO %s %s", tableName, query[strings.Index(query, "("):])
	_, err = d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to create record: %v", err)
	}
	return nil
}

func (d *Database) CreateInBatches(ctx context.Context, tableName string, docs any, batchSize int) error {
	ctx, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	sliceVal := reflect.ValueOf(docs)
	if sliceVal.Kind() != reflect.Slice {
		return fmt.Errorf("docs must be a slice")
	}

	length := sliceVal.Len()
	if length == 0 {
		return nil
	}

	for i := 0; i < length; i += batchSize {
		end := i + batchSize
		if end > length {
			end = length
		}

		batch := sliceVal.Slice(i, end).Interface()
		query, args, err := buildBatchInsertQuery(batch)
		if err != nil {
			return fmt.Errorf("failed to build batch insert query: %v", err)
		}
		query = fmt.Sprintf("INSERT INTO %s %s", tableName, query[strings.Index(query, "("):])
		_, err = d.db.ExecContext(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("failed to create batch: %v", err)
		}
	}
	return nil
}

func (d *Database) Update(ctx context.Context, tableName string, doc any) error {
	ctx, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	query, args, err := buildUpdateQuery(doc)
	if err != nil {
		return fmt.Errorf("failed to build update query: %v", err)
	}
	query = fmt.Sprintf("UPDATE %s SET %s", tableName, query[strings.Index(query, "SET"):])
	_, err = d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update record: %v", err)
	}
	return nil
}

func (d *Database) Delete(ctx context.Context, tableName string, value any, opts ...FindOption) error {
	ctx, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	query, args := buildDeleteQuery(tableName, value, opts...)
	_, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete records: %v", err)
	}
	return nil
}

func (d *Database) FindById(ctx context.Context, tableName string, id string, result any) error {
	ctx, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", tableName)
	err := scanRow(d.db.QueryRowContext(ctx, query, id), result)
	if err != nil {
		return fmt.Errorf("failed to find by ID: %v", err)
	}
	return nil
}

func (d *Database) FindOne(ctx context.Context, tableName string, result any, opts ...FindOption) error {
	ctx, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	query, args := buildSelectQuery(tableName, result, opts...)
	row := d.db.QueryRowContext(ctx, query, args...)
	err := scanRow(row, result)
	if err != nil {
		return fmt.Errorf("failed to find one: %v", err)
	}
	return nil
}

func (d *Database) Find(ctx context.Context, tableName string, result any, opts ...FindOption) error {
	ctx, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	query, args := buildSelectQuery(tableName, result, opts...)
	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to find records: %v", err)
	}
	defer rows.Close()

	err = scanRows(rows, result)
	if err != nil {
		return fmt.Errorf("failed to scan rows: %v", err)
	}
	return nil
}

func (d *Database) Count(ctx context.Context, tableName string, model any, total *int64, opts ...FindOption) error {
	ctx, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	query, args := buildCountQuery(tableName, model, opts...)
	err := d.db.QueryRowContext(ctx, query, args...).Scan(total)
	if err != nil {
		return fmt.Errorf("failed to count records: %v", err)
	}
	return nil
}

// Close
func (d *Database) Close() error {
	return d.Close()
}
