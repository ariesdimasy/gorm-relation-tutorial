package migrations

import (
	"fmt"
	"gorm-relation-tutorial/config"
	"gorm-relation-tutorial/models"
)

func Migration() {
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Locker{},
		&models.Post{},
	)

	if err != nil {
		fmt.Println("cant running migration")
	} else {
		fmt.Println("migrated.")
	}

}
