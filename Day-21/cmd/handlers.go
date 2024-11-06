package main

import (
	"encoding/json"
	"net/http"
)

var users = []User{
	{ID: 1, Username: "testuser", Password: "password1234"},
}

// adds new user tro the list
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

// verifies the credentials and returns a JWT if valid
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials User
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request paylaod", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		// verify the login username & password with the DB username & password
		if user.Username == credentials.Username && user.Password == credentials.Password {
			token, err := GenerateJWT(user.ID) // create a JWT token for the valid user
			if err != nil {
				http.Error(w, "Error generating token", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"token": token})
			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

// create the protected endpoint
// thsi endpoint allows users to access their profile only if authenticated
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to your profile!"))
}
