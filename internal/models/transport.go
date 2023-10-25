package models

import "github.com/google/uuid"

type Transport struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey" validate:"omitempty,uuid"`
	OwnerID       uuid.UUID `json:"ownerId" gorm:"type:uuid" validate:"required"`
	CanBeRented   bool      `json:"canBeRented" gorm:"not null" validate:"required"`
	TransportType string    `json:"transportType" validate:"required"`
	Model         string    `json:"model" gorm:"not null" validate:"required"`
	Color         string    `json:"color" gorm:"not null" validate:"required"`
	Identifier    string    `json:"identifier" gorm:"not null" validate:"required"`
	Description   string    `json:"description"`
	Latitude      float64   `json:"latitude" gorm:"not null" validate:"required"`
	Longitude     float64   `json:"longitude" gorm:"not null" validate:"required"`
	MinutePrice   float64   `json:"minutePrice"`
	DayPrice      float64   `json:"dayPrice"`
}
