package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PortofolioTestimonialInterface interface {
	CreatePortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error
	FetchAllPortofolioTestimonial(ctx context.Context) ([]entity.PortofolioTestimonialEntity, error)
	FetchByIDPortofolioTestimonial(ctx context.Context, id int64) (*entity.PortofolioTestimonialEntity, error)
	EditByIDPortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error
	DeleteByIDPortofolioTestimonial(ctx context.Context, id int64) error
}

type portofolioTestimonial struct {
	DB *gorm.DB
}

// CreatePortofolioTestimonial implements PortofolioTestimonialInterface.
func (p *portofolioTestimonial) CreatePortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error {
	modelPortofolioTestimonial := model.PortofolioTestimonial{
		PortoFolioSectionID: req.PortoFolioSection.ID,
		Thumbnail:           req.Thumbnail,
		Message:             req.Message,
		ClientName:          req.ClientName,
		Role:                req.Role,
	}
	if err = p.DB.Create(&modelPortofolioTestimonial).Error; err != nil {
		log.Errorf("[REPOSITORY] CreatePortofolioTestimonial - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDPortofolioTestimonial implements PortofolioTestimonialInterface.
func (p *portofolioTestimonial) DeleteByIDPortofolioTestimonial(ctx context.Context, id int64) error {
	modelPortofolioTestimonial := model.PortofolioTestimonial{}
	if err = p.DB.Where("id = ?", id).First(&modelPortofolioTestimonial).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioTestimonial - 1: %v", err)
		return err
	}
	if err = p.DB.Delete(&modelPortofolioTestimonial).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioTestimonial -2: %v", err)
		return err
	}
	return nil

}

// EditByIDPortofolioTestimonial implements PortofolioTestimonialInterface.
func (p *portofolioTestimonial) EditByIDPortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error {
	modelPortofolioTestimonial := model.PortofolioTestimonial{}
	if err = p.DB.Where("id = ?", req.ID).First(&modelPortofolioTestimonial).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioTestimonial - 1: %v", err)
		return err
	}

	modelPortofolioTestimonial.ID = req.PortoFolioSection.ID
	modelPortofolioTestimonial.Thumbnail = req.Thumbnail
	modelPortofolioTestimonial.Message = req.Message
	modelPortofolioTestimonial.ClientName = req.ClientName
	modelPortofolioTestimonial.Role = req.Role

	if err = p.DB.Save(&modelPortofolioTestimonial).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioTestimonial - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllPortofolioTestimonial implements PortofolioTestimonialInterface.
func (p *portofolioTestimonial) FetchAllPortofolioTestimonial(ctx context.Context) ([]entity.PortofolioTestimonialEntity, error) {

	rows, err := p.DB.Table("portofolio_testimonials as pt").
		Select("pt.id", "pt.thumbnail", "pt.message", "pt.client_name", "pt.role", "ps.name").
		Joins("inner join portofolio_sections as ps on ps.id = pt.portofolio_section_id").Order("created_at DESC").Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioTestimonial - 1: %v", err)
		return nil, err
	}
	var portofolioTestimonialEntities []entity.PortofolioTestimonialEntity
	for rows.Next() {
		portofolioTestimonial := entity.PortofolioTestimonialEntity{}
		err = rows.Scan(&portofolioTestimonial.ID,
			&portofolioTestimonial.Thumbnail,
			&portofolioTestimonial.Message,
			&portofolioTestimonial.ClientName,
			&portofolioTestimonial.Role,
			&portofolioTestimonial.PortoFolioSection.Name)
		if err != nil {
			log.Errorf("[REPOSITORY] EditIDPortofolioTestimonial - 2: %v", err)
			return nil, err
		}
		portofolioTestimonialEntities = append(portofolioTestimonialEntities, portofolioTestimonial)
	}

	return portofolioTestimonialEntities, nil
}

// FetchByIDPortofolioTestimonial implements PortofolioTestimonialInterface.
func (p *portofolioTestimonial) FetchByIDPortofolioTestimonial(ctx context.Context, id int64) (*entity.PortofolioTestimonialEntity, error) {

	rows, err := p.DB.Table("portofolio_testimonials as pt").
		Select("pt.id", "pt.thumbnail", "pt.message", "pd.client_name", "pd.role", "ps.id", "ps.name", "ps.id", "ps.name", "ps.thumbnail").
		Joins("inner join portofolio_sections as ps on ps.id = pt.portofolio_section_id").
		Where("pt_id = ? ", id).Order("created_at DESC").Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioTestimonial - 1: %v", err)
		return nil, err
	}

	var portofolioTestimonialEntity entity.PortofolioTestimonialEntity
	for rows.Next() {
		err = rows.Scan(
			&portofolioTestimonialEntity.ID,
			&portofolioTestimonialEntity.Thumbnail,
			&portofolioTestimonialEntity.Message,
			&portofolioTestimonialEntity.ClientName,
			&portofolioTestimonialEntity.Role,
			&portofolioTestimonialEntity.PortoFolioSection.ID,
			&portofolioTestimonialEntity.PortoFolioSection.Name,
			&portofolioTestimonialEntity.PortoFolioSection.Thumbnail,
		)

		if err != nil {
			log.Errorf("[REPOSITORY] EditIDPortofolioTestimonial - 2: %v", err)
			return nil, err
		}
	}
	return &portofolioTestimonialEntity, nil

}

func NewPortofolioTestimonialRepository(DB *gorm.DB) PortofolioTestimonialInterface {
	return &portofolioTestimonial{
		DB: DB,
	}
}
