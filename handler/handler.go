package handler

import "net/http"

type Handler struct {
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	repo.GetUsers(r.Context())
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
