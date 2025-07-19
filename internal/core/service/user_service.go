package service

import (
	"context"
	"corporate/config"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"
	"corporate/utils/auth"
	"corporate/utils/conv"
	"errors"

	"github.com/rs/zerolog/log"
)

var (
	err  error
	code string
)

type UserServiceInterface interface {
	LoginAdmin(ctx context.Context, req entity.UserEntity) (string, error)
}

type userService struct {
	userRepo repository.UserRepositoryInterface
	config   *config.Config
	jwtAuth  auth.JwtInterface
}

func (u *userService) LoginAdmin(ctx context.Context, req entity.UserEntity) (string, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		code = "[SERVICE] LoginAdmin - 1"
		log.Err(err).Msg(code)
		return "", err
	}

	if checkPass := conv.CheckPasswordHash(req.Password, user.Password); !checkPass {
		code = "[SERVICE] LoginAdmin - 2"
		err = errors.New("invalid password")
		log.Err(err).Msg(code)
		return "", err
	}
	jwtData := &entity.JwtData{
		UserID: float64(user.ID),
	}
	token, _, err := u.jwtAuth.GenerateToken(jwtData)
	if err != nil {
		code = "[SERVICE] LoginAdmin - 3"
		log.Err(err).Msg(code)
		return "", err
	}
	return token, nil
}

func NewUserService(userRepo repository.UserRepositoryInterface, cfg *config.Config, jwtAuth auth.JwtInterface) UserServiceInterface {
	return &userService{
		userRepo: userRepo,
		config:   cfg,
		jwtAuth:  jwtAuth,
	}
}
