package database

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PostgresConnector struct {
	DB         *sqlx.DB
	SQLBuilder squirrel.StatementBuilderType
}

func NewPostgresConnector(ctx context.Context, db *sqlx.DB) PostgresConnector {
	return PostgresConnector{
		DB:         db,
		SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
