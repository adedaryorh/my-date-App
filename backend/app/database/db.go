package database

import (
	"strconv"

	"backend.app/configs"
	"backend.app/database/postgres"
	"backend.app/internal/services/redisservice"
	"backend.app/pkg/logger"
)

type DB struct {
	Postgres *postgres.Postgres
	Redis    *redisservice.Redis
}

func ConnectDB(config *configs.Config, log *logger.Logger) DB {
	redis := redisservice.NewConnection(config)
	poolMax, _ := strconv.Atoi(config.PgPoolMax)
	pg, err := postgres.New(config, postgres.MaxPoolSize(poolMax))
	if err != nil {
		log.Fatal("postgres connection failed : %v", err)
	}
	return DB{
		Postgres: pg,
		Redis:    &redis,
	}
}
