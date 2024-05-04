package service

import (
	"context"
	"fmt"
	"net/http"
	"projectsphere/cats-social/internal/cat/entity"
	"projectsphere/cats-social/internal/cat/repository"
	"projectsphere/cats-social/pkg/protocol/msg"
	"strconv"
	"strings"
)

type CatService struct {
	catRepo repository.CatRepo
}

func NewCatService(catRepo repository.CatRepo) CatService {
	return CatService{
		catRepo: catRepo,
	}
}

func (s CatService) Create(ctx context.Context, catParam entity.CatParam) (entity.CreateCatResponse, error) {
	// Validate cat parameters
	if err := s.validateCatParam(catParam); err != nil {
		return entity.CreateCatResponse{}, &msg.RespError{
			Code:    400,
			Message: "request doesn't pass validation",
		}
	}

	// Create the cat
	cat, err := s.catRepo.CreateCat(ctx, catParam)
	if err != nil {
		return entity.CreateCatResponse{}, err
	}

	return entity.CreateCatResponse{
		Message: "success",
		Data: entity.CreateCatData{
			ID:        cat.IdCat,
			CreatedAt: cat.CreatedAt,
		},
	}, nil
}

func (s CatService) Update(ctx context.Context, catID int, catParam entity.CatParam) (entity.UpdateCatResponse, error) {
	if err := s.validateCatParam(catParam); err != nil {
		return entity.UpdateCatResponse{}, &msg.RespError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	cat, err := s.catRepo.UpdateCat(ctx, catID, catParam)
	if err != nil {
		return entity.UpdateCatResponse{}, err
	}

	return entity.UpdateCatResponse{
		Message: "success",
		Data: entity.UpdateCatData{
			ID:        cat.IdCat,
			UpdatedAt: cat.UpdatedAt.Time,
		},
	}, nil
}

func (s CatService) validateCatParam(catParam entity.CatParam) error {
	// Map to store validation errors
	validationErrors := make(map[string]string)

	// Validate name if not empty
	if catParam.Name != "" && len(catParam.Name) > 30 {
		validationErrors["name"] = "name must be between 1 and 30 characters"
	}

	// Validate race if not empty
	validRaces := map[string]bool{
		"Persian":           true,
		"Maine Coon":        true,
		"Siamese":           true,
		"Ragdoll":           true,
		"Bengal":            true,
		"Sphynx":            true,
		"British Shorthair": true,
		"Abyssinian":        true,
		"Scottish Fold":     true,
		"Birman":            true,
	}
	if catParam.Race != "" && !validRaces[catParam.Race] {
		validationErrors["race"] = "invalid race"
	}

	// Validate sex if not empty
	if catParam.Sex != "" && catParam.Sex != "male" && catParam.Sex != "female" {
		validationErrors["sex"] = "sex must be either 'male' or 'female'"
	}

	// Validate ageInMonth if not empty
	if catParam.AgeInMonth != 0 && (catParam.AgeInMonth < 1 || catParam.AgeInMonth > 120082) {
		validationErrors["ageInMonth"] = "ageInMonth must be between 1 and 120082"
	}

	// Validate description if not empty
	if catParam.Description != "" && len(catParam.Description) > 200 {
		validationErrors["description"] = "description must be between 1 and 200 characters"
	}

	// Validate imageUrls if not empty
	if len(catParam.ImageURLs) != 0 {
		for i, url := range catParam.ImageURLs {
			if url == "" {
				validationErrors[fmt.Sprintf("imageUrls[%d]", i)] = "imageUrls must not contain empty URLs"
			} else if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
				validationErrors[fmt.Sprintf("imageUrls[%d]", i)] = "invalid URL"
			}
		}
	}

	// Check if there are any validation errors
	if len(validationErrors) > 0 {
		// Construct error message from validation errors
		var errorMsgs []string
		for field, msg := range validationErrors {
			errorMsgs = append(errorMsgs, fmt.Sprintf("%s: %s", field, msg))
		}
		return fmt.Errorf(strings.Join(errorMsgs, "; "))
	}

	return nil
}

func (s CatService) Get(ctx context.Context, getCatParam entity.GetCatParam) ([]entity.GetCatData, error) {
	var age = 0
	var ageOperator = ""
	if getCatParam.AgeInMonth != "" {
		switch getCatParam.AgeInMonth[0:1] {
		case ">":
			ageOperator = ">"
			val, err := strconv.Atoi(getCatParam.AgeInMonth[1:])
			if err != nil {
				ageOperator = ""
			}
			age = val
		case "<":
			ageOperator = "<"
			val, err := strconv.Atoi(getCatParam.AgeInMonth[1:])
			if err != nil {
				ageOperator = ""
			}
			age = val
		default:
			ageOperator = "="
			val, err := strconv.Atoi(getCatParam.AgeInMonth[0:])
			if err != nil {
				ageOperator = ""
			}
			age = val
		}
	}

	cats, err := s.catRepo.GetCat(ctx, getCatParam, ageOperator, age)
	if err != nil {
		return []entity.GetCatData{}, err
	}

	var catRes = []entity.GetCatData{}
	for _, v := range cats {
		var images = []string{}
		for _, ci := range v.CatImage {
			images = append(images, ci.Image)
		}

		var hasMatched bool
		for _, mc := range v.MatchCat {
			if mc.IsMatched {
				hasMatched = true
				break
			}
		}

		data := entity.GetCatData{
			IdCat:       v.IdCat,
			Name:        v.Name,
			Race:        v.Race,
			Sex:         v.Sex,
			AgeInMonth:  v.AgeInMonth,
			Description: v.Description,
			ImageUrl:    images,
			HasMatched:  hasMatched,
			CreatedAt:   v.CreatedAt,
		}
		catRes = append(catRes, data)
	}

	return catRes, nil
}

func (s CatService) Delete(ctx context.Context, catID int, userID int) error {
	err := s.catRepo.DeleteCat(ctx, catID, userID)
	if err != nil {
		return err
	}

	return nil
}
