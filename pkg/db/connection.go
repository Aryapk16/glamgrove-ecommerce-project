package db

import (
	"fmt"
	"glamgrove/pkg/config"
	"glamgrove/pkg/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func to connect data base using config(database config) and return address of a new instnce of gorm DB
func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("Failed to connect with database")
		return nil, err
	}

	// migrate the database tables
	db.AutoMigrate(
		//user
		domain.User{},

		//admin
		domain.Admin{},

		//product
		domain.Product{},

		domain.ProductItem{},

		domain.Category{},
	)

	if err != nil {
		log.Fatal("DB Migration failed")
		return nil, nil
	}
	fmt.Println("DB migration success")
	return db, nil
}
