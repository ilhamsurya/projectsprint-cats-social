package repository

import (
	"context"
	"fmt"
	"projectsphere/cats-social/internal/user/entity"
	"projectsphere/cats-social/pkg/database"
	"projectsphere/cats-social/pkg/protocol/msg"
)

const userTableName string = "users"

type UserRepo struct {
	dbConnector database.PostgresConnector
}

func NewUserRepo(dbConnector database.PostgresConnector) UserRepo {
	return UserRepo{
		dbConnector: dbConnector,
	}
}

func (r UserRepo) CreateUser(ctx context.Context, param entity.UserParam) (entity.User, error) {
	query := `
	    INSERT INTO "cat" (name, race, sex, age_in_month, description)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id_cat
    `

	fmt.Print(ctx, query)

	result, err := r.dbConnector.DB.ExecContext(ctx, query, 3, param.Email, param.Name, param.Password)
	if err != nil {
		return entity.User{}, msg.InternalServerError(err.Error())
	}

	id, _ := result.LastInsertId()
	fmt.Println(id)

	var row entity.User
	err = r.dbConnector.DB.QueryRowContext(ctx,
		"SELECT id_user, email, name, password, created_at, updated_at FROM \"user\" WHERE id_user = $1",
		1,
	).Scan(&row.Id_user, &row.Email, &row.Name, &row.Password, &row.CreatedAt, &row.UpdatedAt)
	if err != nil {
		return entity.User{}, msg.InternalServerError(err.Error())
	}

	return row, nil
}
