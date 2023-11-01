package sql

import (
	"context"
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore() (*MySQLStore, error) {
	dbname := os.Getenv("DATABASE_NAME")
	conurl := os.Getenv("MYSQLDB_URL")
	dburl := conurl + "/" + dbname + "?parseTime=true"

	// Connect to the database.
	db, err := sql.Open("mysql", dburl)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	return &MySQLStore{db: db}, nil
}

func (s *MySQLStore) AddReord(ctx context.Context, sp string, args ...any) (*sql.Rows, error) {
	// Prepare the SQL statement for calling the stored procedure
	stmt, err := s.db.PrepareContext(ctx, sp)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}

func (s *MySQLStore) GetRecords(ctx context.Context, sp string, args ...any) (*sql.Rows, error) {
	// Prepare the SQL statement for calling the stored procedure
	stmt, err := s.db.PrepareContext(ctx, sp)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the stored procedure with parameter values
	return stmt.QueryContext(ctx, args...)
}

func (s *MySQLStore) GetRecordByArgs(ctx context.Context, sp string, args ...any) (*sql.Rows, error) {

	// Prepare the SQL statement for calling the stored procedure
	stmt, err := s.db.PrepareContext(ctx, sp)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the stored procedure with parameter values
	return stmt.QueryContext(ctx, args...)
}

func (s *MySQLStore) DeleteRecordByArgs(ctx context.Context, sp string, args ...any) (*sql.Rows, error) {
	// Prepare the SQL statement for calling the stored procedure
	stmt, err := s.db.PrepareContext(ctx, sp)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the stored procedure with parameter values
	return stmt.QueryContext(ctx, args...)
}
