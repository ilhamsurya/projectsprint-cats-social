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
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (s UserService) Register(ctx context.Context, userParam entity.UserParam) (entity.UserRegisterResponse, error) {
	if !utils.IsValidFullName(userParam.Name) {
		return entity.UserRegisterResponse{}, msg.BadRequest(msg.ErrInvalidFullName)
	}

	if !utils.IsEmailValid(userParam.Email) {
		return entity.UserRegisterResponse{}, &msg.RespError{
			Code:    409,
			Message: msg.ErrInvalidEmail,
		}
	}

	if !utils.IsSolidPassword(userParam.Password) {
		return entity.UserRegisterResponse{}, msg.BadRequest(msg.ErrInvalidPassword)
	}

	user, err := s.userRepo.CreateUser(ctx, userParam)
	if err != nil {
		return entity.UserRegisterResponse{}, err
	}

	accessToken, err := auth.GenerateToken(user.Id_user, "ACCESS_TOKEN")
	if err != nil {
		return entity.UserRegisterResponse{}, err
	}

	return entity.UserRegisterResponse{
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: accessToken,
	}, nil
}
