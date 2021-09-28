package middlewares

import (
	"net/http"

	"github.com/VictorKabata/quotes-api/api/auth"
	"github.com/VictorKabata/quotes-api/api/responses"
)

//Sets the response of every request to JSON format
func JsonMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

//Checks validity of token before every request
func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.IsTokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
		}
		next(w, r)
	}
}
