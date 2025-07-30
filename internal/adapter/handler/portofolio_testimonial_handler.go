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

type PortofolioTestimonialHandlerInterface interface {
	CreatePortofolioTestimonial(c echo.Context) error
	FetchAllPortofolioTestimonial(c echo.Context) error
	FetchByIDPortofolioTestimonial(c echo.Context) error
	EditByIDPortofolioTestimonial(c echo.Context) error
	DeleteByIDPortofolioTestimonial(c echo.Context) error
}
type portofolioTestimonialHandler struct {
	PortofolioTestimonialService service.PortofolioTestimonialServiceInterface
}

// CreatePortofolioTestimonial implements CLientSectionHandlerInterface.
func (p *portofolioTestimonialHandler) CreatePortofolioTestimonial(c echo.Context) error {
	var (
		req       = request.PortofolioTestimonialRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreatePortofolioTestimonial - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioTestimonial - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioTestimonial - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.PortofolioTestimonialEntity{
		Thumbnail:         req.Thumbnail,
		Message:           req.Message,
		ClientName:        req.ClientName,
		Role:              req.Role,
		PortoFolioSection: entity.PortofolioSectionEntity{ID: req.PortoFolioSectionID},
	}

	err = p.PortofolioTestimonialService.CreatePortofolioTestimonial(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreatePortofolioTestimonial - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success create portofolio testimonial"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDPortofolioTestimonial implements CLientSectionHandlerInterface.
func (p *portofolioTestimonialHandler) DeleteByIDPortofolioTestimonial(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeletePortofolioTestimonial - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioTestimonial := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioTestimonial)
	if err != nil {
		log.Errorf("[HANDLER] DeletePortofolioTestimonial - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = p.PortofolioTestimonialService.DeleteByIDPortofolioTestimonial(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeletePortofolioTestimonial - 3: %v", err)
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

// EditByIDPortofolioTestimonial implements CLientSectionHandlerInterface.
func (p *portofolioTestimonialHandler) EditByIDPortofolioTestimonial(c echo.Context) error {
	var (
		req       = request.PortofolioTestimonialRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditPortofolioTestimonial - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioTestimonial := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioTestimonial)
	if err != nil {
		log.Errorf("[HANDLER] EditPortofolioTestimonial - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditPortofolioTestimonial - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditPortofolioTestimonial - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.PortofolioTestimonialEntity{
		ID:         id,
		Thumbnail:  req.Thumbnail,
		Message:    req.Message,
		ClientName: req.ClientName,
		Role:       req.Role,
		PortoFolioSection: entity.PortofolioSectionEntity{
			ID: req.PortoFolioSectionID,
		},
	}

	err = p.PortofolioTestimonialService.EditByIDPortofolioTestimonial(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditPortofolioTestimonial - 6: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success edit portofolio testimonial"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllPortofolioTestimonial implements CLientSectionHandlerInterface.
func (p *portofolioTestimonialHandler) FetchAllPortofolioTestimonial(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = []response.PortofolioTestimonialResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLPortofolioTestimonial - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := p.PortofolioTestimonialService.FetchAllPortofolioTestimonial(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLPortofolioTestimonial - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClient = append(respClient, response.PortofolioTestimonialResponse{
			ID:                val.ID,
			Thumbnail:         val.Thumbnail,
			Message:           val.Message,
			ClientName:        val.ClientName,
			Role:              val.Role,
			PortoFolioSection: response.PortofolioSectionResponse{Name: val.PortoFolioSection.Name},
		})
	}
	resp.Meta.Message = "Success fetch all portofolio testimonials"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDPortofolioTestimonial implements CLientSectionHandlerInterface.
func (p *portofolioTestimonialHandler) FetchByIDPortofolioTestimonial(c echo.Context) error {
	var (
		resp                      = response.DefaultSuccessResponse{}
		respError                 = response.ErrorResponseDefault{}
		ctx                       = c.Request().Context()
		respPortofolioTestimonial = response.PortofolioTestimonialResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID PortofolioTestimonial - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioTestimonial := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioTestimonial)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID PortofolioTestimonial - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := p.PortofolioTestimonialService.FetchByIDPortofolioTestimonial(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID Service Section - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respPortofolioTestimonial.ID = result.ID
	respPortofolioTestimonial.Thumbnail = result.Thumbnail
	respPortofolioTestimonial.Message = result.Message
	respPortofolioTestimonial.ClientName = result.ClientName
	respPortofolioTestimonial.Role = result.Role
	respPortofolioTestimonial.PortoFolioSection.ID = result.PortoFolioSection.ID
	respPortofolioTestimonial.PortoFolioSection.Name = result.PortoFolioSection.Name
	respPortofolioTestimonial.PortoFolioSection.Thumbnail = result.PortoFolioSection.Thumbnail
	resp.Meta.Message = "Success fetch portofolio testimonial by ID"
	resp.Meta.Status = true
	resp.Data = respPortofolioTestimonial
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewPortofolioTestimonialHandler(e *echo.Echo, PortofolioTestimonialService service.PortofolioTestimonialServiceInterface, cfg *config.Config) PortofolioTestimonialHandlerInterface {
	h := &portofolioTestimonialHandler{
		PortofolioTestimonialService: PortofolioTestimonialService,
	}

	mid := middleware.NewMiddleware(cfg)
	portofolioSectionApp := e.Group("/portofolio-testimonial")
	adminApp := portofolioSectionApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", h.CreatePortofolioTestimonial)
	adminApp.GET("", h.FetchAllPortofolioTestimonial)
	adminApp.GET("/:id", h.FetchByIDPortofolioTestimonial)
	adminApp.PUT("/:id", h.EditByIDPortofolioTestimonial)
	adminApp.DELETE("/:id", h.DeleteByIDPortofolioTestimonial)
	return h
}
