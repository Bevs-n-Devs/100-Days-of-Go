package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

// define the GoUser model
type GoUser struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

/*
Implement each handler, starting with creating a user and then continuing with read, update, and delete.
*/

func createUser(w http.ResponseWriter, r *http.Request) {
	var user GoUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request paylaod", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO gousers (name, email) VALUES ($1, $2) RETURNING id"
	err = db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Error create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SEELCT id, name, email FROM gousers")
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []GoUser
	for rows.Next() {
		var user GoUser
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, "Error reading user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user GoUser
	query := "SELECT id, name, email FROM gousers WHERE id = $1"
	err = db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user GoUser
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "UPDATE gousers SET name = $1, email = $2 WHERE id = $3"
	_, err = db.Exec(query, user.Name, user.Email, id)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	user.ID = id
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM gousers WHERE id = $1"
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
