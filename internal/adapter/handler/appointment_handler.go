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

type AppointmentHandlerInterface interface {
	//CreateAppointment(c echo.Context) error
	FetchAllAppointment(c echo.Context) error
	FetchByIDAppointment(c echo.Context) error
	//EditByIDAppointment(c echo.Context) error
	DeleteByIDAppointment(c echo.Context) error
	CreateAppointment(c echo.Context) error
}
type appointmentHandler struct {
	AppointmentService service.AppointmentServiceInterface
}

// DeleteByIDAppointment implements CLientSectionHandlerInterface.
func (a *appointmentHandler) DeleteByIDAppointment(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteAppointment - 1: Unautohrized")
		respError.Meta.Message = "UnAuthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAppointment := c.Param("id")
	id, err := conv.StringToInt64(idAppointment)
	if err != nil {
		log.Errorf("[HANDLER] DeleteAppointment - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false

		return c.JSON(http.StatusBadRequest, respError)
	}

	err = a.AppointmentService.DeleteByIDAppointment(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteAppointment - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success delete appointment"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func (a *appointmentHandler) CreateAppointment(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		req       = request.AppointmentRequest{}
		ctx       = c.Request().Context()
	)

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateAppointment - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateAppointment - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	stringprojectDate, err := time.Parse("2006-01-02", req.MeetAt)
	if err != nil {
		log.Errorf("[HANDLER] CreateAppointment - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.AppointmentEntity{
		ServiceID:   req.ServiceID,
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Brief:       req.Brief,
		Budget:      req.Budget,
		MeetAt:      stringprojectDate,
	}

	err = a.AppointmentService.CreateAppointment(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateAppointment - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create appointment"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)

}

// FetchAllAppointment implements CLientSectionHandlerInterface.
func (a *appointmentHandler) FetchAllAppointment(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = []response.AppointmentResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchALLAppointment - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := a.AppointmentService.FetchAllAppointment(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchALLAppointment - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClient = append(respClient, response.AppointmentResponse{
			ID:          val.ID,
			Name:        val.Name,
			PhoneNumber: val.PhoneNumber,
			Email:       val.Email,
			Brief:       val.Brief,
			Budget:      val.Budget,
			MeetAt:      val.MeetAt.Format("02 Jan 2026 15:04:05"),
			ServiceName: val.ServiceName,
			ServiceID:   val.ServiceID,
		})
	}
	resp.Meta.Message = "Success fetch all appointment"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDAppointment implements CLientSectionHandlerInterface.
func (a *appointmentHandler) FetchByIDAppointment(c echo.Context) error {
	var (
		resp            = response.DefaultSuccessResponse{}
		respError       = response.ErrorResponseDefault{}
		ctx             = c.Request().Context()
		respAppointment = response.AppointmentResponse{}
	)

	user := conv.GetUserByIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByID Appointment - 1: Unautohrized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAppointment := c.Param("id")
	id, err := conv.StringToInt64(idAppointment)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID Appointment - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := a.AppointmentService.FetchByIDAppointment(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByID Service Section - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respAppointment.ID = result.ID
	respAppointment.Name = result.Name
	respAppointment.PhoneNumber = result.PhoneNumber
	respAppointment.Email = result.Email
	respAppointment.Brief = result.Brief
	respAppointment.Budget = result.Budget
	respAppointment.MeetAt = result.MeetAt.Format("02 Jan 2026 15:04:05")
	respAppointment.ServiceName = result.ServiceName
	respAppointment.ServiceID = result.ServiceID
	resp.Meta.Message = "Success fetch appointment by ID"
	resp.Meta.Status = true
	resp.Data = respAppointment
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewAppointmentHandler(e *echo.Echo, AppointmentService service.AppointmentServiceInterface, cfg *config.Config) AppointmentHandlerInterface {
	h := &appointmentHandler{
		AppointmentService: AppointmentService,
	}

	mid := middleware.NewMiddleware(cfg)
	appointmentApp := e.Group("/appointments")
	appointmentApp.POST("", h.CreateAppointment)
	adminApp := appointmentApp.Group("/admin", mid.CheckToken())
	//adminApp.POST("", h.CreateAppointment)
	adminApp.GET("", h.FetchAllAppointment)
	adminApp.GET("/:id", h.FetchByIDAppointment)
	//adminApp.PUT("/:id", h.EditByIDAppointment)
	adminApp.DELETE("/:id", h.DeleteByIDAppointment)
	return h
}
