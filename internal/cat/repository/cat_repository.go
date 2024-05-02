package repository

import (
	"context"
	"database/sql"
	"fmt"
	"projectsphere/cats-social/internal/cat/entity"
	"projectsphere/cats-social/pkg/database"
	"projectsphere/cats-social/pkg/protocol/msg"
)

type CatRepo struct {
	dbConnector database.PostgresConnector
}

func NewCatRepo(dbConnector database.PostgresConnector) CatRepo {
	return CatRepo{
		dbConnector: dbConnector,
	}
}

func (r CatRepo) CreateCat(ctx context.Context, param entity.CatParam) (entity.Cat, error) {
	var cat entity.Cat
	query := `
        INSERT INTO "cat" (name, race, sex, age_in_month, description)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id_cat
    `
	err := r.dbConnector.DB.QueryRow(query, param.Name,
		param.Race,
		param.Sex,
		param.AgeInMonth,
		param.Description).Scan(&cat.IdCat)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Cat{}, msg.BadRequest("no rows were returned")
		}
		return entity.Cat{}, msg.InternalServerError(err.Error())
	}

	// Insert images into cat_image table
	for _, imageURL := range param.ImageURLs {
		_, err := r.insertCatImage(ctx, cat.IdCat, imageURL)
		if err != nil {
			return entity.Cat{}, msg.InternalServerError(err.Error())
		}
	}

	cat.Name = param.Name
	cat.Race = param.Race
	cat.Sex = param.Sex
	cat.AgeInMonth = param.AgeInMonth
	cat.Description = param.Description

	return cat, nil
}

func (r CatRepo) insertCatImage(ctx context.Context, catID uint32, imageURL string) (int, error) {
	query := `
		INSERT INTO cat_image (id_cat, image)
		VALUES ($1, $2)
		RETURNING id_image
	`

	var idImage int
	err := r.dbConnector.DB.QueryRowContext(
		ctx,
		query,
		catID,
		imageURL,
	).Scan(
		&idImage,
	)

	if err != nil {
		return 0, err
	}

	return idImage, nil
}

// UpdateCat updates the cat information in the database.
func (r CatRepo) UpdateCat(ctx context.Context, catID int, catParam entity.CatParam) (entity.Cat, error) {
	query := `
		UPDATE cat
		SET name = $1, race = $2, sex = $3, age_in_month = $4, description = $5, updated_at = CURRENT_TIMESTAMP
		WHERE id_cat = $6
		RETURNING id_cat, name, race, sex, age_in_month, description, created_at, updated_at
	`

	fmt.Print(catID, catParam)
	// Execute the update query
	row := r.dbConnector.DB.QueryRowContext(ctx, query, catParam.Name, catParam.Race, catParam.Sex, catParam.AgeInMonth, catParam.Description, catID)

	// Scan the updated cat from the database row
	var updatedCat entity.Cat
	err := row.Scan(&updatedCat.IdCat, &updatedCat.Name, &updatedCat.Race, &updatedCat.Sex, &updatedCat.AgeInMonth, &updatedCat.Description, &updatedCat.CreatedAt, &updatedCat.UpdatedAt)
	if err != nil {
		return entity.Cat{}, err
	}

	return updatedCat, nil
}
