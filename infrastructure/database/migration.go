package infrastructure_database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migration(dbMigrationFilesPath string) (*migrate.Migrate, error) {
	dburi := "postgres://okaprinarjaya:ok4prin4rj4y4@localhost:5432/ats_simple_testing?sslmode=disable"
	m, err := migrate.New("file://"+dbMigrationFilesPath, dburi)

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return m, nil
}
