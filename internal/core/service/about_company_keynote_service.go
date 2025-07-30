package service

import (
	"context"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"

	"github.com/labstack/gommon/log"
)

type AboutCompanyKeynoteServiceInterface interface {
	CreateAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error
	FetchAllAboutCompanyKeynote(ctx context.Context) ([]entity.AboutCompanyKeynoteEntity, error)
	FetchByIDAboutCompanyKeynote(ctx context.Context, id int64) (*entity.AboutCompanyKeynoteEntity, error)
	EditByIDAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error
	DeleteByIDAboutCompanyKeynote(ctx context.Context, id int64) error
	FetchByCompanyID(ctx context.Context, companyId int64) ([]entity.AboutCompanyKeynoteEntity, error)
}

type aboutCompanyKeynoteService struct {
	aboutCompanyKeynoteRepo repository.AboutCompanyKeynoteInterface
	aboutCompanyRepo        repository.AboutCompanyInterface
}

// FetchByCompanyID implements AboutCompanyKeynoteServiceInterface.
func (a *aboutCompanyKeynoteService) FetchByCompanyID(ctx context.Context, companyId int64) ([]entity.AboutCompanyKeynoteEntity, error) {
	return a.aboutCompanyKeynoteRepo.FetchByCompanyID(ctx, companyId)
}

// CreateAboutCompanyKeynote implements AboutCompanyKeynoteServiceInterface.
func (a *aboutCompanyKeynoteService) CreateAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error {
	_, err := a.aboutCompanyRepo.FetchByIDAboutCompany(ctx, req.AboutCompanyID)
	if err != nil {
		log.Errorf("[SERVICE] CreateAboutCompanyKeynote - 1: %v", err)
		return err
	}

	return a.aboutCompanyKeynoteRepo.CreateAboutCompanyKeynote(ctx, req)
}

// DeleteByIDAboutCompanyKeynote implements AboutCompanyKeynoteServiceInterface.
func (a *aboutCompanyKeynoteService) DeleteByIDAboutCompanyKeynote(ctx context.Context, id int64) error {

	return a.aboutCompanyKeynoteRepo.DeleteByIDAboutCompanyKeynote(ctx, id)
}

// EditByIDAboutCompanyKeynote implements AboutCompanyKeynoteServiceInterface.
func (a *aboutCompanyKeynoteService) EditByIDAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error {
	_, err := a.aboutCompanyRepo.FetchByIDAboutCompany(ctx, req.AboutCompanyID)
	if err != nil {
		log.Errorf("[SERVICE] EditByIDAboutCompanyKeynote - 2: %v", err)
		return err
	}
	return a.aboutCompanyKeynoteRepo.EditByIDAboutCompanyKeynote(ctx, req)
}

// FetchAllAboutCompanyKeynote implements AboutCompanyKeynoteServiceInterface.
func (a *aboutCompanyKeynoteService) FetchAllAboutCompanyKeynote(ctx context.Context) ([]entity.AboutCompanyKeynoteEntity, error) {
	return a.aboutCompanyKeynoteRepo.FetchAllAboutCompanyKeynote(ctx)
}

// FetchByIDAboutCompanyKeynote implements AboutCompanyKeynoteServiceInterface.
func (a *aboutCompanyKeynoteService) FetchByIDAboutCompanyKeynote(ctx context.Context, id int64) (*entity.AboutCompanyKeynoteEntity, error) {
	return a.aboutCompanyKeynoteRepo.FetchByIDAboutCompanyKeynote(ctx, id)
}

func NewAboutCompanyKeynoteService(aboutCompanyKeynoteRepo repository.AboutCompanyKeynoteInterface, aboutCompanyRepo repository.AboutCompanyInterface) AboutCompanyKeynoteServiceInterface {
	return &aboutCompanyKeynoteService{
		aboutCompanyKeynoteRepo: aboutCompanyKeynoteRepo,
		aboutCompanyRepo:        aboutCompanyRepo,
	}
}
