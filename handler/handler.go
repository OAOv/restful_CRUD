package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/OAOv/restful_CRUD/repo"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := repo.GetUsers()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = repo.CreateUser(body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "New User was created.")
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user, err := repo.GetUser(params)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	err = repo.UpdateUser(params, body)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "User with ID = %s was updated", params["id"])
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := repo.DeleteUser(params)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "User with ID = %s was deleted", params["id"])
}
