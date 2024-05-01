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
	Name        string `json:"name"`
	Race        string `json:"race"`
	Sex         string `json:"sex"`
	AgeInMonth  int    `json:"age_in_month"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	IsMatch     bool   `json:"isMatch"`
}

type CreateCatRequest struct {
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	Race        string   `json:"race" validate:"required,oneof=Persian MaineCoon Siamese Ragdoll Bengal Sphynx BritishShorthair Abyssinian ScottishFold Birman"`
	Sex         string   `json:"sex" validate:"required,oneof=male female"`
	AgeInMonth  int      `json:"ageInMonth" validate:"required,min=1,max=120082"`
	Description string   `json:"description" validate:"required,min=1,max=200"`
	ImageURLs   []string `json:"imageUrls" validate:"required,min=1,dive,url"`
}

type CreateCatResponse struct {
	Message string        `json:"message"`
	Data    CreateCatData `json:"data"`
}

type CreateCatData struct {
	ID        uint32    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
