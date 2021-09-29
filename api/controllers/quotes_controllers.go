package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/VictorKabata/quotes-api/api/auth"
	"github.com/VictorKabata/quotes-api/api/models"
	"github.com/VictorKabata/quotes-api/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) NewQuote(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	quote := models.Quote{}
	err = json.Unmarshal(body, &quote)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tokenId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenId != quote.UserID {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Unauthorized"))
		return
	}

	quote.Prepare()
	err = quote.ValidateInput("create")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	createdQuote, err := quote.SaveQuote(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.SUCCESS(w, http.StatusOK, createdQuote)
}

func (server *Server) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	quote := models.Quote{}
	allQuotes, err := quote.GetAllQuotes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.SUCCESS(w, http.StatusOK, allQuotes)
}

func (server *Server) GetQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	qid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	quote := models.Quote{}
	specificQuote, err := quote.GetQuote(server.DB, uint32(qid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.SUCCESS(w, http.StatusOK, specificQuote)
}

func (server *Server) UpdateQuotes(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	quote := models.Quote{}
	err = json.Unmarshal(body, &quote)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tokenId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unathorized"))
		return
	}
	if tokenId != quote.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	updatedQuote, err := quote.UpdateQuote(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.SUCCESS(w, http.StatusOK, updatedQuote)
}

func (server *Server) DeleteQuotes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["user_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	qid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if tokenId != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	quote := models.Quote{}
	_, err = quote.DeleteQuote(server.DB, uint32(qid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.SUCCESS(w, http.StatusOK, "Quote deleted")
}
