package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	Name      string          `json:"name"`
	Price     float64         `json:"price"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt  time.Time       `json:"-"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}
