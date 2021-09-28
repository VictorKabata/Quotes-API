package controllers

import (
	"net/http"

	"github.com/VictorKabata/quotes-api/api/responses"
)

func (server *Server) HomePage(w http.ResponseWriter, r *http.Request) {
	responses.SUCCESS(w, http.StatusOK, "Welcome to quotes endpoint")
}
