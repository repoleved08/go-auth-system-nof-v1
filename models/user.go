package models

import (
	"database/sql"
	"errors"
	"go-auth-system/config"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	err = config.DB.QueryRow(query, user.Username, user.Email, string(hashedPassword)).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func AuthenticateUser(username, password string) (User, error) {
	var user User
	query := "SELECT username, password FROM users WHERE username=$1"
	err := config.DB.QueryRow(query, username).Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return user, errors.New("invalid username or password")
	} else if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid username or password")
	}
	return user, nil
}
