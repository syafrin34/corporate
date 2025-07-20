package service

import (
	"context"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"
)

type HeroSectionServiceInterface interface {
	CreateHeroSection(ctx context.Context, req entity.HeroSectionEntity) error
	FetchAllHeroSection(ctx context.Context) ([]entity.HeroSectionEntity, error)
	FetchByIDHeroSection(ctx context.Context, id int64) (*entity.HeroSectionEntity, error)
	EditByIDHeroSection(ctx context.Context, req entity.HeroSectionEntity) error
	DeleteByIDHeroSection(ctx context.Context, id int64) error
}

type heroSectionService struct {
	heroSectionRepo repository.HeroSectionInterface
}

// CreateHeroSection implements HeroSectionServiceInterface.
func (h *heroSectionService) CreateHeroSection(ctx context.Context, req entity.HeroSectionEntity) error {
	panic("unimplemented")
}

// DeleteByIDHeroSection implements HeroSectionServiceInterface.
func (h *heroSectionService) DeleteByIDHeroSection(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// EditByIDHeroSection implements HeroSectionServiceInterface.
func (h *heroSectionService) EditByIDHeroSection(ctx context.Context, req entity.HeroSectionEntity) error {
	panic("unimplemented")
}

// FetchAllHeroSection implements HeroSectionServiceInterface.
func (h *heroSectionService) FetchAllHeroSection(ctx context.Context) ([]entity.HeroSectionEntity, error) {
	panic("unimplemented")
}

// FetchByIDHeroSection implements HeroSectionServiceInterface.
func (h *heroSectionService) FetchByIDHeroSection(ctx context.Context, id int64) (*entity.HeroSectionEntity, error) {
	panic("unimplemented")
}

func NewHeroSectionService(heroSectionRepo repository.HeroSectionInterface) HeroSectionServiceInterface {
	return &heroSectionService{
		heroSectionRepo: heroSectionRepo,
	}
}
