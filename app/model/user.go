package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"primary_key, not null, autoIncrement" json:"id"  validate:"uuid,omitempty"`
	Name      string         `json:"name" validate:"required"`
	Email     string         `gorm:"unique" json:"email" validate:"required,email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
