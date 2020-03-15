package user

import (
	"encoding/json"
	"github.com/dnlgrwd87/blog-api/helpers"
	"github.com/dnlgrwd87/blog-api/models"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service UserServiceInterface
}

type UserHandlerInterface interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func (handler *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.Service.GetAllUsers()

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondwithJSON(w, http.StatusOK, users)
}

func (handler *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := handler.Service.GetUserById(userID)

	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	helpers.RespondwithJSON(w, http.StatusOK, user)
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if err := user.ValidateUserJson(); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	createdUser, _ := handler.Service.CreateUser(user)
	helpers.RespondwithJSON(w, http.StatusOK, createdUser)
}
