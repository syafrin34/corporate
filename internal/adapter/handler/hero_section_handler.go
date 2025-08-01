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

type HeroSectionHandlerInterface interface {
	CreateHeroSection(c echo.Context) error
	FetchAllHeroSection(c echo.Context) error
	FetchByIDHeroSection(c echo.Context) error
	EditByIDHeroSection(c echo.Context) error
	DeleteByIDHeroSection(c echo.Context) error
	FetchHeroDataHome(c echo.Context) error
}

type heroSectionHandler struct {
	heroSectionService service.HeroSectionServiceInterface
}

// FetchHeroDataHome implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) FetchHeroDataHome(c echo.Context) error {
	var (
		respHero  = response.HeroSectionResponse{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	results, err := h.heroSectionService.FetchAllHeroSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLHeroDataHome - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respHero.Banner = results[0].Banner
	respHero.Heading = results[0].Heading
	respHero.SubHeading = results[0].SubHeading
	respHero.PathVideo = results[0].PathVideo
	respHero.ID = results[0].ID

	resp.Meta.Message = "Success Fetch hero data home"
	resp.Meta.Status = true
	resp.Data = respHero
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// CreateHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) CreateHeroSection(c echo.Context) error {
	var (
		req       = request.HeroSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateHeroSection - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.HeroSectionEntity{
		Heading:    req.Heading,
		SubHeading: req.SubHeading,
		PathVideo:  req.PathVideo,
		Banner:     req.Banner,
	}

	err = h.heroSectionService.CreateHeroSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success create hero section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) DeleteByIDHeroSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteHeroSection - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idHero := c.Param("id")
	id, err := conv.StringToInt64(idHero)
	if err != nil {
		log.Errorf("[HANDLER] DeleteHeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = h.heroSectionService.DeleteByIDHeroSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteHeroSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success delete hero section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) EditByIDHeroSection(c echo.Context) error {
	var (
		req       = request.HeroSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditHeroSection - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idHero := c.Param("id")
	id, err := conv.StringToInt64(idHero)
	if err != nil {
		log.Errorf("[HANDLER] EditHeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditHeroSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditHeroSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.HeroSectionEntity{
		ID:         id,
		Heading:    req.Heading,
		SubHeading: req.SubHeading,
		PathVideo:  req.PathVideo,
		Banner:     req.Banner,
	}

	err = h.heroSectionService.EditByIDHeroSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditHeroSection - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success dedit hero section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)

}

// FetchAllHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) FetchAllHeroSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
		respHero  = []response.HeroSectionResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLHeroSection - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := h.heroSectionService.FetchAllHeroSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLHeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respHero = append(respHero, response.HeroSectionResponse{
			ID:         val.ID,
			Heading:    val.Heading,
			SubHeading: val.SubHeading,
			PathVideo:  val.PathVideo,
			Banner:     val.Banner,
		})
	}
	resp.Meta.Message = "Success fetch all hero section"
	resp.Meta.Status = true
	resp.Data = respHero
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)

}

// FetchByIDHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) FetchByIDHeroSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
		respHero  = response.HeroSectionResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID HeroSection - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idHero := c.Param("id")
	id, err := conv.StringToInt64(idHero)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID HeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := h.heroSectionService.FetchByIDHeroSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID HeroSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respHero.ID = result.ID
	respHero.Heading = result.Heading
	respHero.SubHeading = result.SubHeading
	respHero.PathVideo = result.PathVideo
	respHero.Banner = result.Banner
	resp.Meta.Message = "Success fetch hero section by ID"
	resp.Meta.Status = true
	resp.Data = respHero
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)

}

func NewHeroSectionHandler(c *echo.Echo, cfg *config.Config, heroSectionService service.HeroSectionServiceInterface) HeroSectionHandlerInterface {

	heroHandler := &heroSectionHandler{
		heroSectionService: heroSectionService,
	}
	mid := middleware.NewMiddleware(cfg)
	heroApp := c.Group("/hero-sections")
	heroApp.GET("", heroHandler.FetchHeroDataHome)
	adminApp := heroApp.Group("/admin", mid.CheckToken())
	adminApp.GET("", heroHandler.FetchAllHeroSection)
	adminApp.POST("", heroHandler.CreateHeroSection)
	adminApp.GET("/:id", heroHandler.FetchByIDHeroSection)
	adminApp.PUT("/:id", heroHandler.EditByIDHeroSection)
	adminApp.DELETE("/:id", heroHandler.DeleteByIDHeroSection)
	return heroHandler
}
