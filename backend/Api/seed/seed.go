package seed

import (
	"api/backend/api/models"
	"log"

	"github.com/jinzhu/gorm"
)

func Load(db *gorm.DB) {
	// Drop the todos table if it exists
	err := db.Debug().DropTableIfExists(&models.Todo{}).Error
	if err != nil {
		// Log the error if it occurs
		log.Fatalf("cannot drop table: %v", err)
	}
	// Create the todos table
	err = db.Debug().AutoMigrate(&models.Todo{}).Error
	if err != nil {
		// Log the error if it occurs
		log.Fatalf("cannot migrate table: %v", err)
	}

	// Create two todos
	todos := []models.Todo{
		{
			ID:        1,
			Item:      "I have to find dance class near me",
			Completed: false,
		},
		{
			Item:      "I have to learn coding",
			ID:        2,
			Completed: false,
		},
	}

	// Create the todo in the database
	for i, _ := range todos {
		err = db.Debug().Model(&models.Todo{}).Create(&todos[i]).Error
		if err != nil {
			// Log the error if it occurs
			log.Fatalf("cannot seed todos table: %v", err)
		}
	}
}
