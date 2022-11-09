package database

import (
	"fmt"
	"waysbeans/models"
	"waysbeans/pkg/postgre"
)

func RunMigration() {
	if err := postgre.DB.AutoMigrate(&models.Products{}, &models.User{}); err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("migration success")
}