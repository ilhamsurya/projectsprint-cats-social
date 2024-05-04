package service

import (
	"context"
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

func (s *MatchService) Create(ctx context.Context, matchParam entity.MatchCat) (entity.MatchCatResponse, error) {
	if matchParam.IdMatchedCat == 0 || matchParam.IdUserCat == 0 {
		return entity.MatchCatResponse{}, &msg.RespError{
			Code:    400,
			Message: "Invalid payload: required fields are missing",
		}
	}

	matchCat, err := s.catRepo.GetCatByID(ctx, int(matchParam.IdMatchedCat))
	if err != nil {
		if respErr, ok := err.(*msg.RespError); ok && respErr.Code == 404 {
			return entity.MatchCatResponse{}, &msg.RespError{
				Code:    404,
				Message: "Matched cat not found",
			}
		}
		return entity.MatchCatResponse{}, err
	}

	userCat, err := s.catRepo.GetCatByID(ctx, int(matchParam.IdUserCat))
	if err != nil {
		if respErr, ok := err.(*msg.RespError); ok && respErr.Code == 404 {
			return entity.MatchCatResponse{}, &msg.RespError{
				Code:    404,
				Message: "User cat not found",
			}
		}
		return entity.MatchCatResponse{}, err
	}

	if matchCat.Sex == userCat.Sex {
		return entity.MatchCatResponse{}, &msg.RespError{
			Code:    400,
			Message: "Cat's gender is the same",
		}
	}

	matchedCatOwnerID, err := s.catRepo.GetCatOwner(ctx, int(matchParam.IdMatchedCat))
	if err != nil {
		return entity.MatchCatResponse{}, &msg.RespError{
			Code:    400,
			Message: "Not match owner",
		}
	}

	userCatOwnerID, err := s.catRepo.GetCatOwner(ctx, int(matchParam.IdUserCat))
	if err != nil {
		return entity.MatchCatResponse{}, &msg.RespError{
			Code:    400,
			Message: "Not match owner",
		}
	}

	if matchedCatOwnerID == userCatOwnerID {
		return entity.MatchCatResponse{}, &msg.RespError{
			Code:    400,
			Message: "MatchCatId and UserCatId are from the same owner",
		}
	}

	if matchCat.IdCat == userCat.IdCat {
		return entity.MatchCatResponse{}, &msg.RespError{
			Code:    400,
			Message: "Either MatchCatId or UserCatId is already matched",
		}
	}

	if !s.catRepo.CatExists(ctx, int(matchParam.IdMatchedCat)) || !s.catRepo.CatExists(ctx, int(matchParam.IdUserCat)) {
		return entity.MatchCatResponse{}, &msg.RespError{
			Code:    404,
			Message: "Neither MatchCatId nor UserCatId is found",
		}
	}

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
