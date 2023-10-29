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
	DB.AutoMigrate(&models.Rent{})
	DB.Create(&models.Account{
		Username: "admin",
		Password: "$2a$10$sxDWQ4ABVVr02j4hAv8GVedZ9Jtr4o837kDdagZ47XyP/sH0EMa9m",
		IsAdmin:  true,
		Balance:  0,
	})
}
