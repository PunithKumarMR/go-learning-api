package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name" validate:"required,min=2,max=100"`
	Price     float64        `json:"price" validate:"required,gt=0"`
	Quantity  int            `json:"quantity" validate:"required,gte=0"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
