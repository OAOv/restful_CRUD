package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/OAOv/restful_CRUD/service"
	"github.com/OAOv/restful_CRUD/types"
	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	userService service.UserService
}

type FHandler struct {
}

var isOne = false
var searchID = ""

func (u *UserAPI) CreateUser(c *gin.Context) {
}

func (u *UserAPI) GetUsers(c *gin.Context) {
	users, err := u.userService.GetUsers()
	if err != nil {
		log.Println(types.ErrInvalidParms)
	}
}

func (u *UserAPI) GetUser(c *gin.Context) {
}

func (u *UserAPI) UpdateUser(c *gin.Context) {
}

func (u *UserAPI) DeleteUser(c *gin.Context) {
}

func (fh *FHandler) TmplHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./view/layout.html"))

	var users []types.User
	var body []byte
	var err error
	if !isOne {
		body, err = DoReadAllRequest()
		json.Unmarshal(body, &users)
	} else {
		isOne = false
		body, err = DoReadOneRequest(searchID)
		json.Unmarshal(body, &users)
	}

	if r.Method != http.MethodPost {
		tmpl.Execute(w, struct {
			Title   string
			Input   []string
			Operate []string
			List    []types.User
		}{
			Title:   "users",
			Input:   []string{"ID", "Name", "Age"},
			Operate: []string{"create", "readAll", "readOne", "update", "delete"},
			List:    users,
		})
	} else {
		log.Println("button: " + r.FormValue("btn"))

		switch r.FormValue("btn") {
		case "create":
			body, err = DoCreateRequest(r.FormValue("Name"), r.FormValue("Age"))

		case "readOne":
			isOne = true
			searchID = r.FormValue("ID")

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
