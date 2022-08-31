package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Yefhem/hello-world-case/helpers"
	"github.com/Yefhem/hello-world-case/models"
	"github.com/Yefhem/hello-world-case/service"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// --------------------> Methods

// ----------> Create a New User Account
// ----------> Return 201, 400 or 500 status code
func (cont *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusNotImplemented, err.Error())
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusNotImplemented, err.Error())
		return
	}

	userModel, err := cont.userService.CreateUser(user)
	if err != nil {
		log.Println(err)
		helpers.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.Response(w, http.StatusCreated, userModel)
}
