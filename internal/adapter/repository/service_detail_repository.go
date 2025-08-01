package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ServiceDetailInterface interface {
	CreateServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error
	FetchAllServiceDetail(ctx context.Context) ([]entity.ServiceDetailEntity, error)
	FetchByIDServiceDetail(ctx context.Context, id int64) (*entity.ServiceDetailEntity, error)
	EditByIDServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error
	DeleteByIDServiceDetail(ctx context.Context, id int64) error
	GetByServiceIDDetail(ctx context.Context, serviceId int64) (*entity.ServiceDetailEntity, error)
}

type serviceDetail struct {
	DB *gorm.DB
}

// CreateServiceDetail implements ServiceDetailInterface.
func (s *serviceDetail) CreateServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error {
	modelServiceDetail := model.ServiceDetail{
		ServiceID:   req.ServiceID,
		PathImage:   req.PathImage,
		Description: req.Description,
		PathPdf:     req.PathPdf,
		Title:       req.Title,
		PathDocx:    req.PathDocx,
	}
	if err = s.DB.Create(&modelServiceDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateServiceDetail - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDServiceDetail implements ServiceDetailInterface.
func (s *serviceDetail) DeleteByIDServiceDetail(ctx context.Context, id int64) error {
	modelServiceDetail := model.ServiceDetail{}
	if err = s.DB.Where("id = ?", id).First(&modelServiceDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDServiceDetail - 1: %v", err)
		return err
	}
	if err = s.DB.Delete(&modelServiceDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDServiceDetail -2: %v", err)
		return err
	}
	return nil

}

// EditByIDServiceDetail implements ServiceDetailInterface.
func (s *serviceDetail) EditByIDServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error {
	modelServiceDetail := model.ServiceDetail{}
	if err = s.DB.Where("id = ?", req.ID).First(&modelServiceDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDServiceDetail - 1: %v", err)
		return err
	}

	modelServiceDetail.Description = req.Description
	modelServiceDetail.PathImage = req.PathImage
	modelServiceDetail.ServiceID = req.ServiceID
	modelServiceDetail.PathPdf = req.PathPdf
	modelServiceDetail.PathDocx = req.PathDocx
	modelServiceDetail.Title = req.Title

	if err = s.DB.Save(&modelServiceDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDServiceDetail - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllServiceDetail implements ServiceDetailInterface.
func (s *serviceDetail) FetchAllServiceDetail(ctx context.Context) ([]entity.ServiceDetailEntity, error) {
	modelServiceDetail := []model.ServiceDetail{}
	if err = s.DB.Select("id", "path_image", "description", "title", "path_pdf", "path_docx", "service_id").Find(&modelServiceDetail).Order("created_at DESC").Error; err != nil {
		log.Errorf("[REPOSITORY] FetchAllServiceDetail - 1: %v", err)
		return nil, err
	}

	var serviceDetailEntities []entity.ServiceDetailEntity
	for _, v := range modelServiceDetail {
		serviceDetailEntities = append(serviceDetailEntities, entity.ServiceDetailEntity{
			ID:          v.ID,
			ServiceID:   v.ServiceID,
			PathImage:   v.PathImage,
			Description: v.Description,
			PathPdf:     v.PathPdf,
			PathDocx:    v.PathDocx,
			Title:       v.Title,
		})
	}
	return serviceDetailEntities, nil
}

// FetchByIDServiceDetail implements ServiceDetailInterface.
func (s *serviceDetail) FetchByIDServiceDetail(ctx context.Context, id int64) (*entity.ServiceDetailEntity, error) {
	modelServiceDetail := model.ServiceDetail{}
	if err = s.DB.Select("id", "path_image", "description", "title", "path_pdf", "path_docx", "service_id").Where("id = ?", id).First(&modelServiceDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] EditIDServiceDetail - 1: %v", err)
		return nil, err
	}

	return &entity.ServiceDetailEntity{
		ID:          modelServiceDetail.ID,
		ServiceID:   modelServiceDetail.ServiceID,
		PathImage:   modelServiceDetail.PathImage,
		Description: modelServiceDetail.Description,
		PathPdf:     modelServiceDetail.PathPdf,
		PathDocx:    modelServiceDetail.PathDocx,
		Title:       modelServiceDetail.Title,
	}, nil
}

func (s *serviceDetail) GetByServiceIDDetail(ctx context.Context, serviceId int64) (*entity.ServiceDetailEntity, error) {
	rows, err := s.DB.Table("service_details as ack").
		Select("ack.id", "ack.path_image", "ack.description", "ack.path_pdf", "ack.path_docx", "ac.name").
		Joins("inner join service_sections as ac on ac.id = ack.service_id").
		Where("ack.deleted_at IS NULL").
		Rows()

	if err != nil {
		log.Errorf("[REPOSITORY] GetByServiceIDDetail - 1: %v", err)
		return nil, err
	}

	serviceDetail := entity.ServiceDetailEntity{}
	for rows.Next() {
		err = rows.Scan(&serviceDetail.ID, &serviceDetail.PathImage, &serviceDetail.Description, &serviceDetail.PathPdf, &serviceDetail.PathDocx, &serviceDetail.ServiceName)
		if err != nil {
			log.Errorf("[REPOSITORY] GetByServiceIDDetail - 2: %v", err)
			return nil, err
		}
	}
	return &serviceDetail, nil
}

func NewServiceDetailRepository(DB *gorm.DB) ServiceDetailInterface {
	return &serviceDetail{
		DB: DB,
	}
}
