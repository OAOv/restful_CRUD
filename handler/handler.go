package handler

import "net/http"

type Handler struct {
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {

}
