package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type AboutCompanyKeynoteInterface interface {
	CreateAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error
	FetchAllAboutCompanyKeynote(ctx context.Context) ([]entity.AboutCompanyKeynoteEntity, error)
	FetchByIDAboutCompanyKeynote(ctx context.Context, id int64) (*entity.AboutCompanyKeynoteEntity, error)
	EditByIDAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error
	DeleteByIDAboutCompanyKeynote(ctx context.Context, id int64) error
}

type aboutCompanyKeynote struct {
	DB *gorm.DB
}

// CreateAboutCompanyKeynote implements AboutCompanyKeynoteInterface.
func (a *aboutCompanyKeynote) CreateAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error {
	modelAboutCompanyKeynote := model.AboutCompanyKeynote{
		AboutCompanyID: req.AboutCompanyID,
		Keynote:        req.Keynote,
		PathImage:      &req.PathImage,
	}
	if err = a.DB.Create(&modelAboutCompanyKeynote).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateAboutCompanyKeynote - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDAboutCompanyKeynote implements AboutCompanyKeynoteInterface.
func (a *aboutCompanyKeynote) DeleteByIDAboutCompanyKeynote(ctx context.Context, id int64) error {
	modelAboutCompanyKeynote := model.AboutCompanyKeynote{}
	err = a.DB.Where("id = ?", id).First(&modelAboutCompanyKeynote).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAboutCompanyKeynote - 1: %v", err)
		return err
	}
	err = a.DB.Delete(&modelAboutCompanyKeynote).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAboutCompanyKeynote -2: %v", err)
		return err
	}
	return nil

}

// EditByIDAboutCompanyKeynote implements AboutCompanyKeynoteInterface.
func (a *aboutCompanyKeynote) EditByIDAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error {
	modelAboutCompanyKeynote := model.AboutCompanyKeynote{}
	err = a.DB.Where("id = ?", req.ID).First(&modelAboutCompanyKeynote).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDAboutCompanyKeynote - 1: %v", err)
		return err
	}
	modelAboutCompanyKeynote.AboutCompanyID = req.AboutCompanyID
	modelAboutCompanyKeynote.Keynote = req.Keynote
	modelAboutCompanyKeynote.PathImage = &req.PathImage

	err = a.DB.Save(&modelAboutCompanyKeynote).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDAboutCompanyKeynote - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllAboutCompanyKeynote implements AboutCompanyKeynoteInterface.
func (a *aboutCompanyKeynote) FetchAllAboutCompanyKeynote(ctx context.Context) ([]entity.AboutCompanyKeynoteEntity, error) {

	rows, err := a.DB.Table("about_company_keynotes as ack").Select("ack.id", "ack.keynote",
		"ack.about_company_id", "ack.path_image", "ac.description").Joins("inner join about_companies as ac on ac.id = ack.about_company_id").Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllAboutCompanyKeynote - 1: %v", err)
		return nil, err
	}
	//defer rows.Close()
	var aboutCompanyKeynoteEntities []entity.AboutCompanyKeynoteEntity
	for rows.Next() {
		aboutCompanyKeynote := entity.AboutCompanyKeynoteEntity{}
		err = rows.Scan(&aboutCompanyKeynote.ID, &aboutCompanyKeynote.Keynote, &aboutCompanyKeynote.AboutCompanyID, &aboutCompanyKeynote.PathImage, &aboutCompanyKeynote.AboutCompanyDescription)
		if err != nil {
			log.Errorf("[REPOSITORY] FetchAllAboutCompanyKeynote - 2: %v", err)
			return nil, err
		}
		aboutCompanyKeynoteEntities = append(aboutCompanyKeynoteEntities, aboutCompanyKeynote)
	}

	return aboutCompanyKeynoteEntities, nil
}

// FetchByIDAboutCompanyKeynote implements AboutCompanyKeynoteInterface.
func (a *aboutCompanyKeynote) FetchByIDAboutCompanyKeynote(ctx context.Context, id int64) (*entity.AboutCompanyKeynoteEntity, error) {

	rows, err := a.DB.Table("about_company_keynotes as ack").Select("ack.id", "ack.keynote",
		"ack.about_company_id", "ack.path_image", "ac.description").Joins("inner join about_companies as ac on ac.id = ack.about_company_id").Where("ack.id = ?", id).Rows()

	if err != nil {
		log.Errorf("[REPOSITORY] EditIDAboutCompanyKeynote - 1: %v", err)
		return nil, err
	}

	respEntity := entity.AboutCompanyKeynoteEntity{}
	for rows.Next() {
		err = rows.Scan(&respEntity.ID, &respEntity.Keynote, &respEntity.AboutCompanyID, &respEntity.PathImage, &respEntity.AboutCompanyDescription)
		if err != nil {
			log.Errorf("[REPOSITORY] FetchAllAboutCompanyKeynote - 2: %v", err)
			return nil, err
		}
	}
	return &respEntity, nil
}

func NewAboutCompanyKeynoteRepository(DB *gorm.DB) AboutCompanyKeynoteInterface {
	return &aboutCompanyKeynote{
		DB: DB,
	}
}
