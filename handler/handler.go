package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    users,
		"message": "status OK",
	})
}

func (u *UserAPI) GetUser(c *gin.Context) {
	id := c.Param("id")
	if _, err := strconv.ParseInt(id, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidType.Error(),
		})
		return
	}

	user, err := u.userService.GetUser(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	//userData.data is a list
	var users []types.User
	users = append(users, user)
	c.JSON(http.StatusOK, gin.H{
		"data":    users,
		"message": "status OK",
	})
	return
}

func (u *UserAPI) UpdateUser(c *gin.Context) {
}

func (u *UserAPI) DeleteUser(c *gin.Context) {
}

func (fh *FHandler) TmplHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./view/layout.html"))

	var users types.UserData
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
	log.Println("Body Response: " + string(body))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, struct {
			Title   string
			Input   []string
			Operate []string
			Data    []types.User
			Message string
		}{
			Title:   "users",
			Input:   []string{"ID", "Name", "Age"},
			Operate: []string{"create", "readAll", "readOne", "update", "delete"},
			Data:    users.Data,
			Message: users.Message,
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
