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

type OurTeamHandlerInterface interface {
	CreateOurTeam(c echo.Context) error
	FetchAllOurTeam(c echo.Context) error
	FetchByIDOurTeam(c echo.Context) error
	EditByIDOurTeam(c echo.Context) error
	DeleteByIDOurTeam(c echo.Context) error
}
type ourTeamHandler struct {
	OurTeamService service.OurTeamServiceInterface
}

// CreateOurTeam implements CLientSectionHandlerInterface.
func (o *ourTeamHandler) CreateOurTeam(c echo.Context) error {
	var (
		req       = request.OurTeamRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateOurTeam - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateOurTeam - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateOurTeam - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.OurTeamEntity{
		Name:      req.Name,
		Role:      req.Role,
		PathPhoto: req.PathPhoto,
		Tagline:   req.Tagline,
	}

	err = o.OurTeamService.CreateOurTeam(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateOurTeam - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success create our team"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDOurTeam implements CLientSectionHandlerInterface.
func (o *ourTeamHandler) DeleteByIDOurTeam(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteOurTeam - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idCompany := c.Param("id")
	id, err := conv.StringToInt64(idCompany)
	if err != nil {
		log.Errorf("[HANDLER] DeleteOurTeam - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = o.OurTeamService.DeleteByIDOurTeam(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteOurTeam - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success delete our team"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDOurTeam implements CLientSectionHandlerInterface.
func (o *ourTeamHandler) EditByIDOurTeam(c echo.Context) error {
	var (
		req       = request.OurTeamRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditOurTeam - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idOurTeam := c.Param("id")
	id, err := conv.StringToInt64(idOurTeam)
	if err != nil {
		log.Errorf("[HANDLER] EditOurTeam - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditOurTeam - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditOurTeam - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.OurTeamEntity{
		ID:        id,
		Name:      req.Name,
		Role:      req.Role,
		PathPhoto: req.PathPhoto,
		Tagline:   req.Tagline,
	}

	err = o.OurTeamService.EditByIDOurTeam(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditOurTeam - 5: %v", err)
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

// FetchAllOurTeam implements CLientSectionHandlerInterface.
func (cs *ourTeamHandler) FetchAllOurTeam(c echo.Context) error {
	var (
		resp        = response.DefaultSuccessResponse{}
		respError   = response.ErrorResponseDefault{}
		ctx         = c.Request().Context()
		respOurTeam = []response.OurTeamResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLOurTeam - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.OurTeamService.FetchAllOurTeam(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLOurTeam - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respOurTeam = append(respOurTeam, response.OurTeamResponse{
			ID:        val.ID,
			Name:      val.Name,
			Role:      val.Role,
			PathPhoto: val.PathPhoto,
			Tagline:   val.Tagline,
		})
	}
	resp.Meta.Message = "Success fetch all our team"
	resp.Meta.Status = true
	resp.Data = respOurTeam
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDOurTeam implements CLientSectionHandlerInterface.
func (cs *ourTeamHandler) FetchByIDOurTeam(c echo.Context) error {
	var (
		resp        = response.DefaultSuccessResponse{}
		respError   = response.ErrorResponseDefault{}
		ctx         = c.Request().Context()
		respOurTeam = response.OurTeamResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID OurTeam - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idOurTeam := c.Param("id")
	id, err := conv.StringToInt64(idOurTeam)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID OurTeam - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.OurTeamService.FetchByIDOurTeam(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID Our Team - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respOurTeam.ID = result.ID
	respOurTeam.Name = result.Name
	respOurTeam.Role = result.Role
	respOurTeam.PathPhoto = result.PathPhoto
	respOurTeam.Tagline = result.Tagline
	resp.Meta.Message = "Success fetch our team by ID"
	resp.Meta.Status = true
	resp.Data = respOurTeam
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewOurTeamHandler(e *echo.Echo, OurTeamService service.OurTeamServiceInterface, cfg *config.Config) OurTeamHandlerInterface {
	h := &ourTeamHandler{
		OurTeamService: OurTeamService,
	}

	mid := middleware.NewMiddleware(cfg)
	ourTeamApp := e.Group("/our-teams")
	adminApp := ourTeamApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", h.CreateOurTeam)
	adminApp.GET("", h.FetchAllOurTeam)
	adminApp.GET("/:id", h.FetchByIDOurTeam)
	adminApp.PUT("/:id", h.EditByIDOurTeam)
	adminApp.DELETE("/:id", h.DeleteByIDOurTeam)
	return h
}
