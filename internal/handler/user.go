package handler

import (
	"log"
    "net/http"
	"Go-Boilerplate/internal/helpers"
)

// LoginCheckHandler responds with "user logged in successfully" when a POST request is made to /login.
func LoginCheckHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
		helpers.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
        return
    }
    // Proceed with login logic for POST method...
	log.Printf("Login check successful")
}
