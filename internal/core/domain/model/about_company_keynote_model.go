package model

import (
	"time"

	"gorm.io/gorm"
)

type AboutCompanyKeynote struct {
	ID             int64 `gorm:"id,primaryKey"`
	AboutCompanyID string
	Keynote        string
	PathImage      *string
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
