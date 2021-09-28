package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/VictorKabata/quotes-api/api/auth"
	"github.com/VictorKabata/quotes-api/api/models"
	"github.com/VictorKabata/quotes-api/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user.Prepare()
	err = user.Validate("create")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	createdUser, err := user.SaveUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	responses.SUCCESS(w, http.StatusCreated, createdUser)
}

func (server *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	users, err := user.GetAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	responses.SUCCESS(w, http.StatusOK, users)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]

	user := models.User{}

	specificUser, err := user.GetUser(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	responses.SUCCESS(w, http.StatusOK, specificUser)
}

func (serv *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	tokenId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
	} else if tokenId != uid {
		responses.ERROR(w, http.StatusUnauthorized, err)
	}

	updatedUser, err := user.UpdateUser(serv.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	responses.SUCCESS(w, http.StatusOK, updatedUser)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]

	tokenId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
	} else if tokenId != uid {
		responses.ERROR(w, http.StatusUnauthorized, err)
	}

	user := models.User{}
	_, err = user.DeleteUser(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	responses.SUCCESS(w, http.StatusOK, "Account deleted")
}

func (server *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	userToken, err := user.SignInUser(server.DB, user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	responses.SUCCESS(w, http.StatusUnprocessableEntity, userToken)
}
