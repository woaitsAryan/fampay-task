package initializers

import (
	"fmt"

	"github.com/woaitsAryan/fampay-task/backend/models"
)

func AutoMigrate() {
	fmt.Println("\nStarting Migrations...")
	DB.AutoMigrate(
		&models.Video{},

	)
	fmt.Println("Migrations Finished!")
}
