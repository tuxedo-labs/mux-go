package main

import (
	"fmt"
	"go-api/internal/handlers"
	"go-api/pkg/config"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.DBConnect()
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/{id}", handlers.GetProductById).Methods("GET")
	r.HandleFunc("/", handlers.CreateProduct).Methods("POST")
  r.HandleFunc("/{id}", handlers.DeleteProduct).Methods("DELETE")
  r.HandleFunc("/{id}", handlers.UpdateProduct).Methods("PATCH")
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
