package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type AboutCompanyInterface interface {
	CreateAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error
	FetchAllAboutCompany(ctx context.Context) ([]entity.AboutCompanyEntity, error)
	FetchByIDAboutCompany(ctx context.Context, id int64) (*entity.AboutCompanyEntity, error)
	EditByIDAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error
	DeleteByIDAboutCompany(ctx context.Context, id int64) error
	FetchAllCompanyAndKeynote(ctx context.Context) ([]entity.AboutCompanyEntity, error)
}

type aboutCompany struct {
	DB *gorm.DB
}

// CreateAboutCompany implements AboutCompanyInterface.
func (a *aboutCompany) CreateAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error {
	modelAboutCompany := model.AboutCompany{
		Description: req.Description,
	}
	if err = a.DB.Create(&modelAboutCompany).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateAboutCompany - 1: %v", err)
		return err
	}
	return nil
}

func (a *aboutCompany) FetchAllCompanyAndKeynote(ctx context.Context) ([]entity.AboutCompanyEntity, error) {
	modelAboutCompany := []model.AboutCompany{}
	err := a.DB.Select("id", "description").Find(&modelAboutCompany).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllCompanyAndKeynote - 1: %v", err)
		return nil, err
	}

	var aboutCompanyRepoEntities []entity.AboutCompanyEntity
	for _, v := range modelAboutCompany {
		var aboutCompanyKeynoteModel []model.AboutCompanyKeynote
		err := a.DB.Select("id", "description").Where("about_company_id = ?", v.ID).Find(&aboutCompanyKeynoteModel).Error
		if err != nil {
			log.Errorf("[REPOSITORY] FetchAllCompanyAndKeynote - 2: %v", err)
			return nil, err
		}

		var aboutCompanyKeynoteEntity []entity.AboutCompanyKeynoteEntity
		for _, val := range aboutCompanyKeynoteModel {
			aboutCompanyKeynoteEntity = append(aboutCompanyKeynoteEntity, entity.AboutCompanyKeynoteEntity{
				ID:             val.ID,
				AboutCompanyID: v.ID,
				Keynote:        val.Keynote,
				PathImage:      *val.PathImage,
			})
		}
		aboutCompanyRepoEntities = append(aboutCompanyRepoEntities, entity.AboutCompanyEntity{
			ID:          v.ID,
			Description: v.Description,
			Keynote:     aboutCompanyKeynoteEntity,
		})
	}
	return aboutCompanyRepoEntities, nil
}

// DeleteByIDAboutCompany implements AboutCompanyInterface.
func (h *aboutCompany) DeleteByIDAboutCompany(ctx context.Context, id int64) error {
	modelAboutCompany := model.AboutCompany{}
	err = h.DB.Where("id = ?", id).First(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAboutCompany - 1: %v", err)
		return err
	}
	err = h.DB.Delete(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAboutCompany -2: %v", err)
		return err
	}
	return nil

}

// EditByIDAboutCompany implements AboutCompanyInterface.
func (h *aboutCompany) EditByIDAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error {
	modelAboutCompany := model.AboutCompany{}
	err = h.DB.Where("id = ?", req.ID).First(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDAboutCompany - 1: %v", err)
		return err
	}
	modelAboutCompany.Description = req.Description

	err = h.DB.Save(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDAboutCompany - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllAboutCompany implements AboutCompanyInterface.
func (h *aboutCompany) FetchAllAboutCompany(ctx context.Context) ([]entity.AboutCompanyEntity, error) {
	modelAboutCompany := []model.AboutCompany{}
	err = h.DB.Select("id", "description").Find(&modelAboutCompany).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllAboutCompany - 1: %v", err)
		return nil, err
	}

	var heroSectionsEntities []entity.AboutCompanyEntity
	for _, v := range modelAboutCompany {
		heroSectionsEntities = append(heroSectionsEntities, entity.AboutCompanyEntity{
			ID:          v.ID,
			Description: v.Description,
		})
	}
	return heroSectionsEntities, nil
}

// FetchByIDAboutCompany implements AboutCompanyInterface.
func (h *aboutCompany) FetchByIDAboutCompany(ctx context.Context, id int64) (*entity.AboutCompanyEntity, error) {
	modelAboutCompany := model.AboutCompany{}
	err = h.DB.Select("id", "description").Where("id = ?", id).First(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDAboutCompany - 1: %v", err)
		return nil, err
	}

	return &entity.AboutCompanyEntity{
		ID:          modelAboutCompany.ID,
		Description: modelAboutCompany.Description,
	}, nil
}

func NewAboutCompanyRepository(DB *gorm.DB) AboutCompanyInterface {
	return &aboutCompany{
		DB: DB,
	}
}
