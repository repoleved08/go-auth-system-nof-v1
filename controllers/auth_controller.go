package controllers

import (
	"encoding/json"
	"go-auth-system/models"
	"go-auth-system/utils"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = models.RegisterUser(&user)
	if err != nil {
		http.Error(w, "user registration failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusBadRequest)
		return
	}

	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "invalid input", http.StatusInternalServerError)
		return
	}

	user,err := models.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}
	tokenString, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "error generating jwt token", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"tokenString": tokenString})
	json.NewEncoder(w).Encode(user)
}
