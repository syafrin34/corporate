package service

import (
	"context"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/domain/entity"
)

type OurTeamServiceInterface interface {
	CreateOurTeam(ctx context.Context, req entity.OurTeamEntity) error
	FetchAllOurTeam(ctx context.Context) ([]entity.OurTeamEntity, error)
	FetchByIDOurTeam(ctx context.Context, id int64) (*entity.OurTeamEntity, error)
	EditByIDOurTeam(ctx context.Context, req entity.OurTeamEntity) error
	DeleteByIDOurTeam(ctx context.Context, id int64) error
}

type ourTeamService struct {
	ourTeamRepo repository.OurTeamInterface
}

// CreateOurTeam implements OurTeamServiceInterface.
func (o *ourTeamService) CreateOurTeam(ctx context.Context, req entity.OurTeamEntity) error {
	return o.ourTeamRepo.CreateOurTeam(ctx, req)
}

// DeleteByIDOurTeam implements OurTeamServiceInterface.
func (o *ourTeamService) DeleteByIDOurTeam(ctx context.Context, id int64) error {
	return o.ourTeamRepo.DeleteByIDOurTeam(ctx, id)
}

// EditByIDOurTeam implements OurTeamServiceInterface.
func (o *ourTeamService) EditByIDOurTeam(ctx context.Context, req entity.OurTeamEntity) error {
	return o.ourTeamRepo.EditByIDOurTeam(ctx, req)
}

// FetchAllOurTeam implements OurTeamServiceInterface.
func (o *ourTeamService) FetchAllOurTeam(ctx context.Context) ([]entity.OurTeamEntity, error) {
	return o.ourTeamRepo.FetchAllOurTeam(ctx)
}

// FetchByIDOurTeam implements OurTeamServiceInterface.
func (o *ourTeamService) FetchByIDOurTeam(ctx context.Context, id int64) (*entity.OurTeamEntity, error) {
	return o.ourTeamRepo.FetchByIDOurTeam(ctx, id)
}

func NewOurTeamService(ourTeamRepo repository.OurTeamInterface) OurTeamServiceInterface {
	return &ourTeamService{
		ourTeamRepo: ourTeamRepo,
	}
}
