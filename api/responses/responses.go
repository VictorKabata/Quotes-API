package responses

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

//Formats data to JSON format when request is successful
func SUCCESS(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

//Formats data into custom error struct response if error occours
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		SUCCESS(w, statusCode, Error{err.Error()})
	}
	SUCCESS(w, http.StatusBadRequest, nil)
}
