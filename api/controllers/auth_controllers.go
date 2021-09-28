package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/VictorKabata/quotes-api/api/models"
	"github.com/VictorKabata/quotes-api/api/responses"
)

func (server *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	auth := models.Auth{}
	generatedLogin, err := auth.SignInUser(server.DB, user.Email, user.Username, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.SUCCESS(w, http.StatusOK, generatedLogin)
}
