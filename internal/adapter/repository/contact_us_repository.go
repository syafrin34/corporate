package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ContactUsInterface interface {
	CreateContactUs(ctx context.Context, req entity.ContactUsEntity) error
	FetchAllContactUs(ctx context.Context) ([]entity.ContactUsEntity, error)
	FetchByIDContactUs(ctx context.Context, id int64) (*entity.ContactUsEntity, error)
	EditByIDContactUs(ctx context.Context, req entity.ContactUsEntity) error
	DeleteByIDContactUs(ctx context.Context, id int64) error
}

type contactUs struct {
	DB *gorm.DB
}

// CreateContactUs implements ContactUsInterface.
func (c *contactUs) CreateContactUs(ctx context.Context, req entity.ContactUsEntity) error {
	modelContactUs := model.ContactUs{
		CompanyName:  req.CompanyName,
		LocationName: req.LocationName,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
	}
	if err = c.DB.Create(&modelContactUs).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateContactUs - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDContactUs implements ContactUsInterface.
func (c *contactUs) DeleteByIDContactUs(ctx context.Context, id int64) error {
	modelContactUs := model.ContactUs{}
	err = c.DB.Where("id = ?", id).First(&modelContactUs).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDContactUs - 1: %v", err)
		return err
	}
	err = c.DB.Delete(&modelContactUs).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDContactUs -2: %v", err)
		return err
	}
	return nil

}

// EditByIDContactUs implements ContactUsInterface.
func (c *contactUs) EditByIDContactUs(ctx context.Context, req entity.ContactUsEntity) error {
	modelContactUs := model.ContactUs{}
	err = c.DB.Where("id = ?", req.ID).First(&modelContactUs).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDContactUs - 1: %v", err)
		return err
	}
	modelContactUs.CompanyName = req.CompanyName
	modelContactUs.LocationName = req.LocationName
	modelContactUs.Address = req.Address
	modelContactUs.PhoneNumber = req.PhoneNumber

	err = c.DB.Save(&modelContactUs).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDContactUs - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllContactUs implements ContactUsInterface.
func (h *contactUs) FetchAllContactUs(ctx context.Context) ([]entity.ContactUsEntity, error) {
	modelContactUs := []model.ContactUs{}
	err = h.DB.Select("id", "company_name", "location_name", "address", "phone_number").Find(&modelContactUs).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllContactUs - 1: %v", err)
		return nil, err
	}

	var heroSectionsEntities []entity.ContactUsEntity
	for _, v := range modelContactUs {
		heroSectionsEntities = append(heroSectionsEntities, entity.ContactUsEntity{
			ID:           v.ID,
			CompanyName:  v.CompanyName,
			LocationName: v.LocationName,
			Address:      v.Address,
			PhoneNumber:  v.PhoneNumber,
		})
	}
	return heroSectionsEntities, nil
}

// FetchByIDContactUs implements ContactUsInterface.
func (h *contactUs) FetchByIDContactUs(ctx context.Context, id int64) (*entity.ContactUsEntity, error) {
	modelContactUs := model.ContactUs{}
	err = h.DB.Select("id", "company_name", "location_name", "address", "phone_number").Where("id = ?", id).First(&modelContactUs).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDContactUs - 1: %v", err)
		return nil, err
	}

	return &entity.ContactUsEntity{
		ID:           modelContactUs.ID,
		CompanyName:  modelContactUs.CompanyName,
		LocationName: modelContactUs.LocationName,
		Address:      modelContactUs.Address,
		PhoneNumber:  modelContactUs.PhoneNumber,
	}, nil
}

func NewContactUsRepository(DB *gorm.DB) ContactUsInterface {
	return &contactUs{
		DB: DB,
	}
}
