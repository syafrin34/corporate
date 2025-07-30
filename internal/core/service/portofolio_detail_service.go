package service

import (
	"context"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"

	"github.com/labstack/gommon/log"
)

type PortofolioDetailServiceInterface interface {
	CreatePortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error
	FetchAllPortofolioDetail(ctx context.Context) ([]entity.PortofolioDetailEntity, error)
	FetchByIDPortofolioDetail(ctx context.Context, id int64) (*entity.PortofolioDetailEntity, error)
	EditByIDPortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error
	DeleteByIDPortofolioDetail(ctx context.Context, id int64) error
}

type portofolioDetailService struct {
	portofolioDetail      repository.PortofolioDetailInterface
	portofolioSectionRepo repository.PortofolioSectionInterface
}

// CreatePortofolioDetail implements PortofolioDetailServiceInterface.
func (p *portofolioDetailService) CreatePortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error {
	_, err := p.portofolioSectionRepo.FetchByIDPortofolioSection(ctx, req.PortofolioSection.ID)
	if err != nil {
		log.Errorf("[REPOSITORY] CreatePortofolioDetail - 1: %v", err)
		return err
	}
	return p.portofolioDetail.CreatePortofolioDetail(ctx, req)
}

// DeleteByIDPortofolioDetail implements PortofolioDetailServiceInterface.
func (p *portofolioDetailService) DeleteByIDPortofolioDetail(ctx context.Context, id int64) error {
	return p.portofolioDetail.DeleteByIDPortofolioDetail(ctx, id)
}

// EditByIDPortofolioDetail implements PortofolioDetailServiceInterface.
func (p *portofolioDetailService) EditByIDPortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error {
	_, err := p.portofolioSectionRepo.FetchByIDPortofolioSection(ctx, req.PortofolioSection.ID)
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDPortofolioDetail - 1: %v", err)
		return err
	}
	return p.portofolioDetail.EditByIDPortofolioDetail(ctx, req)
}

// FetchAllPortofolioDetail implements PortofolioDetailServiceInterface.
func (p *portofolioDetailService) FetchAllPortofolioDetail(ctx context.Context) ([]entity.PortofolioDetailEntity, error) {
	return p.portofolioDetail.FetchAllPortofolioDetail(ctx)
}

// FetchByIDPortofolioDetail implements PortofolioDetailServiceInterface.
func (p *portofolioDetailService) FetchByIDPortofolioDetail(ctx context.Context, id int64) (*entity.PortofolioDetailEntity, error) {
	return p.portofolioDetail.FetchByIDPortofolioDetail(ctx, id)
}

func NewPortofolioDetailService(portofolioDetail repository.PortofolioDetailInterface, portofolioSectionRepo repository.PortofolioSectionInterface) PortofolioDetailServiceInterface {
	return &portofolioDetailService{
		portofolioDetail:      portofolioDetail,
		portofolioSectionRepo: portofolioSectionRepo,
	}
}
