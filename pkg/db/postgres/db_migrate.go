package postgres

import (
	"fmt"

	"github.com/nikolaev-roman/simbir-go/internal/models"
	"gorm.io/gorm"
)

func MigrateDb(DB *gorm.DB) {
	fmt.Println("migrate")
	DB.AutoMigrate(&models.Account{})
	DB.AutoMigrate(&models.Transport{})
}
