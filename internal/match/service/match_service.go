package service

import (
	"context"
	"errors"
	catRepository "projectsphere/cats-social/internal/cat/repository"
	"projectsphere/cats-social/internal/match/entity"
	"projectsphere/cats-social/internal/match/repository"
	"projectsphere/cats-social/pkg/protocol/msg"
)

type MatchService struct {
	matchRepo repository.MatchRepo
	catRepo   catRepository.CatRepo
}

func NewMatchService(matchRepo repository.MatchRepo, catRepo catRepository.CatRepo) MatchService {
	return MatchService{
		matchRepo: matchRepo,
		catRepo:   catRepo,
	}
}

func (s MatchService) Create(ctx context.Context, matchParam entity.MatchCat) (entity.MatchCatResponse, error) {
	// Check for missing or expired token
	if ctx.Value("token") == nil {
		return entity.MatchCatResponse{}, errors.New("401: Request token is missing")
	}

	// Fetch the cat information
	matchCat, err := s.catRepo.GetCatByID(ctx, int(matchParam.IdMatchedCat))
	if err != nil {
		return entity.MatchCatResponse{}, err
	}

	userCat, err := s.catRepo.GetCatByID(ctx, int(matchParam.IdUserCat))
	if err != nil {
		return entity.MatchCatResponse{}, err
	}

	// Check if the cat's gender is the same
	if matchCat.Sex == userCat.Sex {
		return entity.MatchCatResponse{}, errors.New("400: Cat's gender is the same")
	}

	// Check if the cat IDs are from the same owner
	matchedCatOwnerID, err := s.catRepo.GetCatOwner(ctx, int(matchParam.IdMatchedCat))
	if err != nil {
		return entity.MatchCatResponse{}, err
	}

	userCatOwnerID, err := s.catRepo.GetCatOwner(ctx, int(matchParam.IdUserCat))
	if err != nil {
		return entity.MatchCatResponse{}, err
	}

	if matchedCatOwnerID == userCatOwnerID {
		return entity.MatchCatResponse{}, errors.New("400: MatchCatId and UserCatId are from the same owner")
	}

	// Check if either matchCatId or userCatId already matched
	if matchCat.IdCat == userCat.IdCat {
		return entity.MatchCatResponse{}, errors.New("400: Either MatchCatId or UserCatId is already matched")
	}

	// Check if neither matchCatId nor userCatId is found
	if !s.catRepo.CatExists(ctx, int(matchParam.IdMatchedCat)) || !s.catRepo.CatExists(ctx, int(matchParam.IdUserCat)) {
		return entity.MatchCatResponse{}, errors.New("404: Neither MatchCatId nor UserCatId is found")
	}

	// // Check if userCatId belongs to the user
	// if exists, err := s.catRepo.IsUserCatAssociationValid(ctx, matchParam.UserID, int(matchParam.IdUserCat)); err != nil {
	// 	return entity.MatchCatResponse{}, err
	// } else if !exists {
	// 	return entity.MatchCatResponse{}, errors.New("404: UserCatId does not belong to the user")
	// }

	// Create the match
	cat, err := s.matchRepo.CreateMatch(ctx, matchParam)
	if cat.IdMatch != 0 && err != nil {
		return entity.MatchCatResponse{}, err
	}

	return entity.MatchCatResponse{
		Message: "success",
	}, nil
}

func (s MatchService) Delete(ctx context.Context, matchID int, userID int) error {

	//Get Match Info
	match, err := s.matchRepo.GetMatchByID(ctx, int(matchID))
	if err != nil {
		return msg.BadRequest(err.Error())
	}

	// Fetch the cat information
	userCat, err := s.catRepo.GetCatByID(ctx, int(match.IdUserCat))
	if err != nil {
		return msg.BadRequest(err.Error())
	}

	// Check if userCatId belongs to the user
	if userCat.IdUser != uint32(userID) {
		return msg.Unauthorization("either cat or match request don't belong to user")
	}

	//Check if match already processed
	if match.ApprovedAt.Valid || match.RejectedAt.Valid {
		return msg.BadRequest("matchId is already approved / rejected")
	}

	// Delete the match
	err = s.matchRepo.DeleteMatchByMatchId(ctx, matchID)
	if err != nil {
		return msg.BadRequest(err.Error())
	}

	return nil
}

func (s MatchService) RejectMatchRequest(ctx context.Context, matchParam entity.ProcessMatchRequest, userID int) error {

	//Get Match Info
	match, err := s.matchRepo.GetMatchByID(ctx, int(matchParam.MatchId))
	if err != nil {
		return msg.BadRequest(err.Error())
	}

	// Fetch the cat information
	matchCat, err := s.catRepo.GetCatByID(ctx, int(match.IdMatchedCat))
	if err != nil {
		return msg.BadRequest(err.Error())
	}

	// Check if matchCatId belongs to the user
	if matchCat.IdUser != uint32(userID) {
		return msg.Unauthorization("either cat or match request don't belong to user")
	}

	//Check if match already processed
	if match.ApprovedAt.Valid || match.RejectedAt.Valid {
		return msg.BadRequest("matchId is no longer valid")
	}

	// Delete the match
	err = s.matchRepo.RejectByMatchId(ctx, int(matchParam.MatchId))
	if err != nil {
		return msg.BadRequest(err.Error())
	}

	return nil
}

func (s MatchService) ApproveMatchRequest(ctx context.Context, matchParam entity.ProcessMatchRequest, userID int) error {

	//Get Match Info
	match, err := s.matchRepo.GetMatchByID(ctx, int(matchParam.MatchId))
	if err != nil {
		return msg.BadRequest(err.Error())
	}

	// Fetch the cat information
	matchCat, err := s.catRepo.GetCatByID(ctx, int(match.IdMatchedCat))
	if err != nil {
		return msg.BadRequest(err.Error())
	}

	// Check if matchCatId belongs to the user
	if matchCat.IdUser != uint32(userID) {
		return msg.Unauthorization("either cat or match request don't belong to user")
	}

	//Check if match already processed
	if match.ApprovedAt.Valid || match.RejectedAt.Valid {
		return msg.BadRequest("matchId is no longer valid")
	}

	// Approve the match request
	err = s.matchRepo.ApproveByMatchId(ctx, int(matchParam.MatchId))
	if err != nil {
		return msg.BadRequest(err.Error())
	}

	err = s.matchRepo.DeleteMatchByApprove(ctx, match)
	if err != nil {
		return msg.BadRequest(err.Error())
	}

	return nil
}

func (s MatchService) GetMatchRequest(ctx context.Context, userID int) ([]entity.DataDetail, error) {
	res, err := s.matchRepo.GetMatchRequest(ctx, userID)
	if err != nil {
		return []entity.DataDetail{}, msg.InternalServerError(err.Error())
	}

	dataDetails := []entity.DataDetail{}

	for _, v := range res {
		hasMatched := false
		if v.ApprovedAt.Valid {
			hasMatched = true
		}

		data := entity.DataDetail{
			ID: int(v.IdMatch),
			IssuedBy: entity.IssuedBy{
				Name:      v.UserCat.User.Name,
				Email:     v.UserCat.User.Email,
				CreatedAt: v.UserCat.CreatedAt,
			},
			UserCatDetail: entity.CatDetail{
				ID:          int(v.UserCat.IdCat),
				Name:        v.UserCat.Name,
				Race:        v.UserCat.Race,
				Sex:         v.UserCat.Sex,
				Description: v.UserCat.Description,
				AgeInMonth:  v.UserCat.AgeInMonth,
				CreatedAt:   v.CreatedAt,
				HasMatched:  hasMatched,
				ImageUrls:   make([]string, 0),
			},
			MatchCatDetail: entity.CatDetail{
				ID:          int(v.MatchedCat.IdCat),
				Name:        v.MatchedCat.Name,
				Race:        v.MatchedCat.Race,
				Sex:         v.MatchedCat.Sex,
				Description: v.MatchedCat.Description,
				AgeInMonth:  v.MatchedCat.AgeInMonth,
				CreatedAt:   v.CreatedAt,
				HasMatched:  hasMatched,
				ImageUrls:   make([]string, 0),
			},
			CreatedAt: v.CreatedAt,
		}

		for _, imgU := range v.UserCat.CatImage {
			data.UserCatDetail.ImageUrls = append(data.UserCatDetail.ImageUrls, imgU.Image)
		}

		for _, imgM := range v.MatchedCat.CatImage {
			data.MatchCatDetail.ImageUrls = append(data.MatchCatDetail.ImageUrls, imgM.Image)
		}

		dataDetails = append(dataDetails, data)
	}

	return dataDetails, nil
}
