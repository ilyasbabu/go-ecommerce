package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Products struct {
	BaseModel
	Name        string
	Description string
	Price       float64 `gorm:"type:decimal(7,2)"`
	Stock       int     `gorm:"default:0"`
	AvgRating   float64 `gorm:"type:decimal(3,1);default:0"`
	RatingCount int     `gorm:"default:0"`
	Slug        string  `gorm:"size:200;unique"`
}

func (p *Products) BeforeCreate(tx *gorm.DB) (err error) {
	p.Slug = slug.Make(p.Name)
	return nil
}
