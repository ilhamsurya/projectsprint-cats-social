package service

import (
	"context"
	"projectsphere/cats-social/internal/user/entity"
	"projectsphere/cats-social/internal/user/repository"
	"projectsphere/cats-social/pkg/middleware/auth"
	"projectsphere/cats-social/pkg/protocol/msg"
	"projectsphere/cats-social/pkg/utils"
)

type UserService struct {
	userRepo repository.UserRepo
	saltLen  int
}

func NewUserService(userRepo repository.UserRepo, saltLen int) UserService {
	return UserService{
		userRepo: userRepo,
		saltLen:  saltLen,
	}
}

func (s UserService) Register(ctx context.Context, userParam *entity.UserParam) (entity.UserRegisterResponse, error) {
	if !utils.IsValidFullName(userParam.Name) {
		return entity.UserRegisterResponse{}, msg.BadRequest(msg.ErrInvalidFullName)
	}

	if !utils.IsEmailValid(userParam.Email) {
		return entity.UserRegisterResponse{}, msg.BadRequest(msg.ErrInvalidEmail)
	}

	if !utils.IsSolidPassword(userParam.Password) {
		return entity.UserRegisterResponse{}, msg.BadRequest(msg.ErrInvalidPassword)
	}

	userParam.Salt = utils.GenerateRandomAlphaNumeric(int(s.saltLen))
	hashedPassword := auth.GenerateHash([]byte(userParam.Password), []byte(userParam.Salt))
	userParam.Password = hashedPassword

	user, err := s.userRepo.CreateUser(ctx, *userParam)
	if err != nil {
		return entity.UserRegisterResponse{}, err
	}

	accessToken, err := auth.GenerateToken(user.IdUser, "ACCESS_TOKEN")
	if err != nil {
		return entity.UserRegisterResponse{}, err
	}

	return entity.UserRegisterResponse{
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: accessToken,
	}, nil
}
