package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewAppRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Celebut API"))
	})

	return r
}
