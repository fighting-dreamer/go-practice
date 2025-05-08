package main

import (
	"github.com/gorilla/mux"
)

// Have the code for the server.
type Server struct {
	Port   int
	Router *mux.Router
}

func NewServer(port int) *Server {
	return &Server{
		Port:   port,
		Router: mux.NewRouter(),
	}
}

func (s *Server) SetupRoutes() {
	router := s.Router

	// mux.CORSMethodMiddleware(router)
	
	router.HandleFunc("/system", SystemHandler).Methods("GET")
	router.HandleFunc("/todo", .AddTodo).Methods("POST")
}
