package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/OAOv/restful_CRUD/repo"
	"github.com/OAOv/restful_CRUD/types"
	"github.com/gorilla/mux"
)

type Handler struct {
}

type FHandler struct {
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := repo.GetUsers()
	if err != nil {
		log.Println(err)
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = repo.CreateUser(body)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintf(w, "New User was created.")
	}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user, err := repo.GetUser(params)
	if err != nil {
		log.Println(err)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	err = repo.UpdateUser(params, body)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintf(w, "User with ID = %s was updated", params["id"])
	}
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := repo.DeleteUser(params)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintf(w, "User with ID = %s was deleted", params["id"])
	}
}

func (fh *FHandler) TmplHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./view/layout.html"))

	req, _ := http.NewRequest("GET", "http://localhost:8000/users", nil)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var users []types.User
	json.Unmarshal(body, &users)
	fmt.Println(users)

	if r.Method != http.MethodPost {
		tmpl.Execute(w, struct {
			Title   string
			Input   []string
			Operate []string
			List    []types.User
		}{
			Title:   "users",
			Input:   []string{"ID", "Name", "Age"},
			Operate: []string{"create", "update", "delete"},
			List:    users,
		})
	} else {
		log.Println("button: " + r.FormValue("btn"))

		switch r.FormValue("btn") {
		case "create":
			body, err = DoCreateRequest(r.FormValue("Name"), r.FormValue("Age"))

		case "update":
			body, err = DoUpdateRequest(r.FormValue("ID"), r.FormValue("Name"), r.FormValue("Age"))

		case "delete":
			body, err = DoDeleteRequset(r.FormValue("ID"))
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Body response: " + string(body))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
