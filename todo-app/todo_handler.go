package main

import "net/http"

// have the code for the todo api routes

type TodoHandler struct {
	todoService TodoService
}

func NewTodoHandler(todoService TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (th *TodoHandler) AddTodo(w http.ResponseWriter, r *http.Request) {
	// for each api call, we check if the body is empty or not ?
	// for each api call, we will need to get the request and construct a request object
	// for each api call, we check if this object is valid.
	// for each api call, we will have some authentication/authorization
	// for each api call, we will have some monitoring setup
	
}
