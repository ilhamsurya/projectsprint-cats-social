package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Id_user   uint32       `db:"id_user" json:"id_user"`
	Email     string       `db:"email" json:"email"`
	Name      string       `db:"name" json:"name"`
	Password  string       `db:"password" json:"-"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
}

type UserParam struct {
	Email    string
	Name     string
	Password string
}

type UserRegisterResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
