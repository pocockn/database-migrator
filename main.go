package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"database-migrator/config"
	"github.com/pocockn/models/api/football"
	"gopkg.in/gormigrate.v1"
)

func main() {
	fmt.Print("Starting migrations...\n")

	config := config.NewConfig()
	DB, err := gorm.Open("mysql", config.Database.URL)
	if err != nil {
		log.Fatal(err)
	}

	// Implement some sort of backing off process here.
	if err := DB.DB().Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Print("Database connection established")
	}
	DB.LogMode(true)

	// TODO: Only run migrations on dev and if we have new migrations added.
	gormMigrator := gormigrate.New(DB, gormigrate.DefaultOptions, GenerateMigrations())

	gormMigrator.InitSchema(func(tx *gorm.DB) error {
		log.Print("Creating initial table schema...")
		err := tx.AutoMigrate(
			&football.Team{},
			&football.Player{},
		)

		if err != nil {
			log.Printf("Err : %+v /n creating initial schema", err)
		}

		return nil
	})

	if err := gormMigrator.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Initial schema migration successfull")

	err = processSeeds(DB)
	if err != nil {
		log.Fatalf("Error will seeding database: %v", err)
	}
	log.Printf("Database seeding successfull")
}
