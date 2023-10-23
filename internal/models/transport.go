package models

import "github.com/google/uuid"

type Transport struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	OwnerID       uuid.UUID `json:"ownerId" gorm:"type:uuid"`
	CanBeRented   bool      `json:"canBeRented"`
	TransportType string    `json:"transportType"`
	Model         string    `json:"model"`
	Color         string    `json:"color"`
	Identifier    string    `json:"identifier"`
	Description   string    `json:"description"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	MinutePrice   float64   `json:"minutePrice"`
	DayPrice      float64   `json:"dayPrice"`
}
