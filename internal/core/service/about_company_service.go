package service

import (
	"context"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"
)

type AboutCompanyServiceInterface interface {
	CreateAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error
	FetchAllAboutCompany(ctx context.Context) ([]entity.AboutCompanyEntity, error)
	FetchByIDAboutCompany(ctx context.Context, id int64) (*entity.AboutCompanyEntity, error)
	EditByIDAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error
	DeleteByIDAboutCompany(ctx context.Context, id int64) error
}

type aboutCompanyService struct {
	aboutCompanyRepo repository.AboutCompanyInterface
}

// CreateAboutCompany implements AboutCompanyServiceInterface.
func (a *aboutCompanyService) CreateAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error {
	return a.aboutCompanyRepo.CreateAboutCompany(ctx, req)
}

// DeleteByIDAboutCompany implements AboutCompanyServiceInterface.
func (a *aboutCompanyService) DeleteByIDAboutCompany(ctx context.Context, id int64) error {
	return a.aboutCompanyRepo.DeleteByIDAboutCompany(ctx, id)
}

// EditByIDAboutCompany implements AboutCompanyServiceInterface.
func (a *aboutCompanyService) EditByIDAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error {
	return a.aboutCompanyRepo.EditByIDAboutCompany(ctx, req)
}

// FetchAllAboutCompany implements AboutCompanyServiceInterface.
func (a *aboutCompanyService) FetchAllAboutCompany(ctx context.Context) ([]entity.AboutCompanyEntity, error) {
	return a.aboutCompanyRepo.FetchAllAboutCompany(ctx)
}

// FetchByIDAboutCompany implements AboutCompanyServiceInterface.
func (a *aboutCompanyService) FetchByIDAboutCompany(ctx context.Context, id int64) (*entity.AboutCompanyEntity, error) {
	return a.aboutCompanyRepo.FetchByIDAboutCompany(ctx, id)
}

func NewAboutCompanyService(aboutCompanyRepo repository.AboutCompanyInterface) AboutCompanyServiceInterface {
	return &aboutCompanyService{
		aboutCompanyRepo: aboutCompanyRepo,
	}
}
