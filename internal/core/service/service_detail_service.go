package service

import (
	"context"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"
)

type ServiceDetailServiceInterface interface {
	CreateServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error
	FetchAllServiceDetail(ctx context.Context) ([]entity.ServiceDetailEntity, error)
	FetchByIDServiceDetail(ctx context.Context, id int64) (*entity.ServiceDetailEntity, error)
	EditByIDServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error
	DeleteByIDServiceDetail(ctx context.Context, id int64) error
	GetByServiceIDDetail(ctx context.Context, serviceId int64) (*entity.ServiceDetailEntity, error)
}

type serviceDetailService struct {
	ServiceDetailRepo repository.ServiceDetailInterface
}

// GetByServiceIDDetail implements ServiceDetailServiceInterface.
func (s *serviceDetailService) GetByServiceIDDetail(ctx context.Context, serviceId int64) (*entity.ServiceDetailEntity, error) {
	return s.ServiceDetailRepo.GetByServiceIDDetail(ctx, serviceId)
}

// CreateServiceDetail implements ServiceDetailServiceInterface.
func (s *serviceDetailService) CreateServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error {
	return s.ServiceDetailRepo.CreateServiceDetail(ctx, req)
}

// DeleteByIDServiceDetail implements ServiceDetailServiceInterface.
func (s *serviceDetailService) DeleteByIDServiceDetail(ctx context.Context, id int64) error {
	return s.ServiceDetailRepo.DeleteByIDServiceDetail(ctx, id)
}

// EditByIDServiceDetail implements ServiceDetailServiceInterface.
func (s *serviceDetailService) EditByIDServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error {
	return s.ServiceDetailRepo.EditByIDServiceDetail(ctx, req)
}

// FetchAllServiceDetail implements ServiceDetailServiceInterface.
func (s *serviceDetailService) FetchAllServiceDetail(ctx context.Context) ([]entity.ServiceDetailEntity, error) {
	return s.ServiceDetailRepo.FetchAllServiceDetail(ctx)
}

// FetchByIDServiceDetail implements ServiceDetailServiceInterface.
func (s *serviceDetailService) FetchByIDServiceDetail(ctx context.Context, id int64) (*entity.ServiceDetailEntity, error) {
	return s.ServiceDetailRepo.FetchByIDServiceDetail(ctx, id)
}

func NewServiceDetailService(ServiceDetailRepo repository.ServiceDetailInterface) ServiceDetailServiceInterface {
	return &serviceDetailService{
		ServiceDetailRepo: ServiceDetailRepo,
	}
}
