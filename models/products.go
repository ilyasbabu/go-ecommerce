package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Products struct {
	BaseModel
	Name        string
	Description string
	Price       float64
	Stock       int
	AvgRating   float64
	RatingCount int
	Slug        string `gorm:"size:200;unique"`
}

func (p *Products) BeforeCreate(tx *gorm.DB) (err error) {
	p.Slug = slug.Make(p.Name)
	return nil
}
