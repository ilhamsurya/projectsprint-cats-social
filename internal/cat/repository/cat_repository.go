package repository

import (
	"context"
	"database/sql"
	"fmt"
	"projectsphere/cats-social/internal/cat/entity"
	"projectsphere/cats-social/pkg/database"
	"projectsphere/cats-social/pkg/protocol/msg"
	"strconv"
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

func (r CatRepo) GetCat(ctx context.Context, param entity.GetCatParam, ageOperator string, age int) ([]entity.Cat, error) {
	query := `
		SELECT 
			c.id_cat, c.name, c.race, c.sex, c.age_in_month, c.description, 
			ci.id_image, ci.id_cat, ci.image, 
			mc.id_match, mc.is_matched
		FROM "cat" c
		JOIN "cat_image" ci ON ci.id_cat = c.id_cat 
		JOIN "user" u ON u.id_user = c.id_user
		LEFT JOIN "match_cat" mc ON mc.id_user_cat = c.id_cat
		WHERE c.deleted_at IS NULL 
	`
	args := []interface{}{}
	argsCount := 1

	if param.IdCat != nil {
		query += fmt.Sprintf(" AND c.id_cat = $%d", argsCount)
		args = append(args, &param.IdCat)
		argsCount++
	}
	if param.Race != "" {
		query += fmt.Sprintf(" AND c.race = LOWER($%d)", argsCount)
		args = append(args, &param.Race)
		argsCount++
	}
	if param.AgeInMonth != "" {
		switch ageOperator {
		case ">":
			query += fmt.Sprintf(" AND c.age_in_month > $%d", argsCount)
		case "<":
			query += fmt.Sprintf(" AND c.age_in_month < $%d", argsCount)
		default:
			query += fmt.Sprintf(" AND c.age_in_month = $%d", argsCount)
		}
		args = append(args, age)
		argsCount++
	}
	if param.Owned != nil {
		query += fmt.Sprintf(" AND c.id_user = $%d", argsCount)
		args = append(args, &param.IdUser)
		fmt.Println(&param.IdUser)
		argsCount++
	}
	if param.HasMatched != nil {
		if *param.HasMatched {
			query += fmt.Sprintf(" AND mc.is_matched = $%d", argsCount)
			args = append(args, &param.HasMatched)
			argsCount++
		} else {
			query += fmt.Sprintf(" AND (mc.is_matched = $%d OR mc.is_matched IS NULL)", argsCount)
			args = append(args, &param.HasMatched)
			argsCount++
		}
	}
	if param.Sex != "" {
		query += fmt.Sprintf(" AND c.sex = $%d", argsCount)
		args = append(args, &param.Sex)
		argsCount++
	}
	if param.Search != "" {
		query += " AND c.name ILIKE '%' || $" + strconv.Itoa(argsCount) + " || '%'"
		args = append(args, &param.Search)
		argsCount++
	}

	// if param.Limit != nil {
	// 	query += fmt.Sprintf(" LIMIT $%d", argsCount)
	// 	args = append(args, *param.Limit)
	// 	argsCount++
	// }

	// if param.Offset != nil {
	// 	query += fmt.Sprintf(" OFFSET $%d", argsCount)
	// 	args = append(args, &param.Offset)
	// 	argsCount++
	// }

	// fmt.Println(query)

	rows, err := r.dbConnector.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return []entity.Cat{}, msg.InternalServerError(err.Error())
	}
	defer rows.Close()

	var cats []entity.Cat
	var catsMap = make(map[int]entity.Cat)

	for rows.Next() {
		var cat = entity.Cat{}
		var image = entity.CatImage{}
		var idMatch *uint32
		var isMatched *bool
		var err = rows.Scan(&cat.IdCat, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &image.IdImage, &image.IdCat, &image.Image, &idMatch, &isMatched)

		if err != nil {
			return []entity.Cat{}, msg.InternalServerError(err.Error())
		}

		catTemp, ok := catsMap[int(cat.IdCat)]
		if !ok {
			catTemp = entity.Cat{
				IdCat:       cat.IdCat,
				Name:        cat.Name,
				Race:        cat.Race,
				Sex:         cat.Sex,
				AgeInMonth:  cat.AgeInMonth,
				Description: cat.Description,
				CreatedAt:   cat.CreatedAt,
				UpdatedAt:   cat.UpdatedAt,
				CatImage:    make([]entity.CatImage, 0),
				MatchCat:    make([]entity.MatchCat, 0),
			}
		}
		catTemp.CatImage = append(catTemp.CatImage, image)
		if idMatch != nil {
			catTemp.MatchCat = append(catTemp.MatchCat, entity.MatchCat{IdMatch: *idMatch, IsMatched: *isMatched})
		}
		catsMap[int(cat.IdCat)] = catTemp
	}

	for _, cat := range catsMap {
		cats = append(cats, cat)
	}

	// rs, err := json.MarshalIndent(cats, "", " ")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Print(string(rs))

	return cats, nil
}

func (r CatRepo) DeleteCat(ctx context.Context, catID int, userID int) error {
	query := `
		UPDATE cat SET deleted_at = NOW() 
		WHERE id_cat = $1 
		AND id_user = $2
		AND deleted_at IS NULL
	`

	res, err := r.dbConnector.DB.ExecContext(ctx, query, catID, userID)
	if err != nil {
		return msg.InternalServerError(err.Error())
	}

	rowEffect, err := res.RowsAffected()
	if err != nil {
		return msg.InternalServerError(err.Error())
	}

	if rowEffect == 0 {
		return msg.NotFound("id is not found")
	}

	return nil
}
