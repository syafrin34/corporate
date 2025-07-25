package service

import (
	"context"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"
)

type FaqSectionServiceInterface interface {
	CreateFaqSection(ctx context.Context, req entity.FaqSectionEntity) error
	FetchAllFaqSection(ctx context.Context) ([]entity.FaqSectionEntity, error)
	FetchByIDFaqSection(ctx context.Context, id int64) (*entity.FaqSectionEntity, error)
	EditByIDFaqSection(ctx context.Context, req entity.FaqSectionEntity) error
	DeleteByIDFaqSection(ctx context.Context, id int64) error
}

type faqSectionService struct {
	faqSectionRepo repository.FaqSectionInterface
}

// CreateFaqSection implements FaqSectionServiceInterface.
func (f *faqSectionService) CreateFaqSection(ctx context.Context, req entity.FaqSectionEntity) error {
	return f.faqSectionRepo.CreateFaqSection(ctx, req)
}

// DeleteByIDFaqSection implements FaqSectionServiceInterface.
func (f *faqSectionService) DeleteByIDFaqSection(ctx context.Context, id int64) error {
	return f.faqSectionRepo.DeleteByIDFaqSection(ctx, id)
}

// EditByIDFaqSection implements FaqSectionServiceInterface.
func (f *faqSectionService) EditByIDFaqSection(ctx context.Context, req entity.FaqSectionEntity) error {
	return f.faqSectionRepo.EditByIDFaqSection(ctx, req)
}

// FetchAllFaqSection implements FaqSectionServiceInterface.
func (f *faqSectionService) FetchAllFaqSection(ctx context.Context) ([]entity.FaqSectionEntity, error) {
	return f.faqSectionRepo.FetchAllFaqSection(ctx)
}

// FetchByIDFaqSection implements FaqSectionServiceInterface.
func (f *faqSectionService) FetchByIDFaqSection(ctx context.Context, id int64) (*entity.FaqSectionEntity, error) {
	return f.faqSectionRepo.FetchByIDFaqSection(ctx, id)
}

func NewFaqSectionService(faqSectionRepo repository.FaqSectionInterface) FaqSectionServiceInterface {
	return &faqSectionService{
		faqSectionRepo: faqSectionRepo,
	}
}
