package models

import (
	"time"

	"gorm.io/gorm"
)

type AirPolution struct {
	gorm.Model
	Id        int       `gorm:"primaryKey" json:"id"`
	Lat       float64   `gorm:"type:float; not null" json:"lat"`
	Lng       float64   `gorm:"type:float; not null" json:"lng"`
	Compound  string    `gorm:"type:varchar(255); not null" json:"compound"`
	CreatedAt time.Time `json:"created_at"`
}
