package models

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Dress struct {
	gorm.Model
	Url         string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:varchar(255);not null"`
	Labels      pq.StringArray `json:"labels" gorm:"type:text[]"`
	Weather     float64        `gorm:"type:decimal(10,2);not null"`
	Location    datatypes.JSON `json:"location" gorm:"type:jsonb"`
}

type Location struct {
	Vertical   string
	Horizontal string
	Locale     string
}
