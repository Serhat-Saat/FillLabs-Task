package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"user-management-api/errors"
)

// User struct's
type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
	UserPhone string `json:"userPhone"`
}

// GetAllUsersHandler to list all users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Sadece GET metodu destekleniyor", http.StatusMethodNotAllowed)
		return
	}

	users, err := GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByIDHandler to fetch the user with a specific ID
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Sadece GET metodu destekleniyor", http.StatusMethodNotAllowed)
		return
	}

	// Get ID parameter from URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID parametresi gerekli", http.StatusBadRequest)
		return
	}

	// Convert ID from string to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Geçersiz ID formatı", http.StatusBadRequest)
		return
	}

	// Bring the user
	user, err := GetUserByID(id)
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUserHandler for create the new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Sadece POST metodu destekleniyor", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// UpdateUserHandler for update the user
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Sadece PUT metodu destekleniyor", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := UpdateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUserHandler for delete the user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Sadece DELETE metodu destekleniyor", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	fmt.Println(r.URL.Query())
	fmt.Println(idStr)
	if idStr == "" {
		http.Error(w, "ID parametresi gerekli", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Geçersiz ID formatı", http.StatusBadRequest)
		return
	}

	if err := DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
