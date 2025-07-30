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

type PortofolioSectionHandlerInterface interface {
	CreatePortofolioSection(c echo.Context) error
	FetchAllPortofolioSection(c echo.Context) error
	FetchByIDPortofolioSection(c echo.Context) error
	EditByIDPortofolioSection(c echo.Context) error
	DeleteByIDPortofolioSection(c echo.Context) error
}
type portofolioSectionHandler struct {
	PortofolioSectionService service.PortofolioSectionServiceInterface
}

// CreatePortofolioSection implements CLientSectionHandlerInterface.
func (p *portofolioSectionHandler) CreatePortofolioSection(c echo.Context) error {
	var (
		req       = request.PortofolioSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreatePortofolioSection - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.PortofolioSectionEntity{
		Name:      req.Name,
		Thumbnail: req.Thumbnail,
		Tagline:   req.Tagline,
	}

	err = p.PortofolioSectionService.CreatePortofolioSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreatePortofolioSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success create portofolio section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDPortofolioSection implements CLientSectionHandlerInterface.
func (p *portofolioSectionHandler) DeleteByIDPortofolioSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeletePortofolioSection - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolio := c.Param("id")
	id, err := conv.StringToInt64(idPortofolio)
	if err != nil {
		log.Errorf("[HANDLER] DeletePortofolioSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = p.PortofolioSectionService.DeleteByIDPortofolioSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeletePortofolioSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success delete portofolio section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDPortofolioSection implements CLientSectionHandlerInterface.
func (p *portofolioSectionHandler) EditByIDPortofolioSection(c echo.Context) error {
	var (
		req       = request.PortofolioSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditPortofolioSection - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolio := c.Param("id")
	id, err := conv.StringToInt64(idPortofolio)
	if err != nil {
		log.Errorf("[HANDLER] EditPortofolioSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditPortofolioSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditPortofolioSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.PortofolioSectionEntity{
		ID:        id,
		Name:      req.Name,
		Thumbnail: req.Thumbnail,
		Tagline:   req.Tagline,
	}

	err = p.PortofolioSectionService.EditByIDPortofolioSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditPortofolioSection - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success edit portofolio section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllPortofolioSection implements CLientSectionHandlerInterface.
func (p *portofolioSectionHandler) FetchAllPortofolioSection(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = []response.PortofolioSectionResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLPortofolioSection - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := p.PortofolioSectionService.FetchAllPortofolioSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLPortofolioSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClient = append(respClient, response.PortofolioSectionResponse{
			ID:        val.ID,
			Name:      val.Name,
			Thumbnail: val.Thumbnail,
			Tagline:   val.Tagline,
		})
	}
	resp.Meta.Message = "Success fetch all portofolio section"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDPortofolioSection implements CLientSectionHandlerInterface.
func (p *portofolioSectionHandler) FetchByIDPortofolioSection(c echo.Context) error {
	var (
		resp                  = response.DefaultSuccessResponse{}
		respError             = response.ErrorResponseDefault{}
		ctx                   = c.Request().Context()
		respPortofolioSection = response.PortofolioSectionResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID PortofolioSection - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolio := c.Param("id")
	id, err := conv.StringToInt64(idPortofolio)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID PortofolioSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := p.PortofolioSectionService.FetchByIDPortofolioSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID Service Section - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respPortofolioSection.ID = result.ID
	respPortofolioSection.Name = result.Name
	respPortofolioSection.Thumbnail = result.Thumbnail
	respPortofolioSection.Tagline = result.Tagline
	resp.Meta.Message = "Success fetch portofolio section by ID"
	resp.Meta.Status = true
	resp.Data = respPortofolioSection
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewPortofolioSectionHandler(e *echo.Echo, PortofolioSectionService service.PortofolioSectionServiceInterface, cfg *config.Config) PortofolioSectionHandlerInterface {
	h := &portofolioSectionHandler{
		PortofolioSectionService: PortofolioSectionService,
	}

	mid := middleware.NewMiddleware(cfg)
	portofolioSectionApp := e.Group("/portofolio-sections")
	adminApp := portofolioSectionApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", h.CreatePortofolioSection)
	adminApp.GET("", h.FetchAllPortofolioSection)
	adminApp.GET("/:id", h.FetchByIDPortofolioSection)
	adminApp.PUT("/:id", h.EditByIDPortofolioSection)
	adminApp.DELETE("/:id", h.DeleteByIDPortofolioSection)
	return h
}
