package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"id,primaryKey"`
	Name      string         `gorm:"name"`
	Email     string         `gorm:"email"`
	Password  string         `gorm:"password"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt *time.Time     `gorm:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
