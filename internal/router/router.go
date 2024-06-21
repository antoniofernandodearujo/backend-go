package router

import (
	"challenge-api/internal/controllers"
	"net/http"

	_ "challenge-api/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Sensedia Challenge API
// @version 1
// @description This is the API for the Sensedia Challenge
// @BasePath /api/v1
func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Root"))
	})

	// User routes
	router.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", controllers.GetAllUsers)
		r.Post("/create", controllers.CreateUser)
		r.Get("/{username}", controllers.GetUserByUsername) // Explicit route for username
		r.Route("/{id:[a-fA-F0-9\\-]+}", func(r chi.Router) { // UUID regex
			r.Get("/albums", controllers.GetAlbumsByUserID)
			r.Get("/posts", controllers.GetPostsByUserID)
			r.Delete("/albums/{album_id}", controllers.RemoveAlbumFromUser)
			r.Get("/", controllers.GetUserByID)
			r.Put("/", controllers.UpdateUser)
			r.Delete("/", controllers.DeleteUser)
		})
	})

	// Album routes
	router.Route("/api/v1/albums", func(r chi.Router) {
		r.Get("/", controllers.GetAllAlbums)
		r.Post("/create", controllers.CreateAlbum)
		r.Post("/save", controllers.AddAlbumToUser)
		r.Route("/{id:[a-fA-F0-9\\-]+}", func(r chi.Router) { // UUID regex
			r.Get("/", controllers.GetAlbumByID)
			r.Put("/", controllers.UpdateAlbum)
			r.Delete("/", controllers.DeleteAlbum)
		})
	})

	// Post routes
	router.Route("/api/v1/posts", func(r chi.Router) {
		r.Get("/", controllers.GetAllPosts)
		r.Post("/create", controllers.CreatePost)
		r.Route("/{id:[a-fA-F0-9\\-]+}", func(r chi.Router) { // UUID regex
			r.Get("/", controllers.GetPostByID)
			r.Put("/", controllers.UpdatePost)
			r.Delete("/", controllers.DeletePost)
		})
	})

	// Swagger route
	router.Route("/swagger", func(r chi.Router) {
		r.Get("/*", httpSwagger.WrapHandler)
	})

	return router
}
