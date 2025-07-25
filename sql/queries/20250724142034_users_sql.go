package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUsersSql, downUsersSql)
}

func upUsersSql(ctx context.Context, tx *sql.Tx) error {
	-- name: CreateUser :one
	INSERT INTO users (id, created_at, updated_at, email)
	VALUES ( gen_random_uuid, Now(), Now(), $1 ) RETURNING *;

	// This code is executed when the migration is applied.
	return nil
}

func downUsersSql(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
