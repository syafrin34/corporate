package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ClientSectionInterface interface {
	CreateClientSection(ctx context.Context, req entity.ClientSectionEntity) error
	FetchAllClientSection(ctx context.Context) ([]entity.ClientSectionEntity, error)
	FetchByIDClientSection(ctx context.Context, id int64) (*entity.ClientSectionEntity, error)
	EditByIDClientSection(ctx context.Context, req entity.ClientSectionEntity) error
	DeleteByIDClientSection(ctx context.Context, id int64) error
}

type clientSection struct {
	DB *gorm.DB
}

// CreateClientSection implements ClientSectionInterface.
func (h *clientSection) CreateClientSection(ctx context.Context, req entity.ClientSectionEntity) error {
	modelClientSection := model.ClientSection{
		Name:     req.Name,
		PathIcon: req.PathIcon,
	}
	if err = h.DB.Create(&modelClientSection).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateClientSEction - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDClientSection implements ClientSectionInterface.
func (h *clientSection) DeleteByIDClientSection(ctx context.Context, id int64) error {
	modelClientSection := model.ClientSection{}
	err = h.DB.Where("id = ?", id).First(&modelClientSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDClientSection - 1: %v", err)
		return err
	}
	err = h.DB.Delete(&modelClientSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDClientSection -2: %v", err)
		return err
	}
	return nil

}

// EditByIDClientSection implements ClientSectionInterface.
func (h *clientSection) EditByIDClientSection(ctx context.Context, req entity.ClientSectionEntity) error {
	modelClientSection := model.ClientSection{}
	err = h.DB.Where("id = ?", req.ID).First(&modelClientSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDClientSection - 1: %v", err)
		return err
	}
	modelClientSection.Name = req.Name
	modelClientSection.PathIcon = req.PathIcon

	err = h.DB.Save(&modelClientSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDClientSection - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllClientSection implements ClientSectionInterface.
func (h *clientSection) FetchAllClientSection(ctx context.Context) ([]entity.ClientSectionEntity, error) {
	modelClientSection := []model.ClientSection{}
	err = h.DB.Select("id", "name", "path_icon").Find(&modelClientSection).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllClientSection - 1: %v", err)
		return nil, err
	}

	var heroSectionsEntities []entity.ClientSectionEntity
	for _, v := range modelClientSection {
		heroSectionsEntities = append(heroSectionsEntities, entity.ClientSectionEntity{
			ID:       v.ID,
			Name:     v.Name,
			PathIcon: v.PathIcon,
		})
	}
	return heroSectionsEntities, nil
}

// FetchByIDClientSection implements ClientSectionInterface.
func (h *clientSection) FetchByIDClientSection(ctx context.Context, id int64) (*entity.ClientSectionEntity, error) {
	modelClientSection := model.ClientSection{}
	err = h.DB.Select("id", "name", "path_icon").Where("id = ?", id).First(&modelClientSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDClientSection - 1: %v", err)
		return nil, err
	}

	return &entity.ClientSectionEntity{
		ID:       modelClientSection.ID,
		Name:     modelClientSection.Name,
		PathIcon: modelClientSection.PathIcon,
	}, nil
}

func NewClientSectionRepository(DB *gorm.DB) ClientSectionInterface {
	return &clientSection{
		DB: DB,
	}
}
