package handler

import (
	"encoding/json"
	"minha-primeira-api/internal/models"
	"net/http"
)

// GET METHOD
func GetUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
}

// DELETE METHOD
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var deleteUser models.User
	if err := json.NewDecoder(r.Body).Decode(&deleteUser); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	found := false
	for i, v := range models.Users {
		if v.Name == deleteUser.Name {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "user deleted succesfully"}
	json.NewEncoder(w).Encode(response)
}

// PATCH/UPDATE METHOD
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updateUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		http.Error(w, "invalid request method", http.StatusBadRequest)
		return
	}

	found := false
	for i, v := range models.Users {
		if v.Name == updateUser.Name {
			models.Users[i].Age = updateUser.Age
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "user not found", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
}

// POST METHOD
func CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	models.Users = append(models.Users, newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
