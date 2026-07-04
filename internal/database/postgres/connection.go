package postgres

import (
    "context"
    "fmt"

    "github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
    Pool *pgxpool.Pool
}

func NewConnection(ctx context.Context, connectionString string) (*DB, error) {
    if connectionString == "" {
        return nil, fmt.Errorf("connection string cannot be empty")
    }

    pool, err := pgxpool.New(ctx, connectionString)
    if err != nil {
        return nil, fmt.Errorf("failed to create connection pool: %w", err)
    }

    if err := pool.Ping(ctx); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    return &DB{Pool: pool}, nil
}

func (db *DB) Close() {
    db.Pool.Close()
}