package db

import (
	"fmt"
	"glamgrove/pkg/config"
	"glamgrove/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func to connect data base using config(database config) and return address of a new instnce of gorm DB
func ConnectDatbase(cfg config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	// migrate the database tables
	db.AutoMigrate(
		//user
		domain.Users{},

		//admin
		domain.Admin{},

		//product
		domain.ProductCategory{},
		domain.Product{},
		domain.Variation{},
		domain.VariationOption{},
		domain.ProductItem{},
		domain.ProductConfiguraion{},
		domain.ProductImage{},
	)

	return db, err
}
