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
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PortofolioDetailHandlerInterface interface {
	CreatePortofolioDetail(c echo.Context) error
	FetchAllPortofolioDetail(c echo.Context) error
	FetchByIDPortofolioDetail(c echo.Context) error
	EditByIDPortofolioDetail(c echo.Context) error
	DeleteByIDPortofolioDetail(c echo.Context) error
	FetchDetailPortofolioByPortoID(c echo.Context) error
}
type portofolioDetailHandler struct {
	PortofolioDetailService service.PortofolioDetailServiceInterface
}

func (p *portofolioDetailHandler) FetchDetailPortofolioByPortoID(c echo.Context) error {

	var (
		respDetail = response.PortofolioDetailResponse{}
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
	)
	idPorto := c.Param("id")
	id, err := conv.StringToInt64(idPorto)
	if err != nil {
		log.Errorf("[HANDLER] DeletePortofolioDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}
	result, err := p.PortofolioDetailService.FetchDetailPortofolioByPortoID(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeletePortofolioDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respDetail.ID = result.ID
	respDetail.Category = result.Category
	respDetail.ClientName = result.ClientName
	respDetail.ProjectDate = result.ProjectDate.Format("02 January 2006")
	respDetail.ProjectUrl = result.ProjectUrl
	respDetail.Title = result.Title
	respDetail.Description = result.Description
	respDetail.PortofolioSection.ID = result.PortofolioSection.ID
	respDetail.PortofolioSection.Name = result.PortofolioSection.Name
	respDetail.PortofolioSection.Thumbnail = result.PortofolioSection.Thumbnail
	respDetail.PortofolioSection.ID = result.PortofolioSection.ID
	resp.Meta.Message = "Success Fetch portofolio detail"
	resp.Meta.Status = true
	resp.Data = respDetail
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)

}

// CreatePortofolioDetail implements CLientSectionHandlerInterface.
func (p *portofolioDetailHandler) CreatePortofolioDetail(c echo.Context) error {
	var (
		req       = request.PortofolioDetailRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreatePortofolioDetail - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	stringprojectDate, err := time.Parse("2006-01-02", req.ProjectDate)
	if err != nil {
		log.Errorf("[HANDLER] CreatePortofolioDetail - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}
	reqEntity := entity.PortofolioDetailEntity{
		Category:    req.Category,
		ClientName:  req.ClientName,
		ProjectDate: stringprojectDate,
		ProjectUrl:  req.ProjectUrl,
		Title:       req.Title,
		Description: req.Description,
		PortofolioSection: entity.PortofolioSectionEntity{
			ID: req.PortofolioSectionID,
		},
	}

	err = p.PortofolioDetailService.CreatePortofolioDetail(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreatePortofolioDetail - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success create portofolio detail"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDPortofolioDetail implements CLientSectionHandlerInterface.
func (p *portofolioDetailHandler) DeleteByIDPortofolioDetail(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeletePortofolioDetail - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioDetail := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioDetail)
	if err != nil {
		log.Errorf("[HANDLER] DeletePortofolioDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = p.PortofolioDetailService.DeleteByIDPortofolioDetail(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeletePortofolioDetail - 3: %v", err)
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

// EditByIDPortofolioDetail implements CLientSectionHandlerInterface.
func (p *portofolioDetailHandler) EditByIDPortofolioDetail(c echo.Context) error {
	var (
		req       = request.PortofolioDetailRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditPortofolioDetail - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioDetail := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioDetail)
	if err != nil {
		log.Errorf("[HANDLER] EditPortofolioDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditPortofolioDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditPortofolioDetail - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	stringprojectDate, err := time.Parse("2006-01-02", req.ProjectDate)
	if err != nil {
		log.Errorf("[HANDLER] EditPortofolioDetail - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}
	reqEntity := entity.PortofolioDetailEntity{
		ID:          id,
		Category:    req.Category,
		ClientName:  req.ClientName,
		ProjectDate: stringprojectDate,
		ProjectUrl:  req.ProjectUrl,
		Title:       req.Title,
		Description: req.Description,
		PortofolioSection: entity.PortofolioSectionEntity{
			ID: req.PortofolioSectionID,
		},
	}

	err = p.PortofolioDetailService.EditByIDPortofolioDetail(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditPortofolioDetail - 6: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success edit portofolio detail"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllPortofolioDetail implements CLientSectionHandlerInterface.
func (p *portofolioDetailHandler) FetchAllPortofolioDetail(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = []response.PortofolioDetailResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLPortofolioDetail - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := p.PortofolioDetailService.FetchAllPortofolioDetail(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLPortofolioDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClient = append(respClient, response.PortofolioDetailResponse{
			ID:          val.ID,
			Category:    val.Category,
			ClientName:  val.ClientName,
			ProjectDate: val.ProjectDate.Format("02 Jauary 2006"),
			ProjectUrl:  val.ProjectUrl,
			Title:       val.Title,
			Description: val.Description,
			PortofolioSection: response.PortofolioSectionResponse{
				ID:        val.PortofolioSection.ID,
				Name:      val.PortofolioSection.Name,
				Thumbnail: val.PortofolioSection.Thumbnail,
			},
		})
	}
	resp.Meta.Message = "Success fetch all portofolio detail"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDPortofolioDetail implements CLientSectionHandlerInterface.
func (p *portofolioDetailHandler) FetchByIDPortofolioDetail(c echo.Context) error {
	var (
		resp                 = response.DefaultSuccessResponse{}
		respError            = response.ErrorResponseDefault{}
		ctx                  = c.Request().Context()
		respPortofolioDetail = response.PortofolioDetailResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID PortofolioDetail - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioDetail := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioDetail)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID PortofolioDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := p.PortofolioDetailService.FetchByIDPortofolioDetail(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID Service Section - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respPortofolioDetail.ID = result.ID
	respPortofolioDetail.Category = result.Category
	respPortofolioDetail.ProjectDate = result.ProjectDate.Format("02 Januari 2006")
	respPortofolioDetail.ProjectUrl = result.ProjectUrl
	respPortofolioDetail.Title = result.Title
	respPortofolioDetail.Description = result.Description
	respPortofolioDetail.PortofolioSection.ID = result.PortofolioSection.ID
	respPortofolioDetail.PortofolioSection.Name = result.PortofolioSection.Name
	respPortofolioDetail.PortofolioSection.Thumbnail = result.PortofolioSection.Thumbnail
	resp.Meta.Message = "Success fetch portofolio detail by ID"
	resp.Meta.Status = true
	resp.Data = respPortofolioDetail
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewPortofolioDetailHandler(e *echo.Echo, PortofolioDetailService service.PortofolioDetailServiceInterface, cfg *config.Config) PortofolioDetailHandlerInterface {
	h := &portofolioDetailHandler{
		PortofolioDetailService: PortofolioDetailService,
	}

	mid := middleware.NewMiddleware(cfg)
	portofolioDetailApp := e.Group("/portofolio-details")
	portofolioDetailApp.GET("/:id", h.FetchDetailPortofolioByPortoID)
	adminApp := portofolioDetailApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", h.CreatePortofolioDetail)
	adminApp.GET("", h.FetchAllPortofolioDetail)
	adminApp.GET("/:id", h.FetchByIDPortofolioDetail)
	adminApp.PUT("/:id", h.EditByIDPortofolioDetail)
	adminApp.DELETE("/:id", h.DeleteByIDPortofolioDetail)
	return h
}
