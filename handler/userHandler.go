package handler

import (
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/OAOv/restful_CRUD/service"
	"github.com/OAOv/restful_CRUD/types"
	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	userService service.Service
}

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
		"message": "create completed",
	})
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
		"message": "readAll completed",
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
		"message": "readOne completed",
	})
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

	key := reflect.TypeOf(user)
	value := reflect.ValueOf(user)
	var data = make(map[string]interface{})
	for i := 0; i < key.NumField(); i++ {
		tmpKey := key.Field(i).Tag.Get("json")
		tmpValue := value.Field(i).Interface()
		if tmpKey == "id" || tmpValue == "" {
			continue
		}
		data[tmpKey] = tmpValue
	}

	err = u.userService.UpdateUser(user.ID, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": "update completed",
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
		"message": "delete completed",
	})
}
