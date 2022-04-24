package main

import (
	"celebut-api/internal/routes"
	"net/http"
)

func init() {

}

func main() {
	routes.NewAppRouter()

	srv := http.Server{
		Addr:    ":8080",
		Handler: routes.NewAppRouter(),
	}

	err := srv.ListenAndServe()
	if err != nil {

	}
}
