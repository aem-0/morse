package postgres

import (
	"context"
	"database/sql"

	"morse/db/migrations"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func MigrateDatabase(ctx context.Context, connectionString string) error {

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(migrations.MigrationFiles)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.UpContext(ctx, db, "."); err != nil {
		return err
	}

	return nil
}
