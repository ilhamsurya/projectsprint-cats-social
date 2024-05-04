package entity

import (
	"database/sql"
	"time"
)

type MatchCat struct {
	IdMatch      uint32       `db:"id_match" json:"id_match"`
	IdUserCat    uint32       `db:"id_user_cat" json:"userCatId"`
	IdMatchedCat uint32       `db:"id_matched_cat" json:"matchCatId"`
	IsMatched    bool         `db:"is_matched" json:"is_matched"`
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
