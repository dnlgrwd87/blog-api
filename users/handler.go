package users

import (
	"encoding/json"
	"github.com/dnlgrwd87/blog-api/helpers"
	"github.com/dnlgrwd87/blog-api/models"
	"github.com/go-chi/chi"
	"net/http"
)

type UserHandler struct {
	Service UserService
}

func (handler *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := handler.Service.GetAllUsers()
	helpers.RespondwithJSON(w, http.StatusOK, users)
}

func (handler *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := handler.Service.GetUserById(id)
	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	helpers.RespondwithJSON(w, http.StatusOK, user)
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	if err := user.ValidateIncomingPayload(); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	createdUser, _ := handler.Service.CreateUser(user)

	helpers.RespondwithJSON(w, http.StatusOK, createdUser)
}
