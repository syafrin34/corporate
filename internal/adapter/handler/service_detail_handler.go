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

type ServiceDetailHandlerInterface interface {
	CreateServiceDetail(c echo.Context) error
	FetchAllServiceDetail(c echo.Context) error
	FetchByIDServiceDetail(c echo.Context) error
	EditByIDServiceDetail(c echo.Context) error
	DeleteByIDServiceDetail(c echo.Context) error
	FetchServiceDetailByServiceID(c echo.Context) error
}
type serviceDetailHandler struct {
	serviceDetailService service.ServiceDetailServiceInterface
}

func (s *serviceDetailHandler) FetchServiceDetailByServiceID(c echo.Context) error {
	var (
		resp              = response.DefaultSuccessResponse{}
		respError         = response.ErrorResponseDefault{}
		ctx               = c.Request().Context()
		respServiceDetail = response.ServiceDetailResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateServiceDetail - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}
	idServiceDetail := c.Param("id")
	id, err := conv.StringToInt64(idServiceDetail)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID ServiceDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := s.serviceDetailService.GetByServiceIDDetail(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID ServiceDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	respServiceDetail.ID = result.ID
	respServiceDetail.ServiceName = result.ServiceName
	respServiceDetail.PathImage = result.PathImage
	respServiceDetail.PathPdf = result.PathPdf
	respServiceDetail.Title = result.Title
	respServiceDetail.Description = result.Description
	respServiceDetail.ServiceID = result.ServiceID

	resp.Meta.Message = "Success fetch service detail by service"
	resp.Meta.Status = true
	resp.Data = respServiceDetail
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)

}

// CreateServiceDetail implements CLientSectionHandlerInterface.
func (s *serviceDetailHandler) CreateServiceDetail(c echo.Context) error {
	var (
		req       = request.ServiceDetailRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateServiceDetail - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateServiceDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateServiceDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ServiceDetailEntity{
		ServiceID:   req.ServiceID,
		PathImage:   req.PathImage,
		PathPdf:     req.PathPdf,
		PathDocx:    req.PathDocx,
		Title:       req.Title,
		Description: req.Description,
		ServiceName: req.ServiceName,
	}

	err = s.serviceDetailService.CreateServiceDetail(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateServiceDetail - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success create service detail"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDServiceDetail implements CLientSectionHandlerInterface.
func (s *serviceDetailHandler) DeleteByIDServiceDetail(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteServiceDetail - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idService := c.Param("id")
	id, err := conv.StringToInt64(idService)
	if err != nil {
		log.Errorf("[HANDLER] DeleteServiceDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = s.serviceDetailService.DeleteByIDServiceDetail(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteServiceDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success delete service detail"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDServiceDetail implements CLientSectionHandlerInterface.
func (s *serviceDetailHandler) EditByIDServiceDetail(c echo.Context) error {
	var (
		req       = request.ServiceDetailRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditServiceDetail - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idServiceDetail := c.Param("id")
	id, err := conv.StringToInt64(idServiceDetail)
	if err != nil {
		log.Errorf("[HANDLER] EditServiceDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditServiceDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditServiceDetail - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ServiceDetailEntity{
		ID:          id,
		ServiceID:   req.ServiceID,
		PathImage:   req.PathImage,
		PathPdf:     req.PathPdf,
		PathDocx:    req.PathDocx,
		Title:       req.Title,
		Description: req.Description,
		ServiceName: req.ServiceName,
	}

	err = s.serviceDetailService.EditByIDServiceDetail(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditServiceDetail - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success edit service detail"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllServiceDetail implements CLientSectionHandlerInterface.
func (cs *serviceDetailHandler) FetchAllServiceDetail(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = []response.ServiceDetailResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLServiceDetail - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.serviceDetailService.FetchAllServiceDetail(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLServiceDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClient = append(respClient, response.ServiceDetailResponse{
			ID:          val.ID,
			ServiceID:   val.ServiceID,
			PathImage:   val.PathImage,
			PathPdf:     val.PathPdf,
			PathDocx:    val.PathDocx,
			Title:       val.Title,
			Description: val.Description,
			ServiceName: val.ServiceName,
		})
	}
	resp.Meta.Message = "Success fetch all service detail"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDServiceDetail implements CLientSectionHandlerInterface.
func (cs *serviceDetailHandler) FetchByIDServiceDetail(c echo.Context) error {
	var (
		resp              = response.DefaultSuccessResponse{}
		respError         = response.ErrorResponseDefault{}
		ctx               = c.Request().Context()
		respServiceDetail = response.ServiceDetailResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID ServiceDetail - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idService := c.Param("id")
	id, err := conv.StringToInt64(idService)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID ServiceDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.serviceDetailService.FetchByIDServiceDetail(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID Service Section - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respServiceDetail.ID = result.ID
	respServiceDetail.ServiceID = result.ServiceID
	respServiceDetail.PathImage = result.PathImage
	respServiceDetail.PathDocx = result.PathDocx
	respServiceDetail.PathPdf = result.PathPdf
	respServiceDetail.Title = result.Title
	respServiceDetail.Description = result.Description
	respServiceDetail.ServiceName = result.ServiceName

	resp.Meta.Message = "Success fetch service detail by ID"
	resp.Meta.Status = true
	resp.Data = respServiceDetail
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewServiceDetailHandler(e *echo.Echo, ServiceDetailService service.ServiceDetailServiceInterface, cfg *config.Config) ServiceDetailHandlerInterface {
	h := &serviceDetailHandler{
		serviceDetailService: ServiceDetailService,
	}

	mid := middleware.NewMiddleware(cfg)
	ServiceDetailApp := e.Group("/service-detail")
	ServiceDetailApp.GET("", h.FetchServiceDetailByServiceID)
	adminApp := ServiceDetailApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", h.CreateServiceDetail)
	adminApp.GET("", h.FetchAllServiceDetail)
	adminApp.GET("/:id", h.FetchByIDServiceDetail)
	adminApp.PUT("/:id", h.EditByIDServiceDetail)
	adminApp.DELETE("/:id", h.DeleteByIDServiceDetail)
	return h
}
