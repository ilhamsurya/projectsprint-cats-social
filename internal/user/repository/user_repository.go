package repository

import (
	"context"
	"projectsphere/cats-social/internal/user/entity"
	"projectsphere/cats-social/pkg/database"
	"projectsphere/cats-social/pkg/protocol/msg"
	"strings"
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
	query, args, err := r.dbConnector.SQLBuilder.
		Insert(userTableName).
		Columns("email", "name", "password").
		Values(param.Email, param.Name, param.Password).
		Suffix("RETURNING id_user, email, name, password, created_at, updated_at").
		ToSql()

	if err != nil {
		return entity.User{}, msg.InternalServerError(err.Error())
	}

	var row entity.User

	if err = r.dbConnector.DB.GetContext(ctx, &row, query, args...); err != nil {
		if strings.Contains(err.Error(), "unique") {
			return entity.User{}, msg.BadRequest(msg.ErrEmailAlreadyExist)
		} else {
			return entity.User{}, msg.InternalServerError(err.Error())
		}
	}

	return row, nil
}
