package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

type Todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var Todos = make([]Todo, 0)

func main() {
	r := mux.NewRouter()

	todoRouter := r.PathPrefix("/todo-backend").Subrouter()

	todoRouter.HandleFunc("/todos", getTodos).Methods("GET")
	todoRouter.HandleFunc("/todos", postTodos).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	handler := c.Handler(r)

	// Start the server
	port := "80" // For testing
	fmt.Printf("Server started in port %s\n", port)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		return
	}
}

func postTodos(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	todo := Todo{}
	err := decoder.Decode(&todo)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
	}
	Todos = append(Todos, todo)
	responseWithJSON(w, http.StatusOK, todo)
}

func getTodos(w http.ResponseWriter, request *http.Request) {
	responseWithJSON(w, http.StatusOK, Todos)
}
