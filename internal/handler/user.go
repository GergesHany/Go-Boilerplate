package handler

import (
	"Go-Boilerplate/internal/helpers"
	"log"
	"net/http"
)

// LoginCheckHandler responds with "user logged in successfully" when a POST request is made to /login.
func LoginCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	log.Printf("Login check successful")
}
