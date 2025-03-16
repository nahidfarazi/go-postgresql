package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nahidfarazi/go-postgresql/handler"
)

func main() {
	handler.InitDB()
	r := chi.NewRouter()
	r.Get("/", handler.GetAllUser)
	r.Get("/users", handler.GetAllUser)
	r.Get("/user/{id}", handler.GetUserByID)
	r.Post("/users", handler.CreateUser)
	r.Put("/user/{id}", handler.UpdateUser)
	r.Delete("/user/{id}", handler.DeleteUser)

	log.Println(http.ListenAndServe(":4040", r))
}
