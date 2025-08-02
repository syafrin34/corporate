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

type FaqSectionHandlerInterface interface {
	CreateFaqSection(c echo.Context) error
	FetchAllFaqSection(c echo.Context) error
	FetchByIDFaqSection(c echo.Context) error
	EditByIDFaqSection(c echo.Context) error
	DeleteByIDFaqSection(c echo.Context) error
	FetchAllFaqSectionHome(c echo.Context) error
}
type faqSectionHandler struct {
	FaqSectionService service.FaqSectionServiceInterface
}

// CreateFaqSection implements CLientSectionHandlerInterface.
func (cs *faqSectionHandler) CreateFaqSection(c echo.Context) error {
	var (
		req       = request.FaqSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateFaqSection - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateFaqSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateFaqSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.FaqSectionEntity{
		Description: req.Description,
	}

	err = cs.FaqSectionService.CreateFaqSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateFaqSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success create faq section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

func (cs *faqSectionHandler) FetchAllFaqSectionHome(c echo.Context) error {
	var (
		respFaqs  = []response.FaqSectionResponse{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	results, err := cs.FaqSectionService.FetchAllFaqSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] DeleteFaqSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	for _, val := range results {
		respFaqs = append(respFaqs, response.FaqSectionResponse{
			ID:          val.ID,
			Title:       val.Title,
			Description: val.Description,
		})
	}
	resp.Meta.Message = "Success fetch All faq section home"
	resp.Meta.Status = true
	resp.Data = respFaqs
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)

}

// DeleteByIDFaqSection implements CLientSectionHandlerInterface.
func (cs *faqSectionHandler) DeleteByIDFaqSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteFaqSection - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idCompany := c.Param("id")
	id, err := conv.StringToInt64(idCompany)
	if err != nil {
		log.Errorf("[HANDLER] DeleteFaqSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.FaqSectionService.DeleteByIDFaqSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteFaqSection - 3: %v", err)
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

// EditByIDFaqSection implements CLientSectionHandlerInterface.
func (cs *faqSectionHandler) EditByIDFaqSection(c echo.Context) error {
	var (
		req       = request.FaqSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditFaqSection - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idFaqSection := c.Param("id")
	id, err := conv.StringToInt64(idFaqSection)
	if err != nil {
		log.Errorf("[HANDLER] EditFaqSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditFaqSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditFaqSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.FaqSectionEntity{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
	}

	err = cs.FaqSectionService.EditByIDFaqSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditFaqSection - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success edit faq section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllFaqSection implements CLientSectionHandlerInterface.
func (cs *faqSectionHandler) FetchAllFaqSection(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = []response.FaqSectionResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLFaqSection - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.FaqSectionService.FetchAllFaqSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLFaqSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClient = append(respClient, response.FaqSectionResponse{
			ID:          val.ID,
			Title:       val.Title,
			Description: val.Description,
		})
	}
	resp.Meta.Message = "Success fetch faq section"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDFaqSection implements CLientSectionHandlerInterface.
func (cs *faqSectionHandler) FetchByIDFaqSection(c echo.Context) error {
	var (
		resp           = response.DefaultSuccessResponse{}
		respError      = response.ErrorResponseDefault{}
		ctx            = c.Request().Context()
		respFaqSection = response.FaqSectionResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID FaqSection - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idFaqSection := c.Param("id")
	id, err := conv.StringToInt64(idFaqSection)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID FaqSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.FaqSectionService.FetchByIDFaqSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID Faq Section - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respFaqSection.ID = result.ID
	respFaqSection.Description = result.Description
	resp.Meta.Message = "Success fetch faq section by ID"
	resp.Meta.Status = true
	resp.Data = respFaqSection
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewFaqSectionHandler(e *echo.Echo, FaqSectionService service.FaqSectionServiceInterface, cfg *config.Config) FaqSectionHandlerInterface {
	h := &faqSectionHandler{
		FaqSectionService: FaqSectionService,
	}

	mid := middleware.NewMiddleware(cfg)
	aboutCompanyApp := e.Group("/faq-sections")
	aboutCompanyApp.GET("", h.FetchAllFaqSectionHome)
	adminApp := aboutCompanyApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", h.CreateFaqSection)
	adminApp.GET("", h.FetchAllFaqSection)
	adminApp.GET("/:id", h.FetchByIDFaqSection)
	adminApp.PUT("/:id", h.EditByIDFaqSection)
	adminApp.DELETE("/:id", h.DeleteByIDFaqSection)
	return h
}
