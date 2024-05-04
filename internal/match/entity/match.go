package entity

import (
	"database/sql"
	"projectsphere/cats-social/internal/cat/entity"
	"time"
)

type MatchCat struct {
	IdMatch      uint32       `db:"id_match" json:"id_match"`
	IdUserCat    uint32       `db:"id_user_cat" json:"id_user_cat"`
	IdMatchedCat uint32       `db:"id_matched_cat" json:"id_matched_cat"`
	IsMatched    bool         `db:"is_matched" json:"is_matched"`
	UserCat      entity.Cat   `json:"user_cat"`
	MatchedCat   entity.Cat   `json:"matched_cat"`
	CreatedAt    time.Time    `db:"created_at" json:"created_at"`
	ApprovedAt   sql.NullTime `db:"approved_at" json:"approved_at"`
	RejectedAt   sql.NullTime `db:"rejected_at" json:"rejected_at"`
}

type MatchCatRequest struct {
	MatchCatId uint32 `json:"matchCatId"`
	UserCatId  uint32 `json:"userCatId"`
	Message    string `json:"message" validate:"required,min=5,max=120"`
}

type ProcessMatchRequest struct {
	MatchId uint32 `json:"matchId"`
}

type MatchCatResponse struct {
	Message string `json:"message"`
}

type IssuedBy struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type CatDetail struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	Description string    `json:"description"`
	AgeInMonth  int       `json:"ageInMonth"`
	ImageUrls   []string  `json:"imageUrls"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}

type DataDetail struct {
	ID             int       `json:"id"`
	IssuedBy       IssuedBy  `json:"issuedBy"`
	MatchCatDetail CatDetail `json:"matchCatDetail"`
	UserCatDetail  CatDetail `json:"userCatDetail"`
	Message        string    `json:"message"`
	CreatedAt      time.Time `json:"createdAt"`
}
