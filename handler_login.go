package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var userDB = map[string]string{ 
	"exampleUser": hashPassword("examplePassword"), // This is just an example
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expectedPassword, ok := userDB[user.Username]

	if !ok || !checkPasswordHash(user.Password, expectedPassword) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid username or password\n")
		return
	}

	// Upon successful login, you usually generate and return a JWT (JSON Web Token) here 
	// (requires a third-party library and some setup).
	// For the sake of simplicity, we just return a success message.

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Login successful\n")
}