package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

//Formats data to JSON format when request is successful
func SUCCESS(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

//Formats data into custom error struct response if error occours
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		SUCCESS(w, statusCode, Error{err.Error()})
		return
	}
	SUCCESS(w, http.StatusBadRequest, nil)
}
