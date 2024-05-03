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
	jwtAuth  auth.JWTAuth
}

func NewUserService(userRepo repository.UserRepo, saltLen int, jwtAuth auth.JWTAuth) UserService {
	return UserService{
		userRepo: userRepo,
		saltLen:  saltLen,
		jwtAuth:  jwtAuth,
	}
}

func (s UserService) Register(ctx context.Context, userParam *entity.UserParam) (entity.UserResponse, error) {
	if !utils.IsValidFullName(userParam.Name) {
		return entity.UserResponse{}, msg.BadRequest(msg.ErrInvalidFullName)
	}

	if !utils.IsEmailValid(userParam.Email) {
		return entity.UserResponse{}, msg.BadRequest(msg.ErrInvalidEmail)
	}

	if !utils.IsSolidPassword(userParam.Password) {
		return entity.UserResponse{}, msg.BadRequest(msg.ErrInvalidPassword)
	}

	userParam.Salt = utils.GenerateRandomAlphaNumeric(int(s.saltLen))
	hashedPassword := auth.GenerateHash([]byte(userParam.Password), []byte(userParam.Salt))
	userParam.Password = hashedPassword

	user, err := s.userRepo.CreateUser(ctx, *userParam)
	if err != nil {
		return entity.UserResponse{}, err
	}

	accessToken, err := s.jwtAuth.GenerateToken(user.IdUser)
	if err != nil {
		return entity.UserResponse{}, err
	}

	return entity.UserResponse{
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: accessToken,
	}, nil
}

func (s UserService) Login(ctx context.Context, loginParam *entity.UserLoginParam) (entity.UserResponse, error) {
	if !utils.IsEmailValid(loginParam.Email) {
		return entity.UserResponse{}, msg.BadRequest(msg.ErrInvalidEmail)
	}

	if !utils.IsSolidPassword(loginParam.Password) {
		return entity.UserResponse{}, msg.BadRequest(msg.ErrWrongPassword)
	}

	user, err := s.userRepo.GetUserByEmail(ctx, loginParam.Email)
	if err != nil {
		return entity.UserResponse{}, err
	}

	err = auth.CompareHash(user.Password, loginParam.Password, user.Salt)
	if err != nil {
		return entity.UserResponse{}, msg.BadRequest(msg.ErrWrongPassword)
	}

	accessToken, err := s.jwtAuth.GenerateToken(user.IdUser)
	if err != nil {
		return entity.UserResponse{}, err
	}

	return entity.UserResponse{
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: accessToken,
	}, nil
}
