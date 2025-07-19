package auth

import (
	"corporate/config"
	"corporate/internal/core/domain/entity"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtInterface interface {
	GenerateToken(data *entity.JwtData) (string, int64, error)
	VerifyAccesstoken(token string) (*entity.JwtData, error)
}

type Options struct {
	siginKey string
	issuer   string
}

func (o *Options) GenerateToken(data *entity.JwtData) (string, int64, error) {
	now := time.Now().Local()
	expireAt := now.Add(time.Hour * 24)
	data.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expireAt)
	data.RegisteredClaims.Issuer = o.issuer
	data.RegisteredClaims.NotBefore = jwt.NewNumericDate(now)
	acToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	accessToken, err := acToken.SignedString([]byte(o.siginKey))
	if err != nil {
		return "", 0, err
	}
	return accessToken, expireAt.Unix(), nil
}

func (o *Options) VerifyAccesstoken(token string) (*entity.JwtData, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		}
		return []byte(o.siginKey), nil
	})

	if err != nil {
		return nil, err
	}

	if parsedToken.Valid {
		claim, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			return nil, err
		}
		jwtData := &entity.JwtData{
			UserID: claim["user_id"].(float64),
		}
		return jwtData, nil
	}
	return nil, fmt.Errorf("Token is not valid")
}

func NewJwt(cfg *config.Config) JwtInterface {
	opt := new(Options)
	opt.siginKey = cfg.App.JwtSecretKey
	opt.issuer = cfg.App.JwtIssuer

	return opt
}
