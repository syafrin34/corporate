package service

import (
	"context"
	"corporate/internal/adapter/messaging"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"
	"fmt"

	"github.com/labstack/gommon/log"
)

type AppointmentServiceInterface interface {
	//CreateAppointment(ctx context.Context, req entity.AppointmentEntity) error
	FetchAllAppointment(ctx context.Context) ([]entity.AppointmentEntity, error)
	FetchByIDAppointment(ctx context.Context, id int64) (*entity.AppointmentEntity, error)
	//EditByIDAppointment(ctx context.Context, req entity.AppointmentEntity) error
	DeleteByIDAppointment(ctx context.Context, id int64) error
	CreateAppointment(ctx context.Context, req entity.AppointmentEntity) error
}

type appointmentService struct {
	appointmentRepo repository.AppointmentInterface
	sendEmail       messaging.EmailMessagingInterface
}

// CreateAppointment implements AppointmentServiceInterface.
func (s *appointmentService) CreateAppointment(ctx context.Context, req entity.AppointmentEntity) error {
	email, err := s.appointmentRepo.CreateAppointment(ctx, req)
	if err != nil {
		log.Errorf("[SERVICE] CreateAppointment -1: %v", err)
		return err
	}

	body := fmt.Sprintf("You have received a new appointment request fro %s", email)
	err = s.sendEmail.SendEMailAppointment(nil, email, "New Appointment", body)
	if err != nil {
		log.Errorf("[SERVICE] CreateAppointment -2: %v", err)
		return err
	}
	return nil
}

// DeleteByIDAppointment implements AppointmentServiceInterface.
func (s *appointmentService) DeleteByIDAppointment(ctx context.Context, id int64) error {
	return s.appointmentRepo.DeleteByIDAppointment(ctx, id)
}

// FetchAllAppointment implements AppointmentServiceInterface.
func (s *appointmentService) FetchAllAppointment(ctx context.Context) ([]entity.AppointmentEntity, error) {
	return s.appointmentRepo.FetchAllAppointment(ctx)
}

// FetchByIDAppointment implements AppointmentServiceInterface.
func (s *appointmentService) FetchByIDAppointment(ctx context.Context, id int64) (*entity.AppointmentEntity, error) {
	return s.appointmentRepo.FetchByIDAppointment(ctx, id)
}

func NewAppointmentService(appointmentRepo repository.AppointmentInterface, sendEmail messaging.EmailMessagingInterface) AppointmentServiceInterface {
	return &appointmentService{
		appointmentRepo: appointmentRepo,
		sendEmail:       sendEmail,
	}
}
