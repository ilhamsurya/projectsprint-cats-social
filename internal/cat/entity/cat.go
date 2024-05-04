package entity

import (
	"database/sql"
	"time"
)

type Cat struct {
	IdCat       uint32       `db:"id_cat" json:"id_cat"`
	IdUser      uint32       `db:"id_user" json:"id_user"`
	Name        string       `db:"name" json:"name"`
	Race        string       `db:"race" json:"race"`
	Sex         string       `db:"sex" json:"sex"`
	AgeInMonth  int          `db:"age_in_month" json:"age_in_month"`
	Description string       `db:"description" json:"description"`
	CatImage    []CatImage   `json:"cat_image"`
	MatchCat    []MatchCat   `json:"match_cat"`
	HasMatched  bool         `db:"has_matched" json:"has_matched"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

type CatImage struct {
	IdImage uint32 `db:"id_image" json:"id_image"`
	IdCat   uint32 `db:"id_cat" json:"id_cat"`
	Image   string `db:"image" json:"image"`
}

type CatParam struct {
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	AgeInMonth  int      `json:"ageInMonth"`
	Description string   `json:"description"`
	HasMatched  bool     `json:"hasMatched"`
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

type GetCatParam struct {
	IdUser     int    `json:"id_user"`
	IdCat      *int   `json:"id"`
	Limit      *int   `json:"limit"`
	Offset     *int   `json:"offset"`
	Race       string `json:"race"`
	Sex        string `json:"sex"`
	HasMatched *bool  `json:"hasMatched"`
	AgeInMonth string `json:"ageInMonth"`
	Owned      *bool  `json:"owned"`
	Search     string `json:"search"`
}

type GetCatData struct {
	IdCat       uint32    `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	AgeInMonth  int       `json:"ageInMonth"`
	Description string    `json:"description"`
	ImageUrl    []string  `json:"imageUrl"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}

type MatchCat struct {
	IdMatch      uint32       `db:"id_match" json:"id_match"`
	IdUserCat    uint32       `db:"id_user_cat" json:"id_user_cat"`
	IdMatchedCat uint32       `db:"id_matched_cat" json:"id_matched_cat"`
	CreatedAt    time.Time    `db:"created_at" json:"created_at"`
	ApprovedAt   sql.NullTime `db:"approved_at" json:"approved_at"`
	RejectedAt   sql.NullTime `db:"rejected_at" json:"rejected_at"`
}
