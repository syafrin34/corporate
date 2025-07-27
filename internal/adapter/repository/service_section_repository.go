package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ServiceSectionInterface interface {
	CreateServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error
	FetchAllServiceSection(ctx context.Context) ([]entity.ServiceSectionEntity, error)
	FetchByIDServiceSection(ctx context.Context, id int64) (*entity.ServiceSectionEntity, error)
	EditByIDServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error
	DeleteByIDServiceSection(ctx context.Context, id int64) error
}

type serviceSection struct {
	DB *gorm.DB
}

// CreateServiceSection implements ServiceSectionInterface.
func (s *serviceSection) CreateServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error {
	modelServiceSection := model.ServiceSection{
		Name:     req.Name,
		Tagline:  req.Tagline,
		PathIcon: req.PathIcon,
	}
	if err = s.DB.Create(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateServiceSection - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDServiceSection implements ServiceSectionInterface.
func (s *serviceSection) DeleteByIDServiceSection(ctx context.Context, id int64) error {
	modelServiceSection := model.ServiceSection{}
	if err = s.DB.Where("id = ?", id).First(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDServiceSection - 1: %v", err)
		return err
	}
	if err = s.DB.Delete(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDServiceSection -2: %v", err)
		return err
	}
	return nil

}

// EditByIDServiceSection implements ServiceSectionInterface.
func (s *serviceSection) EditByIDServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error {
	modelServiceSection := model.ServiceSection{}
	if err = s.DB.Where("id = ?", req.ID).First(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDServiceSection - 1: %v", err)
		return err
	}

	modelServiceSection.Name = req.Name
	modelServiceSection.PathIcon = req.PathIcon
	modelServiceSection.Tagline = req.Tagline

	if err = s.DB.Save(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDServiceSection - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllServiceSection implements ServiceSectionInterface.
func (s *serviceSection) FetchAllServiceSection(ctx context.Context) ([]entity.ServiceSectionEntity, error) {
	modelServiceSection := []model.ServiceSection{}
	if err = s.DB.Select("id", "path_icon", "tagline", "name").Find(&modelServiceSection).Order("created_at DESC").Error; err != nil {
		log.Errorf("[REPOSITORY] FetchAllServiceSection - 1: %v", err)
		return nil, err
	}

	var heroSectionsEntities []entity.ServiceSectionEntity
	for _, v := range modelServiceSection {
		heroSectionsEntities = append(heroSectionsEntities, entity.ServiceSectionEntity{
			ID:       v.ID,
			Name:     v.Name,
			PathIcon: v.PathIcon,
			Tagline:  v.Tagline,
		})
	}
	return heroSectionsEntities, nil
}

// FetchByIDServiceSection implements ServiceSectionInterface.
func (s *serviceSection) FetchByIDServiceSection(ctx context.Context, id int64) (*entity.ServiceSectionEntity, error) {
	modelServiceSection := model.ServiceSection{}
	if err = s.DB.Select("id", "path_icon", "tagline", "name").Where("id = ?", id).First(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDServiceSection - 1: %v", err)
		return nil, err
	}

	return &entity.ServiceSectionEntity{
		ID:       modelServiceSection.ID,
		Name:     modelServiceSection.Name,
		PathIcon: modelServiceSection.PathIcon,
		Tagline:  modelServiceSection.Tagline,
	}, nil
}

func NewServiceSectionRepository(DB *gorm.DB) ServiceSectionInterface {
	return &serviceSection{
		DB: DB,
	}
}
