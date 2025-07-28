package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type AppointmentInterface interface {
	//CreateAppointment(ctx context.Context, req entity.AppointmentEntity) error
	FetchAllAppointment(ctx context.Context) ([]entity.AppointmentEntity, error)
	FetchByIDAppointment(ctx context.Context, id int64) (*entity.AppointmentEntity, error)
	//EditByIDAppointment(ctx context.Context, req entity.AppointmentEntity) error
	DeleteByIDAppointment(ctx context.Context, id int64) error
}

type appointment struct {
	DB *gorm.DB
}

// DeleteByIDAppointment implements AppointmentInterface.
func (a *appointment) DeleteByIDAppointment(ctx context.Context, id int64) error {
	modelAppointment := model.Appointment{}
	if err = a.DB.Where("id = ?", id).First(&modelAppointment).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAppointment - 1: %v", err)
		return err
	}
	if err = a.DB.Delete(&modelAppointment).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAppointment -2: %v", err)
		return err
	}
	return nil

}

// FetchAllAppointment implements AppointmentInterface.
func (a *appointment) FetchAllAppointment(ctx context.Context) ([]entity.AppointmentEntity, error) {

	rows, err := a.DB.Table("appointments as a").Select("a.id", "a.name", "a.budget", "ss.name").
		Joins("inner join service_sections as ss on ss.id = a.service_id").Rows()

	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllAppointment - 1: %v", err)
		return nil, err
	}

	var appointmentEntities []entity.AppointmentEntity
	for rows.Next() {
		var appointment entity.AppointmentEntity
		err = rows.Scan(&appointment.ID, &appointment.Name, &appointment.Email, &appointment.Budget, &appointment.ServiceName)
		if err != nil {
			log.Errorf("[REPOSITORY] FetchAllAppointment - 2: %v", err)
			return nil, err
		}
		appointmentEntities = append(appointmentEntities, appointment)

	}
	return appointmentEntities, nil
}

// FetchByIDAppointment implements AppointmentInterface.
func (a *appointment) FetchByIDAppointment(ctx context.Context, id int64) (*entity.AppointmentEntity, error) {

	rows, err := a.DB.Table("appointments as a").Select("a.id", "a.phone_number", "brief", "meet_at", "a.name", "a.email", "a.budget", "ss.id", "ss.name").
		Joins("inner join service_sections as ss on ss.id = a.service_id").
		Where("a.id = ?", id).Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDAppointment - 1: %v", err)
		return nil, err
	}

	var appointment = &entity.AppointmentEntity{}
	for rows.Next() {
		err = rows.Scan(&appointment.ID, &appointment.PhoneNumber, &appointment.Brief, &appointment.MeetAt, &appointment.Name, &appointment.Email, &appointment.Budget, &appointment.ServiceID, &appointment.ServiceName)
	}
	return appointment, nil
}

func NewAppointmentRepository(DB *gorm.DB) AppointmentInterface {
	return &appointment{
		DB: DB,
	}
}
