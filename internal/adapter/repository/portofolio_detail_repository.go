package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PortofolioDetailInterface interface {
	CreatePortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error
	FetchAllPortofolioDetail(ctx context.Context) ([]entity.PortofolioDetailEntity, error)
	FetchByIDPortofolioDetail(ctx context.Context, id int64) (*entity.PortofolioDetailEntity, error)
	EditByIDPortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error
	DeleteByIDPortofolioDetail(ctx context.Context, id int64) error
	FetchdetailPortofolioByPortoID(ctx context.Context, id int64) (*entity.PortofolioDetailEntity, error)
}

type portofolioDetail struct {
	DB *gorm.DB
}

// FetchdetailPortofolioByPortoID implements PortofolioDetailInterface.
func (p *portofolioDetail) FetchdetailPortofolioByPortoID(ctx context.Context, portoID int64) (*entity.PortofolioDetailEntity, error) {
	rows, err := p.DB.Table("portofolio_details as pd").
		Select("pd.id", "pd.title", "pd.category", "pd.client_name", "pd.project_date", "pd.description", "pd.project_url", "ps.id", "ps.name", "ps.thumbnail").
		Joins("inner join portofolio_sections as ps on ps.id = pd.portofolio_section_id").
		Where("ps_id = ? AND pd.deleted_at IS NULL", portoID).Order("pd.created_at DESC").Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchDetailPortofolioByPortoID - 1: %v", err)
		return nil, err
	}

	var portofolioDetailEntity entity.PortofolioDetailEntity
	for rows.Next() {
		err = rows.Scan(
			&portofolioDetailEntity.ID,
			&portofolioDetailEntity.Title,
			&portofolioDetailEntity.Category,
			&portofolioDetailEntity.ClientName,
			&portofolioDetailEntity.ProjectDate,
			&portofolioDetailEntity.Description,
			&portofolioDetailEntity.ProjectUrl,
			&portofolioDetailEntity.PortofolioSection.ID,
			&portofolioDetailEntity.PortofolioSection.Name,
			&portofolioDetailEntity.PortofolioSection.Thumbnail,
		)

		if err != nil {
			log.Errorf("[REPOSITORY] FetchDetailPortofolioByPortoID - 2: %v", err)
			return nil, err
		}
	}
	return &portofolioDetailEntity, nil
}

// CreatePortofolioDetail implements PortofolioDetailInterface.
func (p *portofolioDetail) CreatePortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error {
	modelPortofolioDetail := model.PortofolioDetail{
		PortoFolioSectionID: req.PortofolioSection.ID,
		Category:            req.Category,
		ClientName:          req.ClientName,
		ProjectDate:         req.ProjectDate,
		ProjectUrl:          req.ProjectUrl,
		Title:               req.Title,
		Description:         req.Description,
	}
	if err = p.DB.Create(&modelPortofolioDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] CreatePortofolioDetail - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDPortofolioDetail implements PortofolioDetailInterface.
func (p *portofolioDetail) DeleteByIDPortofolioDetail(ctx context.Context, id int64) error {
	modelPortofolioDetail := model.PortofolioDetail{}
	if err = p.DB.Where("id = ?", id).First(&modelPortofolioDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioDetail - 1: %v", err)
		return err
	}
	if err = p.DB.Delete(&modelPortofolioDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioDetail -2: %v", err)
		return err
	}
	return nil

}

// EditByIDPortofolioDetail implements PortofolioDetailInterface.
func (p *portofolioDetail) EditByIDPortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error {
	modelPortofolioDetail := model.PortofolioDetail{}
	if err = p.DB.Where("id = ?", req.ID).First(&modelPortofolioDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioDetail - 1: %v", err)
		return err
	}

	modelPortofolioDetail.ID = req.PortofolioSection.ID
	modelPortofolioDetail.Category = req.Category
	modelPortofolioDetail.ClientName = req.ClientName
	modelPortofolioDetail.ProjectDate = req.ProjectDate
	modelPortofolioDetail.ProjectUrl = req.ProjectUrl
	modelPortofolioDetail.PortoFolioSectionID = req.PortofolioSection.ID

	if err = p.DB.Save(&modelPortofolioDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioDetail - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllPortofolioDetail implements PortofolioDetailInterface.
func (p *portofolioDetail) FetchAllPortofolioDetail(ctx context.Context) ([]entity.PortofolioDetailEntity, error) {

	rows, err := p.DB.Table("portofolio_details as pd").
		Select("pd.id", "pd.title", "pd.category", "pd.client_name", "pd.project_date", "ps.name").
		Joins("inner join portofolio_sections as ps on ps.id = pd.portofolio_section_id").
		Where("pd.deleted_at IS NULL").
		Order("created_at DESC").
		Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioDetail - 1: %v", err)
		return nil, err
	}
	var portofolioDetailEntities []entity.PortofolioDetailEntity
	for rows.Next() {
		portofolioDetail := entity.PortofolioDetailEntity{}
		err = rows.Scan(&portofolioDetail.ID,
			&portofolioDetail.Title,
			&portofolioDetail.Category,
			&portofolioDetail.ClientName,
			&portofolioDetail.ProjectDate,
			&portofolioDetail.PortofolioSection.Name)
		if err != nil {
			log.Errorf("[REPOSITORY] EditIDPortofolioDetail - 2: %v", err)
			return nil, err
		}
		portofolioDetailEntities = append(portofolioDetailEntities, portofolioDetail)
	}

	return portofolioDetailEntities, nil
}

// FetchByIDPortofolioDetail implements PortofolioDetailInterface.
func (p *portofolioDetail) FetchByIDPortofolioDetail(ctx context.Context, id int64) (*entity.PortofolioDetailEntity, error) {

	rows, err := p.DB.Table("portofolio_details as pd").
		Select("pd.id", "pd.title", "pd.category", "pd.client_name", "pd.project_date", "pd.description", "pd.project_url", "ps.id", "ps.name", "ps.thumbnail").
		Joins("inner join portofolio_sections as ps on ps.id = pd.portofolio_section_id").
		Where("pd_id = ? AND pd.deleted_at IS NULL", id).Order("pd.created_at DESC").Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioDetail - 1: %v", err)
		return nil, err
	}

	var portofolioDetailEntity entity.PortofolioDetailEntity
	for rows.Next() {
		err = rows.Scan(
			&portofolioDetailEntity.ID,
			&portofolioDetailEntity.Title,
			&portofolioDetailEntity.Category,
			&portofolioDetailEntity.ClientName,
			&portofolioDetailEntity.ProjectDate,
			&portofolioDetailEntity.Description,
			&portofolioDetailEntity.ProjectUrl,
			&portofolioDetailEntity.PortofolioSection.ID,
			&portofolioDetailEntity.PortofolioSection.Name,
			&portofolioDetailEntity.PortofolioSection.Thumbnail,
		)

		if err != nil {
			log.Errorf("[REPOSITORY] EditIDPortofolioDetail - 2: %v", err)
			return nil, err
		}
	}
	return &portofolioDetailEntity, nil

}

func NewPortofolioDetailRepository(DB *gorm.DB) PortofolioDetailInterface {
	return &portofolioDetail{
		DB: DB,
	}
}
