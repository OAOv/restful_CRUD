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
var isGet = true
var searchID = ""
var body []byte
var err error

func (u *UserAPI) CreateUser(c *gin.Context) {
	user := types.User{}
	c.BindJSON(&user)
	if user.Name == "" || user.Age == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrEmptyInput.Error(),
		})
		return
	}

	err = u.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":    nil,
		"message": "create compeleted",
	})
	return
}

func (u *UserAPI) GetUsers(c *gin.Context) {
	users, err := u.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    users,
		"message": "readAll compeleted",
	})
}

func (u *UserAPI) GetUser(c *gin.Context) {
	id := c.Param("id")
	if val, err := strconv.ParseInt(id, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidType.Error(),
		})
		return
	} else if val <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidInputRange.Error(),
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
		"message": "readOne compeleted",
	})
	return
}

func (u *UserAPI) UpdateUser(c *gin.Context) {
	user := types.User{}
	c.BindJSON(&user)
	if val, err := strconv.ParseInt(user.ID, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidType.Error(),
		})
		return
	} else if val <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidInputRange.Error(),
		})
		return
	} else if user.Name == "" && user.Age == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrEmptyInput.Error(),
		})
		return
	}

	err = u.userService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": "update compeleted",
	})
}

func (u *UserAPI) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if val, err := strconv.ParseInt(id, 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidType.Error(),
		})
		return
	} else if val <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": types.ErrInvalidInputRange.Error(),
		})
		return
	}

	err = u.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": "delete compeleted",
	})
}

func (fh *FHandler) TmplHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./view/layout.html"))

	var users types.UserData

	if r.Method != http.MethodPost {
		if !isOne && isGet {
			body, err = DoReadAllRequest()
		} else if isGet {
			isOne = false
			body, err = DoReadOneRequest(searchID)
		}

		log.Println("Body Response: " + string(body))
		json.Unmarshal(body, &users)

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
			isGet = false
			body, err = DoCreateRequest(r.FormValue("ID"), r.FormValue("Name"), r.FormValue("Age"))

		case "readAll":
			isGet = true

		case "readOne":
			isOne = true
			isGet = true
			searchID = r.FormValue("ID")

		case "update":
			isGet = false
			body, err = DoUpdateRequest(r.FormValue("ID"), r.FormValue("Name"), r.FormValue("Age"))

		case "delete":
			isGet = false
			body, err = DoDeleteRequset(r.FormValue("ID"))
		}
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
