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

type ServiceSectionHandlerInterface interface {
	CreateServiceSection(c echo.Context) error
	FetchAllServiceSection(c echo.Context) error
	FetchByIDServiceSection(c echo.Context) error
	EditByIDServiceSection(c echo.Context) error
	DeleteByIDServiceSection(c echo.Context) error
	FetchAllServiceHome(c echo.Context) error
}
type serviceSectionHandler struct {
	ServiceSectionService service.ServiceSectionServiceInterface
}

func (s *serviceSectionHandler) FetchAllServiceHome(c echo.Context) error {
	var (
		respServices = []response.ServiceSectionResponse{}
		resp         = response.DefaultSuccessResponse{}
		respError    = response.ErrorResponseDefault{}
		ctx          = c.Request().Context()
	)

	results, err := s.ServiceSectionService.FetchAllServiceSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllServiceHome - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	for _, val := range results {
		respServices = append(respServices, response.ServiceSectionResponse{
			ID:       val.ID,
			Name:     val.Name,
			Tagline:  val.Tagline,
			PathIcon: val.PathIcon,
		})
	}

	resp.Data = respServices
	resp.Meta.Message = "Success fetch all service home"
	resp.Meta.Status = true
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// CreateServiceSection implements CLientSectionHandlerInterface.
func (s *serviceSectionHandler) CreateServiceSection(c echo.Context) error {
	var (
		req       = request.ServiceSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateServiceSection - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateServiceSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateServiceSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ServiceSectionEntity{
		Name:     req.Name,
		PathIcon: req.PathIcon,
		Tagline:  req.Tagline,
	}

	err = s.ServiceSectionService.CreateServiceSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateServiceSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success create service section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDServiceSection implements CLientSectionHandlerInterface.
func (s *serviceSectionHandler) DeleteByIDServiceSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteServiceSection - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idService := c.Param("id")
	id, err := conv.StringToInt64(idService)
	if err != nil {
		log.Errorf("[HANDLER] DeleteServiceSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = s.ServiceSectionService.DeleteByIDServiceSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteServiceSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success delete service section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDServiceSection implements CLientSectionHandlerInterface.
func (s *serviceSectionHandler) EditByIDServiceSection(c echo.Context) error {
	var (
		req       = request.ServiceSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditServiceSection - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idService := c.Param("id")
	id, err := conv.StringToInt64(idService)
	if err != nil {
		log.Errorf("[HANDLER] EditServiceSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditServiceSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditServiceSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ServiceSectionEntity{
		ID:       id,
		Name:     req.Name,
		PathIcon: req.PathIcon,
		Tagline:  req.Tagline,
	}

	err = s.ServiceSectionService.EditByIDServiceSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditServiceSection - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success edit service section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllServiceSection implements CLientSectionHandlerInterface.
func (cs *serviceSectionHandler) FetchAllServiceSection(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = []response.ServiceSectionResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLServiceSection - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.ServiceSectionService.FetchAllServiceSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLServiceSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClient = append(respClient, response.ServiceSectionResponse{
			ID:       val.ID,
			Name:     val.Name,
			PathIcon: val.Name,
			Tagline:  val.Tagline,
		})
	}
	resp.Meta.Message = "Success fetch all service section"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDServiceSection implements CLientSectionHandlerInterface.
func (cs *serviceSectionHandler) FetchByIDServiceSection(c echo.Context) error {
	var (
		resp               = response.DefaultSuccessResponse{}
		respError          = response.ErrorResponseDefault{}
		ctx                = c.Request().Context()
		respServiceSection = response.ServiceSectionResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID ServiceSection - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idService := c.Param("id")
	id, err := conv.StringToInt64(idService)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID ServiceSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.ServiceSectionService.FetchByIDServiceSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID Service Section - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respServiceSection.ID = result.ID
	respServiceSection.Name = result.Name
	respServiceSection.PathIcon = result.PathIcon
	respServiceSection.Tagline = result.Tagline
	resp.Meta.Message = "Success fetch service section by ID"
	resp.Meta.Status = true
	resp.Data = respServiceSection
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewServiceSectionHandler(e *echo.Echo, ServiceSectionService service.ServiceSectionServiceInterface, cfg *config.Config) ServiceSectionHandlerInterface {
	h := &serviceSectionHandler{
		ServiceSectionService: ServiceSectionService,
	}

	mid := middleware.NewMiddleware(cfg)
	serviceSectionApp := e.Group("/service-sections")
	serviceSectionApp.GET("", h.FetchAllServiceHome)
	adminApp := serviceSectionApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", h.CreateServiceSection)
	adminApp.GET("", h.FetchAllServiceSection)
	adminApp.GET("/:id", h.FetchByIDServiceSection)
	adminApp.PUT("/:id", h.EditByIDServiceSection)
	adminApp.DELETE("/:id", h.DeleteByIDServiceSection)
	return h
}
