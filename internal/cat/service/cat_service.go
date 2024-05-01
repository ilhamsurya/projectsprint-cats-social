package service

import (
	"context"
	"projectsphere/cats-social/internal/cat/entity"
	"projectsphere/cats-social/internal/cat/repository"

	"github.com/go-playground/validator/v10"
)

type CatService struct {
	catRepo repository.CatRepo
}

func NewCatService(catRepo repository.CatRepo) CatService {
	return CatService{
		catRepo: catRepo,
	}
}

func (s CatService) Create(ctx context.Context, catParam entity.CatParam) (entity.CreateCatData, error) {
	if err := validateCatParam(catParam); err != nil {
		return entity.CreateCatData{}, err
	}

	cat, err := s.catRepo.CreateCat(ctx, catParam)
	if err != nil {
		return entity.CreateCatData{}, err
	}

	return entity.CreateCatData{
		ID:        cat.IdCat,
		CreatedAt: cat.CreatedAt,
	}, nil
}

func validateCatParam(catParam entity.CatParam) error {
	validate := validator.New()
	err := validate.Struct(catParam)
	if err != nil {
		return err
	}
	return nil
}
