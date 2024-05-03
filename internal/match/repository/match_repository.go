package repository

import (
	"context"
	"projectsphere/cats-social/internal/match/entity"
	"projectsphere/cats-social/pkg/database"
	"time"
)

type MatchRepo struct {
	dbConnector database.PostgresConnector
}

func NewMatchRepo(dbConnector database.PostgresConnector) MatchRepo {
	return MatchRepo{
		dbConnector: dbConnector,
	}
}

func (r MatchRepo) CreateMatch(ctx context.Context, param entity.MatchCat) (entity.MatchCat, error) {
	var match entity.MatchCat
	query := `
        INSERT INTO "match_cats" (id_user_cat, id_matched_cat, is_matched, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id_match, id_user_cat, id_matched_cat, is_matched, created_at
    `
	err := r.dbConnector.DB.QueryRowContext(ctx, query,
		param.IdUserCat,
		param.IdMatchedCat,
		false,      // Default is_matched value
		time.Now(), // Current time as created_at
	).Scan(
		&match.IdMatch,
		&match.IdUserCat,
		&match.IdMatchedCat,
		&match.IsMatched,
		&match.CreatedAt,
	)
	if err != nil {
		return entity.MatchCat{}, err
	}

	return match, nil
}
