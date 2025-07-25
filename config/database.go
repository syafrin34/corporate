package config

import (
	//	"corporate/database/migrations/seeds"
	"corporate/database/migrations/seeds"
	"corporate/internal/core/domain/model"
	"fmt"

	"github.com/rs/zerolog/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*Postgres, error) {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Psql.User,
		cfg.Psql.Password,
		cfg.Psql.Host,
		cfg.Psql.Port,
		cfg.Psql.DBName)
	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})

	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-1] failed to connect to database " + cfg.Psql.Host)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-1] failed to get database ")
		return nil, err
	}

	// if err := db.AutoMigrate(&model.User{}, &model); err != nil {
	// 	log.Fatal().Err(err).Msg("AutoMigrate failed")
	// 	return nil, err
	// }

	if err := runMigration(db); err != nil {
		log.Fatal().Err(err).Msg("Migration failed")
		return nil, err
	}

	seeds.SeedAdmin(db)

	sqlDB.SetMaxOpenConns(cfg.Psql.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.Psql.DBMaxIdle)

	return &Postgres{DB: db}, nil

}

func runMigration(db *gorm.DB) error {
	log.Info().Msg("Running AutoMigrate...")

	return db.AutoMigrate(
		&model.User{},
		&model.AboutCompany{},
		&model.AboutCompanyKeynote{},
		&model.Appointment{},
		&model.ClientSection{},
		&model.FAQSection{},
		&model.HeroSection{},
		&model.OurTeam{},
		&model.PortoFolioSection{},
		&model.PortofolioDetail{},
		&model.PortofolioTestimonial{},
		&model.ServiceSection{},
		// Tambahkan model lain di sini sesuai kebutuhan
	)
}
