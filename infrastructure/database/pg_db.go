package infrastructure_database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBNAME         = "ats_simple"
	DBNAME_TESTING = "ats_simple_testing"
)

func PgConnection() *gorm.DB {
	dburi := fmt.Sprintf("postgres://okaprinarjaya:ok4prin4rj4y4@localhost:5432/%s?sslmode=disable", DBNAME)
	db, err := gorm.Open(postgres.Open(dburi), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func PgConnectionTesting(purpose string) *gorm.DB {
	var dburi string

	if purpose == "setup" {
		dburi = fmt.Sprintf("postgres://okaprinarjaya:ok4prin4rj4y4@localhost:5432/%s?sslmode=disable", "postgres")
	} else {
		dburi = fmt.Sprintf("postgres://okaprinarjaya:ok4prin4rj4y4@localhost:5432/%s?sslmode=disable", DBNAME_TESTING)
	}

	db, err := gorm.Open(postgres.Open(dburi), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func ManageDatabase() error {
	dbconn := PgConnectionTesting("setup")
	stmt := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname = '%s';", DBNAME_TESTING)
	rs := dbconn.Raw(stmt)

	if rs.Error != nil {
		log.Fatalln(rs.Error.Error())
		return rs.Error
	}

	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		log.Println("Creating database: " + DBNAME_TESTING)

		stmt := fmt.Sprintf("CREATE DATABASE %s OWNER %s;", DBNAME_TESTING, "okaprinarjaya")
		if rs := dbconn.Exec(stmt); rs.Error != nil {
			log.Fatalln(rs.Error.Error())
			return rs.Error
		}

	} else {
		// Drop first
		log.Println("Dropping database: " + DBNAME_TESTING)

		stmtDrop := fmt.Sprintf("DROP DATABASE %s;", DBNAME_TESTING)
		if rs := dbconn.Exec(stmtDrop); rs.Error != nil {
			log.Fatalln(rs.Error.Error())
		}

		// Then create new
		log.Println("Creating database: " + DBNAME_TESTING)

		stmtCreate := fmt.Sprintf("CREATE DATABASE %s OWNER %s;", DBNAME_TESTING, "okaprinarjaya")
		if rs := dbconn.Exec(stmtCreate); rs.Error != nil {
			log.Fatalln(rs.Error.Error())
			return rs.Error
		}
	}

	sql, err := dbconn.DB()
	defer func() {
		_ = sql.Close()
	}()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return nil
}
