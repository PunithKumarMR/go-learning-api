package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UID       uint           `json:"user_id" gorm:"primarykey"`
	Email     string         `json:"email" gorm:"unique" validate:"required,email"`
	Password  string         `json:"-" validate:"required, min=6"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
