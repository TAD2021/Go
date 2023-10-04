package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var userDB = map[string]string{ 
	"exampleUser": "examplePassword", // This is just an example, storing plaintext passwords is a bad idea.
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expectedPassword, ok := userDB[user.Username]

	if !ok || expectedPassword != user.Password {
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