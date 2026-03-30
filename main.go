package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"fitness-api/db"
	"fitness-api/handlers"
)

func main() {
	// konekcija na bazu
	db.InitDB()

	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	// ROUTES
	r.Route("/api", func(r chi.Router) {

		// svi exercises (bez filtera)
		r.Get("/exercises", handlers.GetAllExercises)

		// filteri (bodyPart, equipment)
		r.Get("/exercises/filter", handlers.GetExercises)

		// get by id
		r.Get("/exercises/{id}", handlers.GetExerciseByID)

		// get instructions by exrcise id
		r.Get("/exercises/{id}/instructions", handlers.GetInstructionsByExerciseID)

		// get related exercises by id
		r.Get("/exercises/{id}/related", handlers.GetRelatedExercises)
	})

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
