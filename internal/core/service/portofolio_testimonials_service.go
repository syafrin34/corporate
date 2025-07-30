package service

import (
	"context"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"

	"github.com/labstack/gommon/log"
)

type PortofolioTestimonialServiceInterface interface {
	CreatePortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error
	FetchAllPortofolioTestimonial(ctx context.Context) ([]entity.PortofolioTestimonialEntity, error)
	FetchByIDPortofolioTestimonial(ctx context.Context, id int64) (*entity.PortofolioTestimonialEntity, error)
	EditByIDPortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error
	DeleteByIDPortofolioTestimonial(ctx context.Context, id int64) error
}

type portofolioTestimonialService struct {
	portofolioTestimonial repository.PortofolioTestimonialInterface
	portofolioSectionRepo repository.PortofolioSectionInterface
}

// CreatePortofolioTestimonial implements PortofolioTestimonialServiceInterface.
func (p *portofolioTestimonialService) CreatePortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error {
	_, err := p.portofolioSectionRepo.FetchByIDPortofolioSection(ctx, req.PortoFolioSection.ID)
	if err != nil {
		log.Errorf("[SERVICE] CreatePortofolioTestimonial - 1: %v", err)
		return err
	}
	return p.portofolioTestimonial.CreatePortofolioTestimonial(ctx, req)
}

// DeleteByIDPortofolioTestimonial implements PortofolioTestimonialServiceInterface.
func (p *portofolioTestimonialService) DeleteByIDPortofolioTestimonial(ctx context.Context, id int64) error {
	return p.portofolioTestimonial.DeleteByIDPortofolioTestimonial(ctx, id)
}

// EditByIDPortofolioTestimonial implements PortofolioTestimonialServiceInterface.
func (p *portofolioTestimonialService) EditByIDPortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error {
	_, err := p.portofolioSectionRepo.FetchByIDPortofolioSection(ctx, req.PortoFolioSection.ID)
	if err != nil {
		log.Errorf("[SERVICE] EditIDPortofolioTestimonial - 2: %v", err)
		return err
	}
	return p.portofolioTestimonial.EditByIDPortofolioTestimonial(ctx, req)
}

// FetchAllPortofolioTestimonial implements PortofolioTestimonialServiceInterface.
func (p *portofolioTestimonialService) FetchAllPortofolioTestimonial(ctx context.Context) ([]entity.PortofolioTestimonialEntity, error) {
	return p.portofolioTestimonial.FetchAllPortofolioTestimonial(ctx)
}

// FetchByIDPortofolioTestimonial implements PortofolioTestimonialServiceInterface.
func (p *portofolioTestimonialService) FetchByIDPortofolioTestimonial(ctx context.Context, id int64) (*entity.PortofolioTestimonialEntity, error) {
	return p.portofolioTestimonial.FetchByIDPortofolioTestimonial(ctx, id)
}

func NewPortofolioTestimonialService(portofolioTestimonial repository.PortofolioTestimonialInterface, portofolioSectionRepo repository.PortofolioSectionInterface) PortofolioTestimonialServiceInterface {
	return &portofolioTestimonialService{
		portofolioTestimonial: portofolioTestimonial,
		portofolioSectionRepo: portofolioSectionRepo,
	}
}
