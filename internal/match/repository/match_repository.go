package repository

import (
	"context"
	"database/sql"
	"errors"
	"projectsphere/cats-social/internal/match/entity"
	"projectsphere/cats-social/pkg/database"
	"projectsphere/cats-social/pkg/protocol/msg"
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
        INSERT INTO "match_cats" (id_user_cat, id_matched_cat, created_at)
        VALUES ($1, $2, $3)
        RETURNING id_match, id_user_cat, id_matched_cat, created_at
    `
	err := r.dbConnector.DB.QueryRowContext(ctx, query,
		param.IdUserCat,
		param.IdMatchedCat,
		time.Now(), // Current time as created_at
	).Scan(
		&match.IdMatch,
		&match.IdUserCat,
		&match.IdMatchedCat,
		&match.CreatedAt,
	)
	if err != nil {
		return entity.MatchCat{}, err
	}

	return match, nil
}

func (r MatchRepo) GetMatchByID(ctx context.Context, matchID int) (entity.MatchCat, error) {
	var match entity.MatchCat
	query := `
        SELECT id_match, id_user_cat, id_matched_cat, approved_at, rejected_at
        FROM "match_cats" WHERE id_match = $1
    `
	err := r.dbConnector.DB.QueryRowContext(ctx, query, matchID).Scan(
		&match.IdMatch, &match.IdUserCat, &match.IdMatchedCat, &match.ApprovedAt, &match.RejectedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.MatchCat{}, errors.New("404: Match not found")
		}
		return entity.MatchCat{}, errors.New("500: Internal server error")
	}

	return match, nil
}

func (r MatchRepo) DeleteMatchByMatchId(ctx context.Context, matchID int) error {
	query := `
		DELETE FROM match_cats WHERE id_match = $1;
    `
	res, err := r.dbConnector.DB.ExecContext(ctx, query, matchID)
	if err != nil {
		return msg.InternalServerError(err.Error())
	}

	rowEffect, err := res.RowsAffected()
	if err != nil {
		return msg.InternalServerError(err.Error())
	}

	if rowEffect == 0 {
		return msg.NotFound("match request is not found")
	}

	return nil
}

func (r MatchRepo) RejectByMatchId(ctx context.Context, matchID int) error {
	query := `
		UPDATE match_cats SET rejected_at = NOW() 
		WHERE match_id = $1
		AND rejected_at IS NULL AND approved_at IS NULL
	`
	res, err := r.dbConnector.DB.ExecContext(ctx, query, matchID)
	if err != nil {
		return msg.InternalServerError(err.Error())
	}

	rowEffect, err := res.RowsAffected()
	if err != nil {
		return msg.InternalServerError(err.Error())
	}

	if rowEffect == 0 {
		return msg.NotFound("match request is not found")
	}

	return nil
}

func (r MatchRepo) ApproveByMatchId(ctx context.Context, matchID int) error {
	query := `
		UPDATE match_cats SET approved_at = NOW() 
		WHERE match_id = $1
		AND rejected_at IS NULL AND approved_at IS NULL
	`
	res, err := r.dbConnector.DB.ExecContext(ctx, query, matchID)
	if err != nil {
		return msg.InternalServerError(err.Error())
	}

	rowEffect, err := res.RowsAffected()
	if err != nil {
		return msg.InternalServerError(err.Error())
	}

	if rowEffect == 0 {
		return msg.NotFound("match request is not found")
	}

	return nil
}

func (r MatchRepo) DeleteMatchByApprove(ctx context.Context, param entity.MatchCat) error {
	query := `
		DELETE FROM match_cats
		WHERE (id_user_cat IN ($2, $3) OR id_matched_cat IN ($2, $3))
		AND id_match != $1;
    `
	res, err := r.dbConnector.DB.ExecContext(ctx, query, param.IdMatch, param.IdUserCat, param.IdMatchedCat)
	if err != nil {
		return msg.InternalServerError(err.Error())
	}

	rowEffect, err := res.RowsAffected()

	if rowEffect == 0 {
		return nil
	}

	return nil
}
