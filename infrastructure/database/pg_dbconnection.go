package infrastructure_database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PgConnection() *gorm.DB {
	dburi := "postgres://okaprinarjaya:ok4prin4rj4y4@localhost:5432/ats_simple?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dburi), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
