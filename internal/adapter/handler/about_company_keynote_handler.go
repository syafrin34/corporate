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

type AboutCompanyKeynoteHandlerInterface interface {
	CreateAboutCompanyKeynote(c echo.Context) error
	FetchAllAboutCompanyKeynote(c echo.Context) error
	FetchByIDAboutCompanyKeynote(c echo.Context) error
	EditByIDAboutCompanyKeynote(c echo.Context) error
	DeleteByIDAboutCompanyKeynote(c echo.Context) error
	FetchByCompanyID(c echo.Context) error
}
type aboutCompanyKeynoteHandler struct {
	AboutCompanyKeynoteService service.AboutCompanyKeynoteServiceInterface
}

// CreateAboutCompanyKeynote implements CLientSectionHandlerInterface.
func (a *aboutCompanyKeynoteHandler) CreateAboutCompanyKeynote(c echo.Context) error {
	var (
		req       = request.AboutCompanyKeynoteRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateAboutCompanyKeynote - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateAboutCompanyKeynote - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateAboutCompanyKeynote - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.AboutCompanyKeynoteEntity{
		AboutCompanyID: req.AboutCompanyID,
		Keynote:        req.Keynote,
		PathImage:      req.PathImage,
	}

	err = a.AboutCompanyKeynoteService.CreateAboutCompanyKeynote(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateAboutCompanyKeynote - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success create about company keynote"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

func (a *aboutCompanyKeynoteHandler) FetchByCompanyID(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = []response.AboutCompanyKeynoteResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByCompanyId - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAboutCompany := c.Param("companyId")
	id, err := conv.StringToInt64(idAboutCompany)
	if err != nil {
		log.Errorf("[HANDLER] FetchByCompanyID  - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	results, err := a.AboutCompanyKeynoteService.FetchByCompanyID(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByCompanyID - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClient = append(respClient, response.AboutCompanyKeynoteResponse{
			ID:                      val.ID,
			AboutCompanyID:          val.AboutCompanyID,
			Keynote:                 val.Keynote,
			PathImage:               val.PathImage,
			AboutCompanyDescription: val.AboutCompanyDescription,
		})
	}
	resp.Meta.Message = "Success fetch by company id keynote"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// DeleteByIDAboutCompanyKeynote implements CLientSectionHandlerInterface.
func (a *aboutCompanyKeynoteHandler) DeleteByIDAboutCompanyKeynote(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteAboutCompanyKeynote - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idCompanyKeynote := c.Param("id")
	id, err := conv.StringToInt64(idCompanyKeynote)
	if err != nil {
		log.Errorf("[HANDLER] DeleteAboutCompanyKeynote - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = a.AboutCompanyKeynoteService.DeleteByIDAboutCompanyKeynote(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteAboutCompanyKeynote - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success delete about company keynote"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDAboutCompanyKeynote implements CLientSectionHandlerInterface.
func (a *aboutCompanyKeynoteHandler) EditByIDAboutCompanyKeynote(c echo.Context) error {
	var (
		req       = request.AboutCompanyKeynoteRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditAboutCompanyKeynote - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idCompanyKeynote := c.Param("id")
	id, err := conv.StringToInt64(idCompanyKeynote)
	if err != nil {
		log.Errorf("[HANDLER] EditAboutCompanyKeynote - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditAboutCompanyKeynote - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditAboutCompanyKeynote - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.AboutCompanyKeynoteEntity{
		ID:             id,
		AboutCompanyID: req.AboutCompanyID,
		Keynote:        req.Keynote,
		PathImage:      req.PathImage,
	}

	err = a.AboutCompanyKeynoteService.EditByIDAboutCompanyKeynote(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditAboutCompanyKeynote - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success edit about company keynote"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllAboutCompanyKeynote implements CLientSectionHandlerInterface.
func (a *aboutCompanyKeynoteHandler) FetchAllAboutCompanyKeynote(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = []response.AboutCompanyKeynoteResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLAboutCompanyKeynote - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := a.AboutCompanyKeynoteService.FetchAllAboutCompanyKeynote(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLAboutCompanyKeynote - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClient = append(respClient, response.AboutCompanyKeynoteResponse{
			ID:                      val.ID,
			AboutCompanyID:          val.AboutCompanyID,
			Keynote:                 val.Keynote,
			PathImage:               val.PathImage,
			AboutCompanyDescription: val.AboutCompanyDescription,
		})
	}
	resp.Meta.Message = "Success fetch all about company keynote"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDAboutCompanyKeynote implements CLientSectionHandlerInterface.
func (a *aboutCompanyKeynoteHandler) FetchByIDAboutCompanyKeynote(c echo.Context) error {
	var (
		resp                    = response.DefaultSuccessResponse{}
		respError               = response.ErrorResponseDefault{}
		ctx                     = c.Request().Context()
		respAboutCompanyKeynote = response.AboutCompanyKeynoteResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID AboutCompanyKeynote - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idCompanyKeynote := c.Param("id")
	id, err := conv.StringToInt64(idCompanyKeynote)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID AboutCompanyKeynote - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := a.AboutCompanyKeynoteService.FetchByIDAboutCompanyKeynote(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID About Company - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respAboutCompanyKeynote.ID = result.ID
	respAboutCompanyKeynote.AboutCompanyID = result.AboutCompanyID
	respAboutCompanyKeynote.Keynote = result.Keynote
	respAboutCompanyKeynote.PathImage = result.PathImage
	respAboutCompanyKeynote.AboutCompanyDescription = result.AboutCompanyDescription
	resp.Meta.Message = "Success fetch about company keynote by ID"
	resp.Meta.Status = true
	resp.Data = respAboutCompanyKeynote
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewAboutCompanyKeynoteHandler(e *echo.Echo, AboutCompanyKeynoteService service.AboutCompanyKeynoteServiceInterface, cfg *config.Config) AboutCompanyKeynoteHandlerInterface {
	h := &aboutCompanyKeynoteHandler{
		AboutCompanyKeynoteService: AboutCompanyKeynoteService,
	}

	mid := middleware.NewMiddleware(cfg)
	AboutCompanyKeynoteApp := e.Group("/about-company-keynotes")
	adminApp := AboutCompanyKeynoteApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", h.CreateAboutCompanyKeynote)
	adminApp.GET("", h.FetchAllAboutCompanyKeynote)
	adminApp.GET("/:id", h.FetchByIDAboutCompanyKeynote)
	adminApp.GET("/:company_id", nil)
	adminApp.PUT("/:id", h.EditByIDAboutCompanyKeynote)
	adminApp.DELETE("/:id", h.DeleteByIDAboutCompanyKeynote)
	return h
}
