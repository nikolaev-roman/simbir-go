package models

import (
	"time"

	"github.com/google/uuid"
)

type Rent struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey" validate:"omitempty,uuid"`
	TransportID uuid.UUID `json:"transportId" gorm:"type:uuid;not null" validate:"required,uuid"`
	UserID      uuid.UUID `json:"userId" gorm:"type:uuid;not null" validate:"required,uuid"`
	TimeStart   time.Time `json:"timeStart" gorm:"not null" validate:"required"`
	TimeEnd     time.Time `json:"timeEnd,omitempty"`
	PriceOfUnit float64   `json:"priceOfUnit" gorm:"not null" validate:"required"`
	PriceType   string    `json:"priceType" gorm:"not null" validate:"required"`
	FinalPrice  float64   `json:"finalPrice"`
}

type SearchToRent struct {
	Lat    float64 `json:"lat"`
	Long   float64 `json:"long"`
	Radius float64 `json:"radius"`
	Type   string  `json:"type"`
}

func (p *SearchToRent) IsValidType() bool {
	switch p.Type {
	case "Car", "Bike", "Scooter", "All":
		return true
	}

	return false
}
