package handler

import "github.com/gorilla/mux"

func Router(router *mux.Router) {
	h := &Handler{}

	router.HandleFunc("/users", h.getUsers).Methods("GET")
	router.HandleFunc("/users", h.createUser).Methods("POST")
	router.HandleFunc("/users/{id}", h.getUser).Methods("GET")
	router.HandleFunc("/users/{id}", h.updateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", h.deleteUser).Methods("DELETE")
}
