package entity

import (
	"database/sql"
	"time"
)

type Cat struct {
	IdCat       uint32       `db:"id_cat" json:"id_cat"`
	Name        string       `db:"name" json:"name"`
	Race        string       `db:"race" json:"race"`
	Sex         string       `db:"sex" json:"sex"`
	AgeInMonth  int          `db:"age_in_month" json:"age_in_month"`
	Description string       `db:"description" json:"description"`
	ImageURL    string       `db:"image_url" json:"image_url"`
	IsMatch     bool         `db:"isMatch" json:"isMatch"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at" json:"updated_at"`
}

type CatParam struct {
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	AgeInMonth  int      `json:"ageInMonth"`
	Description string   `json:"description"`
	ImageURLs   []string `json:"imageUrls"`
}

type CreateCatRequest struct {
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	AgeInMonth  int      `json:"ageInMonth"`
	Description string   `json:"description"`
	ImageURLs   []string `json:"imageUrls"`
}

type CreateCatResponse struct {
	Message string        `json:"message"`
	Data    CreateCatData `json:"data"`
}

type UpdateCatResponse struct {
	Message string        `json:"message"`
	Data    UpdateCatData `json:"data"`
}

type CreateCatData struct {
	ID        uint32    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type UpdateCatData struct {
	ID        uint32    `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserCat struct {
	UserID int `json:"user_id"`
	CatID  int `json:"cat_id"`
}
