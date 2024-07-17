package controllers

import (
	"blog-platform/models"
	"blog-platform/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Basic input validation
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	found := false
	for _, u := range models.Users {
		if u.Username == user.Username {
			err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
			if err == nil {
				token, err := utils.GenerateJWT(u.Username, u.Role)
				if err != nil {
					http.Error(w, "Error generating JWT", http.StatusInternalServerError)
					fmt.Println("Error generating JWT:", err)
					return
				}
				utils.Respond(w, map[string]interface{}{"token": token})
				return
			}
			fmt.Println("Password does not match for user:", u.Username)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("User not found:", user.Username)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Basic input validation
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Additional validation if needed (e.g., password complexity, username uniqueness)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error while hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hash)

	resp := user.Create()
	utils.Respond(w, map[string]interface{}{"message": resp})
}
