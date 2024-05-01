package repository

import (
	"context"
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
	query, args, err := r.dbConnector.SQLBuilder.
		Insert(catTableName).
		Columns("name", "race", "sex", "age_in_month", "description", "image_url").
		Values(param.Name, param.Race, param.Sex, param.AgeInMonth, param.Description, param.ImageURL).
		Suffix("RETURNING id_cat, name, race, sex, age_in_month, description, image_url, isMatch, created_at, updated_at").
		ToSql()

	if err != nil {
		return entity.Cat{}, msg.InternalServerError(err.Error())
	}

	var row entity.Cat

	if err = r.dbConnector.DB.GetContext(ctx, &row, query, args...); err != nil {
		return entity.Cat{}, msg.InternalServerError(err.Error())
	}

	return row, nil
}
