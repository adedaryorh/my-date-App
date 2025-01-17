package app

import (
	"backend.app/internal/repo"
	"backend.app/internal/repo/postgres/accounts"
	"backend.app/internal/repo/postgres/posts"
	"backend.app/pkg/postgres"
)

var (
	industryRepo    repo.Industry
	clientRepo      repo.Client
	userRepo        repo.User
	registerOTPRepo repo.UserOTP
	postsRepo       repo.Post
)

func initialiseRepositories(pg *postgres.Postgres) {
	industryRepo = accounts.NewIndustryRepo(pg)
	clientRepo = accounts.NewClientRepo(pg)
	userRepo = accounts.NewUserRepo(pg)
	registerOTPRepo = accounts.NewUserOTPRepo(pg)
	postsRepo = posts.NewPostsRepo(pg)
}
