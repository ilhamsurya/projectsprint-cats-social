package service

import (
	"context"
	"errors"
	catRepository "projectsphere/cats-social/internal/cat/repository"
	"projectsphere/cats-social/internal/match/entity"
	"projectsphere/cats-social/internal/match/repository"
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
