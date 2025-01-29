package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/rs/zerolog/log"
)

const dir = "migrations"

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	_ = flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 2 {
		flags.Usage()
		return
	}

	command := args[0]

	// get the environment
	if os.Getenv("APP_ENV") != "prod" && os.Getenv("APP_ENV") != "stg" && os.Getenv("APP_ENV") != "beta" {
		if err := godotenv.Load("../../.env"); err != nil {
			log.Fatal().Err(err).Msg(err.Error())
		}
	}
	sslMode := "disable"
	if os.Getenv("APP_ENV") != "dev" {
		sslMode = "require"
	}
	db, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("PG_USER"),
			os.Getenv("PG_PASSWORD"),
			os.Getenv("PG_HOST"),
			os.Getenv("PG_PORT"),
			os.Getenv("PG_DATABASE"),
			sslMode,
		))

	if err != nil {
		log.Fatal().Err(err).Msgf("goose %v: %v", command, err)
	}

	var arguments []string
	for _, val := range args[1:] {
		if len(val) > 0 {
			arguments = append(arguments, val)
		}
	}

	log.Info().Msgf("running goose %s %v : args=%d", command, arguments, len(arguments))
	if err := goose.Run(command, db, dir, arguments...); err != nil {
		log.Fatal().Err(err).Msgf("goose %v: %v", command, err)
	}
}
