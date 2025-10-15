package handler

import (
	"encoding/json"
	"fmt"
	"minha-primeira-api/internal/models"
	"minha-primeira-api/pkg/auth"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "erro ao ler json")
		return
	}

	err = models.AuthenticateUser(&u)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, err.Error())
		return
	}

	tokenString, err := auth.CreateToken(u.Name)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token":tokenString})
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := auth.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "invalid token")
		return
	}
	fmt.Fprint(w, "Welcome to the protected area")
}
