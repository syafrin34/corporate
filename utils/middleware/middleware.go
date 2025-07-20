package middleware

import (
	"corporate/config"
	"corporate/internal/adapter/handler/response"
	"corporate/utils/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Middleware interface {
	CheckToken() echo.MiddlewareFunc
}

type Options struct {
	authJwt auth.JwtInterface
}

func (o *Options) CheckToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var errorResponse response.ErrorResponseDefault
			//ambil header authorization
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				errorResponse.Meta.Message = "Missing Authorization Header"
				errorResponse.Meta.Status = false
				return c.JSON(http.StatusUnauthorized, errorResponse)
			}
			// validasi format header authorization
			parts := strings.Split(authHeader, "Bearer ")
			if len(parts) != 2 {
				errorResponse.Meta.Message = "Invalid Authorization Header Format"
				errorResponse.Meta.Status = false
				return c.JSON(http.StatusUnauthorized, errorResponse)
			}
			//ambil token dari header
			tokenString := parts[1]
			// verifikasi token
			claims, err := o.authJwt.VerifyAccesstoken(tokenString)
			if err != nil {
				errorResponse.Meta.Message = "Invalid Token"
				errorResponse.Meta.Status = false
				return c.JSON(http.StatusUnauthorized, errorResponse)
			}
			c.Set("user", claims)
			// lanjutka ke handler berikutnya
			return next(c)
		}
	}
}

func NewMiddleware(cfg *config.Config) Middleware {
	opt := new(Options)
	opt.authJwt = auth.NewJwt(cfg)
	return opt
}
