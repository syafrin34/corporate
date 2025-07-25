package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type FaqSectionInterface interface {
	CreateFaqSection(ctx context.Context, req entity.FaqSectionEntity) error
	FetchAllFaqSection(ctx context.Context) ([]entity.FaqSectionEntity, error)
	FetchByIDFaqSection(ctx context.Context, id int64) (*entity.FaqSectionEntity, error)
	EditByIDFaqSection(ctx context.Context, req entity.FaqSectionEntity) error
	DeleteByIDFaqSection(ctx context.Context, id int64) error
}

type faqSection struct {
	DB *gorm.DB
}

// CreateFaqSection implements FaqSectionInterface.
func (f *faqSection) CreateFaqSection(ctx context.Context, req entity.FaqSectionEntity) error {
	modelFaqSection := model.FaqSection{
		Title:       req.Title,
		Description: req.Description,
	}
	if err = f.DB.Create(&modelFaqSection).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateFaqSection - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDFaqSection implements FaqSectionInterface.
func (f *faqSection) DeleteByIDFaqSection(ctx context.Context, id int64) error {
	modelFaqSection := model.FaqSection{}
	err = f.DB.Where("id = ?", id).First(&modelFaqSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDFaqSection - 1: %v", err)
		return err
	}
	err = f.DB.Delete(&modelFaqSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDFaqSection -2: %v", err)
		return err
	}
	return nil

}

// EditByIDFaqSection implements FaqSectionInterface.
func (f *faqSection) EditByIDFaqSection(ctx context.Context, req entity.FaqSectionEntity) error {
	modelFaqSection := model.FaqSection{}
	err = f.DB.Where("id = ?", req.ID).First(&modelFaqSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDFaqSection - 1: %v", err)
		return err
	}
	modelFaqSection.Title = req.Title
	modelFaqSection.Description = req.Description

	err = f.DB.Save(&modelFaqSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDFaqSection - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllFaqSection implements FaqSectionInterface.
func (f *faqSection) FetchAllFaqSection(ctx context.Context) ([]entity.FaqSectionEntity, error) {
	modelFaqSection := []model.FaqSection{}
	err = f.DB.Select("id", "title", "description").Find(&modelFaqSection).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllFaqSection - 1: %v", err)
		return nil, err
	}

	var faqSectionsEntities []entity.FaqSectionEntity
	for _, v := range modelFaqSection {
		faqSectionsEntities = append(faqSectionsEntities, entity.FaqSectionEntity{
			ID:          v.ID,
			Description: v.Description,
		})
	}
	return faqSectionsEntities, nil
}

// FetchByIDFaqSection implements FaqSectionInterface.
func (f *faqSection) FetchByIDFaqSection(ctx context.Context, id int64) (*entity.FaqSectionEntity, error) {
	modelFaqSection := model.FaqSection{}
	err = f.DB.Select("id", "title", "description").Where("id = ?", id).First(&modelFaqSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDFaqSection - 1: %v", err)
		return nil, err
	}

	return &entity.FaqSectionEntity{
		ID:          modelFaqSection.ID,
		Title:       modelFaqSection.Title,
		Description: modelFaqSection.Description,
	}, nil
}

func NewFaqSectionRepository(DB *gorm.DB) FaqSectionInterface {
	return &faqSection{
		DB: DB,
	}
}
