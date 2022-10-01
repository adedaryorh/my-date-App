package app

import (
	"celebut-api/internal/repo"
	"celebut-api/internal/repo/postgres/accounts"
	"celebut-api/internal/repo/postgres/celebrations"
	"celebut-api/pkg/postgres"
)

var (
	industryRepo     repo.Industry
	clientRepo       repo.Client
	userRepo         repo.User
	registerOTPRepo  repo.UserOTP
	celebrationsRepo repo.Celebration
)

func initialiseRepositories(pg *postgres.Postgres) {
	industryRepo = accounts.NewIndustryRepo(pg)
	clientRepo = accounts.NewClientRepo(pg)
	userRepo = accounts.NewUserRepo(pg)
	registerOTPRepo = accounts.NewRegisterOTPRepo(pg)
	celebrationsRepo = celebrations.NewCelebrationRepo(pg)
}
