package repository

import (
	"context"
	"database/sql"
	"projectsphere/cats-social/internal/cat/entity"
	"projectsphere/cats-social/pkg/database"
	"projectsphere/cats-social/pkg/protocol/msg"
)

const catTableName string = "cats"

type CatRepo struct {
	dbConnector database.PostgresConnector
}

func NewCatRepo(dbConnector database.PostgresConnector) CatRepo {
	return CatRepo{
		dbConnector: dbConnector,
	}
}

func (r CatRepo) CreateCat(ctx context.Context, param entity.CatParam) (entity.Cat, error) {
	query := `
		INSERT INTO cats (name, race, sex, age_in_month, description, image_url)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id_cat, name, race, sex, age_in_month, description, image_url, isMatch, created_at, updated_at
	`

	var cat entity.Cat
	err := r.dbConnector.DB.QueryRowContext(
		ctx,
		query,
		param.Name,
		param.Race,
		param.Sex,
		param.AgeInMonth,
		param.Description,
		param.ImageURL,
	).Scan(
		&cat.IdCat,
		&cat.Name,
		&cat.Race,
		&cat.Sex,
		&cat.AgeInMonth,
		&cat.Description,
		&cat.ImageURL,
		&cat.IsMatch,
		&cat.CreatedAt,
		&cat.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Cat{}, msg.BadRequest("no rows were returned")
		}
		return entity.Cat{}, msg.InternalServerError(err.Error())
	}

	return cat, nil
}
