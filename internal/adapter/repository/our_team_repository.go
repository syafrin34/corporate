package repository

import (
	"context"
	"corporate/internal/core/domain/entity"
	"corporate/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type OurTeamInterface interface {
	CreateOurTeam(ctx context.Context, req entity.OurTeamEntity) error
	FetchAllOurTeam(ctx context.Context) ([]entity.OurTeamEntity, error)
	FetchByIDOurTeam(ctx context.Context, id int64) (*entity.OurTeamEntity, error)
	EditByIDOurTeam(ctx context.Context, req entity.OurTeamEntity) error
	DeleteByIDOurTeam(ctx context.Context, id int64) error
}

type ourTeam struct {
	DB *gorm.DB
}

// CreateOurTeam implements OurTeamInterface.
func (o *ourTeam) CreateOurTeam(ctx context.Context, req entity.OurTeamEntity) error {
	modelOurTeam := model.OurTeam{
		Name:      req.Name,
		Role:      req.PathPhoto,
		PathPhoto: req.PathPhoto,
		Tagline:   req.Tagline,
	}
	if err = o.DB.Create(&modelOurTeam).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateOurTeam - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDOurTeam implements OurTeamInterface.
func (o *ourTeam) DeleteByIDOurTeam(ctx context.Context, id int64) error {
	modelOurTeam := model.OurTeam{}
	err = o.DB.Where("id = ?", id).First(&modelOurTeam).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDOurTeam - 1: %v", err)
		return err
	}
	err = o.DB.Delete(&modelOurTeam).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDOurTeam -2: %v", err)
		return err
	}
	return nil

}

// EditByIDOurTeam implements OurTeamInterface.
func (o *ourTeam) EditByIDOurTeam(ctx context.Context, req entity.OurTeamEntity) error {
	modelOurTeam := model.OurTeam{}
	err = o.DB.Where("id = ?", req.ID).First(&modelOurTeam).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDOurTeam - 1: %v", err)
		return err
	}
	modelOurTeam.Name = req.Name
	modelOurTeam.Role = req.Role
	modelOurTeam.PathPhoto = req.PathPhoto
	modelOurTeam.Tagline = req.Tagline

	err = o.DB.Save(&modelOurTeam).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDOurTeam - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllOurTeam implements OurTeamInterface.
func (o *ourTeam) FetchAllOurTeam(ctx context.Context) ([]entity.OurTeamEntity, error) {
	modelOurTeam := []model.OurTeam{}
	err = o.DB.Select("id", "name", "role", "path_photo", "tag_line").Find(&modelOurTeam).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllOurTeam - 1: %v", err)
		return nil, err
	}

	var ourTeamEntities []entity.OurTeamEntity
	for _, v := range modelOurTeam {
		ourTeamEntities = append(ourTeamEntities, entity.OurTeamEntity{
			ID:        v.ID,
			Name:      v.Name,
			Role:      v.Role,
			PathPhoto: v.PathPhoto,
			Tagline:   v.Tagline,
		})
	}
	return ourTeamEntities, nil
}

// FetchByIDOurTeam implements OurTeamInterface.
func (o *ourTeam) FetchByIDOurTeam(ctx context.Context, id int64) (*entity.OurTeamEntity, error) {
	modelOurTeam := model.OurTeam{}
	err = o.DB.Select("id", "name", "role", "path_photo", "tag_line").Where("id = ?", id).First(&modelOurTeam).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditIDOurTeam - 1: %v", err)
		return nil, err
	}

	return &entity.OurTeamEntity{
		ID:        modelOurTeam.ID,
		Name:      modelOurTeam.Name,
		PathPhoto: modelOurTeam.PathPhoto,
		Tagline:   modelOurTeam.Tagline,
	}, nil
}

func NewOurTeamRepository(DB *gorm.DB) OurTeamInterface {
	return &ourTeam{
		DB: DB,
	}
}
