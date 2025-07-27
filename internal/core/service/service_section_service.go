package service

import (
	"context"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"
)

type ServiceSectionServiceInterface interface {
	CreateServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error
	FetchAllServiceSection(ctx context.Context) ([]entity.ServiceSectionEntity, error)
	FetchByIDServiceSection(ctx context.Context, id int64) (*entity.ServiceSectionEntity, error)
	EditByIDServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error
	DeleteByIDServiceSection(ctx context.Context, id int64) error
}

type serviceSectionService struct {
	serviceSectionRepo repository.ServiceSectionInterface
}

// CreateServiceSection implements ServiceSectionServiceInterface.
func (s *serviceSectionService) CreateServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error {
	return s.serviceSectionRepo.CreateServiceSection(ctx, req)
}

// DeleteByIDServiceSection implements ServiceSectionServiceInterface.
func (s *serviceSectionService) DeleteByIDServiceSection(ctx context.Context, id int64) error {
	return s.serviceSectionRepo.DeleteByIDServiceSection(ctx, id)
}

// EditByIDServiceSection implements ServiceSectionServiceInterface.
func (s *serviceSectionService) EditByIDServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error {
	return s.serviceSectionRepo.EditByIDServiceSection(ctx, req)
}

// FetchAllServiceSection implements ServiceSectionServiceInterface.
func (s *serviceSectionService) FetchAllServiceSection(ctx context.Context) ([]entity.ServiceSectionEntity, error) {
	return s.serviceSectionRepo.FetchAllServiceSection(ctx)
}

// FetchByIDServiceSection implements ServiceSectionServiceInterface.
func (s *serviceSectionService) FetchByIDServiceSection(ctx context.Context, id int64) (*entity.ServiceSectionEntity, error) {
	return s.serviceSectionRepo.FetchByIDServiceSection(ctx, id)
}

func NewServiceSectionService(serviceSectionRepo repository.ServiceSectionInterface) ServiceSectionServiceInterface {
	return &serviceSectionService{
		serviceSectionRepo: serviceSectionRepo,
	}
}
