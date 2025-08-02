package handler

import (
	"corporate/config"
	"corporate/internal/adapter/handler/request"
	"corporate/internal/adapter/handler/response"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/service"
	"corporate/utils/conv"
	"corporate/utils/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type AboutCompanyHandlerInterface interface {
	CreateAboutCompany(c echo.Context) error
	FetchAllAboutCompany(c echo.Context) error
	FetchByIDAboutCompany(c echo.Context) error
	EditByIDAboutCompany(c echo.Context) error
	DeleteByIDAboutCompany(c echo.Context) error
	FetchAllCompanyHome(c echo.Context) error
}
type aboutCompanyHandler struct {
	AboutCompanyService service.AboutCompanyServiceInterface
}

// FetchAllCompanyHome implements AboutCompanyHandlerInterface.
func (cs *aboutCompanyHandler) FetchAllCompanyHome(c echo.Context) error {
	var (
		respCompany = response.AboutCompanyResponse{}
		resp        = response.DefaultSuccessResponse{}
		respError   = response.ErrorResponseDefault{}
		ctx         = c.Request().Context()
	)
	result, err := cs.AboutCompanyService.FetchAllCompanyAndKeynote(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllCompanyHome - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	respCompany.ID = result.ID
	respCompany.Description = result.Description
	for _, val := range result.Keynote {
		respCompany.CompanyKeyNotes = append(respCompany.CompanyKeyNotes, response.AboutCompanyKeynoteResponse{
			ID:             val.ID,
			AboutCompanyID: val.AboutCompanyID,
			Keynote:        val.Keynote,
			PathImage:      val.PathImage,
		})
	}
	resp.Meta.Message = "Success fetch all company home"
	resp.Meta.Status = true
	resp.Data = respCompany
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// CreateAboutCompany implements CLientSectionHandlerInterface.
func (cs *aboutCompanyHandler) CreateAboutCompany(c echo.Context) error {
	var (
		req       = request.AboutCompanyRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateAboutCompany - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateAboutCompany - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateAboutCompany - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.AboutCompanyEntity{
		Description: req.Description,
	}

	err = cs.AboutCompanyService.CreateAboutCompany(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateAboutCompany - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success create about company"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDAboutCompany implements CLientSectionHandlerInterface.
func (cs *aboutCompanyHandler) DeleteByIDAboutCompany(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteAboutCompany - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idCompany := c.Param("id")
	id, err := conv.StringToInt64(idCompany)
	if err != nil {
		log.Errorf("[HANDLER] DeleteAboutCompany - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.AboutCompanyService.DeleteByIDAboutCompany(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteAboutCompany - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success delete about company"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDAboutCompany implements CLientSectionHandlerInterface.
func (cs *aboutCompanyHandler) EditByIDAboutCompany(c echo.Context) error {
	var (
		req       = request.AboutCompanyRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditAboutCompany - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idCompany := c.Param("id")
	id, err := conv.StringToInt64(idCompany)
	if err != nil {
		log.Errorf("[HANDLER] EditAboutCompany - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditAboutCompany - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditAboutCompany - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.AboutCompanyEntity{
		ID:          id,
		Description: req.Description,
	}

	err = cs.AboutCompanyService.EditByIDAboutCompany(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditAboutCompany - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success edit about company"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllAboutCompany implements CLientSectionHandlerInterface.
func (cs *aboutCompanyHandler) FetchAllAboutCompany(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = []response.AboutCompanyResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLAboutCompany - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.AboutCompanyService.FetchAllAboutCompany(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLAboutCompany - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClient = append(respClient, response.AboutCompanyResponse{
			ID:          val.ID,
			Description: val.Description,
		})
	}
	resp.Meta.Message = "Success fetch all about company"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDAboutCompany implements CLientSectionHandlerInterface.
func (cs *aboutCompanyHandler) FetchByIDAboutCompany(c echo.Context) error {
	var (
		resp             = response.DefaultSuccessResponse{}
		respError        = response.ErrorResponseDefault{}
		ctx              = c.Request().Context()
		respAboutCompany = response.AboutCompanyResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID AboutCompany - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idCompany := c.Param("id")
	id, err := conv.StringToInt64(idCompany)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID AboutCompany - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.AboutCompanyService.FetchByIDAboutCompany(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID About Company - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respAboutCompany.ID = result.ID
	respAboutCompany.Description = result.Description
	resp.Meta.Message = "Success fetch about company by ID"
	resp.Meta.Status = true
	resp.Data = respAboutCompany
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewAboutCompanyHandler(e *echo.Echo, AboutCompanyService service.AboutCompanyServiceInterface, cfg *config.Config) AboutCompanyHandlerInterface {
	h := &aboutCompanyHandler{
		AboutCompanyService: AboutCompanyService,
	}

	mid := middleware.NewMiddleware(cfg)
	aboutCompanyApp := e.Group("/about-company")
	aboutCompanyApp.GET("", h.FetchAllCompanyHome)
	adminApp := aboutCompanyApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", h.CreateAboutCompany)
	adminApp.GET("", h.FetchAllAboutCompany)
	adminApp.GET("/:id", h.FetchByIDAboutCompany)
	adminApp.PUT("/:id", h.EditByIDAboutCompany)
	adminApp.DELETE("/:id", h.DeleteByIDAboutCompany)
	return h
}
