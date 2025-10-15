package handler

import (
	"encoding/json"
	"minha-primeira-api/internal/models"
	"net/http"
	"strconv"
)

// GET METHOD
func GetUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
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

	if err := models.InsertUsers(&newUser); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// DELETE METHOD
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	if err := models.DeleteUsersById(id); err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
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

	if err := models.UpdateUser(&updateUser); err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&updateUser)
}
