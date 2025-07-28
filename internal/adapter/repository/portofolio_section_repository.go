package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PortofolioSectionInterface interface {
	CreatePortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error
	FetchAllPortofolioSection(ctx context.Context) ([]entity.PortofolioSectionEntity, error)
	FetchByIDPortofolioSection(ctx context.Context, id int64) (*entity.PortofolioSectionEntity, error)
	EditByIDPortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error
	DeleteByIDPortofolioSection(ctx context.Context, id int64) error
}

type portofolioSection struct {
	DB *gorm.DB
}

// CreatePortofolioSection implements PortofolioSectionInterface.
func (p *portofolioSection) CreatePortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error {
	modelPortofolioSection := model.PortofolioSection{
		Name:      req.Name,
		Tagline:   req.Tagline,
		Thumbnail: &req.Thumbnail,
	}
	if err = p.DB.Create(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] CreatePortofolioSection - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDPortofolioSection implements PortofolioSectionInterface.
func (p *portofolioSection) DeleteByIDPortofolioSection(ctx context.Context, id int64) error {
	modelPortofolioSection := model.PortofolioSection{}
	if err = p.DB.Where("id = ?", id).First(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioSection - 1: %v", err)
		return err
	}
	if err = p.DB.Delete(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioSection -2: %v", err)
		return err
	}
	return nil

}

// EditByIDPortofolioSection implements PortofolioSectionInterface.
func (p *portofolioSection) EditByIDPortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error {
	modelPortofolioSection := model.PortofolioSection{}
	if err = p.DB.Where("id = ?", req.ID).First(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioSection - 1: %v", err)
		return err
	}

	modelPortofolioSection.Name = req.Name
	modelPortofolioSection.Tagline = req.Tagline
	modelPortofolioSection.Thumbnail = &req.Thumbnail

	if err = p.DB.Save(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioSection - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllPortofolioSection implements PortofolioSectionInterface.
func (p *portofolioSection) FetchAllPortofolioSection(ctx context.Context) ([]entity.PortofolioSectionEntity, error) {
	modelPortofolioSection := []model.PortofolioSection{}
	if err = p.DB.Select("id", "name", "tagline", "thumbnail").Find(&modelPortofolioSection).Order("created_at DESC").Error; err != nil {
		log.Errorf("[REPOSITORY] FetchAllPortofolioSection - 1: %v", err)
		return nil, err
	}

	var heroSectionsEntities []entity.PortofolioSectionEntity
	for _, v := range modelPortofolioSection {
		heroSectionsEntities = append(heroSectionsEntities, entity.PortofolioSectionEntity{
			ID:        v.ID,
			Name:      v.Name,
			Thumbnail: *v.Thumbnail,
			Tagline:   v.Tagline,
		})
	}
	return heroSectionsEntities, nil
}

// FetchByIDPortofolioSection implements PortofolioSectionInterface.
func (p *portofolioSection) FetchByIDPortofolioSection(ctx context.Context, id int64) (*entity.PortofolioSectionEntity, error) {
	modelPortofolioSection := model.PortofolioSection{}
	if err = p.DB.Select("id", "path_icon", "tagline", "name").Where("id = ?", id).First(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioSection - 1: %v", err)
		return nil, err
	}

	return &entity.PortofolioSectionEntity{
		ID:        modelPortofolioSection.ID,
		Name:      modelPortofolioSection.Name,
		Thumbnail: *modelPortofolioSection.Thumbnail,
		Tagline:   modelPortofolioSection.Tagline,
	}, nil
}

func NewPortofolioSectionRepository(DB *gorm.DB) PortofolioSectionInterface {
	return &portofolioSection{
		DB: DB,
	}
}
