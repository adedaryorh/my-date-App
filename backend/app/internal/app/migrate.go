package app

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"

	"backend.app/configs"
	"backend.app/pkg/logger"
)

func IsLocal() bool {
	return os.Getenv("APP_ENV") == "" || os.Getenv("APP_ENV") == "dev"
}

func runMigrations(env *configs.Config, log *logger.Logger) {
	dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		env.PGUser,
		env.PGPassword,
		env.PGHost,
		env.PGPort,
		env.PGDatabase)

	if IsLocal() {
		dbConnectionString += "?sslmode=disable"
	}
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal("sql.Open failed: %v", err)
	}
	defer db.Close()
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal("goose.SetDialect failed: %v", err)
	}
	// Assuming migrations are located in "db/migrations"
	if err := goose.Up(db, "database/goose/migrations"); err != nil {
		log.Fatal("goose.Up failed: %v", err)
	}
}
