package service

import (
	"context"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"
)

type PortofolioSectionServiceInterface interface {
	CreatePortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error
	FetchAllPortofolioSection(ctx context.Context) ([]entity.PortofolioSectionEntity, error)
	FetchByIDPortofolioSection(ctx context.Context, id int64) (*entity.PortofolioSectionEntity, error)
	EditByIDPortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error
	DeleteByIDPortofolioSection(ctx context.Context, id int64) error
}

type portofolioSectionService struct {
	serviceSectionRepo repository.PortofolioSectionInterface
}

// CreatePortofolioSection implements PortofolioSectionServiceInterface.
func (p *portofolioSectionService) CreatePortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error {
	return p.serviceSectionRepo.CreatePortofolioSection(ctx, req)
}

// DeleteByIDPortofolioSection implements PortofolioSectionServiceInterface.
func (p *portofolioSectionService) DeleteByIDPortofolioSection(ctx context.Context, id int64) error {
	return p.serviceSectionRepo.DeleteByIDPortofolioSection(ctx, id)
}

// EditByIDPortofolioSection implements PortofolioSectionServiceInterface.
func (p *portofolioSectionService) EditByIDPortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error {
	return p.serviceSectionRepo.EditByIDPortofolioSection(ctx, req)
}

// FetchAllPortofolioSection implements PortofolioSectionServiceInterface.
func (p *portofolioSectionService) FetchAllPortofolioSection(ctx context.Context) ([]entity.PortofolioSectionEntity, error) {
	return p.serviceSectionRepo.FetchAllPortofolioSection(ctx)
}

// FetchByIDPortofolioSection implements PortofolioSectionServiceInterface.
func (p *portofolioSectionService) FetchByIDPortofolioSection(ctx context.Context, id int64) (*entity.PortofolioSectionEntity, error) {
	return p.serviceSectionRepo.FetchByIDPortofolioSection(ctx, id)
}

func NewPortofolioSectionService(serviceSectionRepo repository.PortofolioSectionInterface) PortofolioSectionServiceInterface {
	return &portofolioSectionService{
		serviceSectionRepo: serviceSectionRepo,
	}
}
