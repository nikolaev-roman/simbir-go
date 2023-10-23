package models

import (
	"github.com/google/uuid"
)

type Account struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Username string    `json:"username" gorm:"unique"`
	Password string    `json:"password"`
}
