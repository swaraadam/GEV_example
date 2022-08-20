package database

import (
	"database/sql"
	"gev_example/helpers"
	sqlc "gev_example/models/sqlc"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Query *sqlc.Queries
}

var PG = &Postgres{
	Query: sqlc.New(GetPGConnection()),
}

func GetPGConnection() *sql.DB {
	connection, err := sql.Open(helpers.Env.PGDatabaseDriver, helpers.Env.PGDatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	return connection
}

func (pg *Postgres) InitializeDB() {
	log.Print("Initializing database...")
	// Creating gev_example DB
	db, err := sql.Open(helpers.Env.PGDatabaseDriver, helpers.Env.PGDatabaseInitURL)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE DATABASE gev_example;")
	if err, ok := err.(*pq.Error); ok {
		if err.Code.Name() == "duplicate_database" {
			log.Print("gev_example database already exists.")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Print("gev_example database created.")
	}

	// Migrating gev_example Schemas
	db, err = sql.Open(helpers.Env.PGDatabaseDriver, helpers.Env.PGDatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://database/migration", helpers.Env.DatabaseName, driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("schema already migrated.")
		}
	}
}
