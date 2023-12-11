package db

import (
	"fmt"
	"glamgrove/pkg/config"
	"glamgrove/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func to connect data base using config(database config) and return address of a new instnce of gorm DB
func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	fmt.Println("Connection string", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	fmt.Println("Error is", err)

	// migrate the database tables
	db.AutoMigrate(
		//user
		domain.User{},

		//admin
		domain.Admin{},

		//product
		domain.Product{},
		

		domain.Category{},
	)

	return db, err
}
