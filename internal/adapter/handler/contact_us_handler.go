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

type ContactUsHandlerInterface interface {
	CreateContactUs(c echo.Context) error
	FetchAllContactUsHome(c echo.Context) error
	FetchAllContactUs(c echo.Context) error
	FetchByIDContactUs(c echo.Context) error
	EditByIDContactUs(c echo.Context) error
	DeleteByIDContactUs(c echo.Context) error
}
type contactUsHandler struct {
	contactUsService service.ContactUsServiceInterface
}

// CreateContactUs implements ContactUsHandlerInterface.
func (cs *contactUsHandler) CreateContactUs(c echo.Context) error {
	var (
		req       = request.ContactUsRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateContactUs - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateContactUs - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateContactUs - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ContactUsEntity{
		CompanyName:  req.CompanyName,
		LocationName: req.LocationName,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
	}

	err = cs.contactUsService.CreateContactUs(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateContactUs - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success create contact us"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDContactUs implements ContactUsHandlerInterface.
func (cs *contactUsHandler) DeleteByIDContactUs(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteContactUs - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idContactUs := c.Param("id")
	id, err := conv.StringToInt64(idContactUs)
	if err != nil {
		log.Errorf("[HANDLER] DeleteContactUs - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.contactUsService.DeleteByIDContactUs(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteContactUs - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success delete contact us"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDContactUs implements ContactUsHandlerInterface.
func (cs *contactUsHandler) EditByIDContactUs(c echo.Context) error {
	var (
		req       = request.ContactUsRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditContactUs - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idContactUs := c.Param("id")
	id, err := conv.StringToInt64(idContactUs)
	if err != nil {
		log.Errorf("[HANDLER] EditContactUs - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditContactUs - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditContactUs - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ContactUsEntity{
		ID:           id,
		CompanyName:  req.CompanyName,
		LocationName: req.LocationName,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
	}

	err = cs.contactUsService.EditByIDContactUs(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditContactUs - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success edit contact us"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllContactUs implements ContactUsHandlerInterface.
func (cs *contactUsHandler) FetchAllContactUsHome(c echo.Context) error {
	var (
		resp          = response.DefaultSuccessResponse{}
		respError     = response.ErrorResponseDefault{}
		ctx           = c.Request().Context()
		respContactUs = response.ContactUsResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLContactUs - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.contactUsService.FetchAllContactUs(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLContactUs - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respContactUs = response.ContactUsResponse{
		ID:           results[0].ID,
		CompanyName:  results[0].CompanyName,
		LocationName: results[0].LocationName,
		Address:      results[0].Address,
		PhoneNumber:  results[0].PhoneNumber,
	}

	resp.Meta.Message = "Success fetch all contact us home"
	resp.Meta.Status = true
	resp.Data = respContactUs
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllContactUs implements ContactUsHandlerInterface.
func (cs *contactUsHandler) FetchAllContactUs(c echo.Context) error {
	var (
		resp          = response.DefaultSuccessResponse{}
		respError     = response.ErrorResponseDefault{}
		ctx           = c.Request().Context()
		respContactUs = []response.ContactUsResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLContactUs - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.contactUsService.FetchAllContactUs(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLContactUs - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	for _, val := range results {
		respContactUs = append(respContactUs, response.ContactUsResponse{
			ID:           val.ID,
			CompanyName:  val.CompanyName,
			LocationName: val.LocationName,
			Address:      val.Address,
			PhoneNumber:  val.PhoneNumber,
		})
	}

	resp.Meta.Message = "Success fetch all contact us"
	resp.Meta.Status = true
	resp.Data = respContactUs
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDContactUs implements ContactUsHandlerInterface.
func (cs *contactUsHandler) FetchByIDContactUs(c echo.Context) error {
	var (
		resp          = response.DefaultSuccessResponse{}
		respError     = response.ErrorResponseDefault{}
		ctx           = c.Request().Context()
		respContactUs = response.ContactUsResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID ContactUs - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idContactUs := c.Param("id")
	id, err := conv.StringToInt64(idContactUs)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID ContactUs - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.contactUsService.FetchByIDContactUs(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID ContactUs - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respContactUs.ID = result.ID
	respContactUs.CompanyName = result.CompanyName
	respContactUs.LocationName = result.LocationName
	respContactUs.Address = result.Address
	respContactUs.PhoneNumber = result.PhoneNumber
	resp.Meta.Message = "Success fetch contact us section by ID"
	resp.Meta.Status = true
	resp.Data = respContactUs
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewContactUsHandler(e *echo.Echo, contactUsService service.ContactUsServiceInterface, cfg *config.Config) ContactUsHandlerInterface {
	h := &contactUsHandler{
		contactUsService: contactUsService,
	}

	mid := middleware.NewMiddleware(cfg)
	contactUsApp := e.Group("/contact-us")
	contactUsApp.GET("", h.FetchAllContactUsHome)
	adminApp := contactUsApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", h.CreateContactUs)
	adminApp.GET("", h.FetchAllContactUs)
	adminApp.GET("/:id", h.FetchByIDContactUs)
	adminApp.PUT("/:id", h.EditByIDContactUs)
	adminApp.DELETE("/:id", h.DeleteByIDContactUs)
	return h
}
