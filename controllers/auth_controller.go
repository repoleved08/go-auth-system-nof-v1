package controllers

import (
	"encoding/json"
	"go-auth-system/config"
	"go-auth-system/models"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
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
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2. $3)"
	err = config.DB.Query (query, user.Username, user.Email, user.Password)
}

func LoginUser() {

}
