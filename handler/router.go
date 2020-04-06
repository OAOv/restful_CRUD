package handler

import "github.com/gorilla/mux"

func Router(router *mux.Router) {
	h := &Handler{}

	router.HandleFunc("/users", h.GetUsers).Methods("GET")
	router.HandleFunc("/users", h.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", h.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
}

func FRouter(router *mux.Router) {
	fh := &FHandler{}

	router.HandleFunc("/", fh.TmplHandler)
}
