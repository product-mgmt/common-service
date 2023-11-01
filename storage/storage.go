package storage

import (
	"context"
	"database/sql"
)

type MySQLStorage interface {
	AddReord(ctx context.Context, sp string, args ...any) (*sql.Rows, error)
	GetRecords(ctx context.Context, sp string, args ...any) (*sql.Rows, error)
	GetRecordByArgs(ctx context.Context, sp string, args ...any) (*sql.Rows, error)
	DeleteRecordByArgs(ctx context.Context, sp string, args ...any) (*sql.Rows, error)
}
