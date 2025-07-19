package handler

import (
	"corporate/internal/adapter/handler/request"
	"corporate/internal/adapter/handler/response"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/service"
	"corporate/utils/conv"
	"net/http"

	//"github.com/gofiber/fiber/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UserHandler interface {
	LoginAdmin(c echo.Context) error
}

type userHandler struct {
	userService service.UserServiceInterface
}

var (
	err  error
	code string
)

func (u *userHandler) LoginAdmin(c echo.Context) error {
	var (
		req       = request.LoginRequest{}
		resp      = response.DefaultSuccessResponse{}
		respLogin = response.LoginResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	if err = c.Bind(&req); err != nil {
		code = "[HANDLER] LoginAdmin - 1"
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		code = "[HANDLER] LoginAdmin - 2"
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.UserEntity{
		Email:    req.Email,
		Password: req.Password,
	}
	token, err := u.userService.LoginAdmin(ctx, reqEntity)
	if err != nil {
		code = "[HANDLER] LoginAdmin - 3"
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respLogin.Token = token
	resp.Meta.Status = true
	resp.Meta.Message = "Success Login"
	resp.Data = respLogin
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewUserHandler(e *echo.Echo, userService service.UserServiceInterface) UserHandler {
	userHandler := &userHandler{
		userService: userService,
	}
	e.Use(middleware.Recover())
	e.POST("/login", userHandler.LoginAdmin)
	return userHandler
}
