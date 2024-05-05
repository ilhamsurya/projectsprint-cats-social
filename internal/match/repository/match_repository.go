package repository

import (
	"context"
	"database/sql"
	"errors"
	catEntity "projectsphere/cats-social/internal/cat/entity"
	"projectsphere/cats-social/internal/match/entity"
	userEntity "projectsphere/cats-social/internal/user/entity"
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

func (r MatchRepo) GetMatchRequest(ctx context.Context, userID int) ([]entity.MatchCat, error) {
	query := `
		SELECT 
			mc.id_match, mc.id_user_cat, mc.id_matched_cat, mc.created_at, mc.approved_at, mc.rejected_at,
			cr.id_cat, cr.name, cr.race, cr.sex, cr.age_in_month, cr.description,
			cri.id_image, cri.id_cat, cri.image, 
			mr.id_cat, mr.name, mr.race, mr.sex, mr.age_in_month, mr.description,
			mri.id_image, mri.id_cat, mri.image, 
			u.id_user, u.name, u.email
		FROM "match_cats" as mc 
		JOIN "cats" cr ON cr.id_cat = mc.id_user_cat
		JOIN "cat_images" cri ON cri.id_cat = cr.id_cat
		JOIN "users" u ON u.id_user = cr.id_user
		JOIN "cats" mr ON  mr.id_cat = mc.id_matched_cat
		JOIN "cat_images" mri ON mri.id_cat = mr.id_cat
    `

	rows, err := r.dbConnector.DB.QueryContext(ctx, query)
	if err != nil {
		return []entity.MatchCat{}, msg.InternalServerError(err.Error())
	}
	defer rows.Close()

	var matchCatMap = make(map[int]entity.MatchCat)

	for rows.Next() {
		matchCat := entity.MatchCat{}
		matchCatDetail := catEntity.Cat{}
		imageUrlsMatchCat := catEntity.CatImage{}
		userCatDetail := catEntity.Cat{}
		imageUrlsUserCat := catEntity.CatImage{}
		issuer := userEntity.User{}

		var err = rows.Scan(
			&matchCat.IdMatch, &matchCat.IdUserCat, &matchCat.IdMatchedCat, &matchCat.CreatedAt, &matchCat.ApprovedAt, &matchCat.RejectedAt,
			&matchCatDetail.IdCat, &matchCatDetail.Name, &matchCatDetail.Race, &matchCatDetail.Sex, &matchCatDetail.AgeInMonth, &matchCatDetail.Description,
			&imageUrlsMatchCat.IdImage, &imageUrlsMatchCat.IdCat, &imageUrlsMatchCat.Image,
			&userCatDetail.IdCat, &userCatDetail.Name, &userCatDetail.Race, &userCatDetail.Sex, &userCatDetail.AgeInMonth, &userCatDetail.Description,
			&imageUrlsUserCat.IdImage, &imageUrlsUserCat.IdCat, &imageUrlsUserCat.Image,
			&issuer.IdUser, &issuer.Name, &issuer.Email,
		)

		if err != nil {
			return []entity.MatchCat{}, msg.InternalServerError(err.Error())
		}

		matchCatTmp, ok := matchCatMap[int(matchCat.IdMatch)]
		if !ok {
			matchCatTmp = entity.MatchCat{
				IdMatch:      matchCat.IdMatch,
				IdUserCat:    matchCat.IdUserCat,
				IdMatchedCat: matchCat.IdMatchedCat,
				MatchedCat:   matchCatDetail,
				UserCat:      userCatDetail,
				CreatedAt:    matchCat.CreatedAt,
				ApprovedAt:   matchCat.ApprovedAt,
				RejectedAt:   matchCat.RejectedAt,
			}

			matchCatTmp.UserCat.User = issuer
			matchCatTmp.UserCat.CatImage = make([]catEntity.CatImage, 0)
			matchCatTmp.MatchedCat.CatImage = make([]catEntity.CatImage, 0)
		}

		matchCatTmp.UserCat.CatImage = append(matchCatTmp.UserCat.CatImage, imageUrlsUserCat)
		matchCatTmp.MatchedCat.CatImage = append(matchCatTmp.MatchedCat.CatImage, imageUrlsMatchCat)
		matchCatMap[int(matchCat.IdMatch)] = matchCatTmp
	}

	var matchCats []entity.MatchCat
	for _, mc := range matchCatMap {
		matchCats = append(matchCats, mc)
	}

	return matchCats, nil
}
