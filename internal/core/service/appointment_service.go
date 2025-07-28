package service

import (
	"context"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"
)

type AppointmentServiceInterface interface {
	//CreateAppointment(ctx context.Context, req entity.AppointmentEntity) error
	FetchAllAppointment(ctx context.Context) ([]entity.AppointmentEntity, error)
	FetchByIDAppointment(ctx context.Context, id int64) (*entity.AppointmentEntity, error)
	//EditByIDAppointment(ctx context.Context, req entity.AppointmentEntity) error
	DeleteByIDAppointment(ctx context.Context, id int64) error
}

type appointmentService struct {
	appointmentRepo repository.AppointmentInterface
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

func NewAppointmentService(appointmentRepo repository.AppointmentInterface) AppointmentServiceInterface {
	return &appointmentService{
		appointmentRepo: appointmentRepo,
	}
}
