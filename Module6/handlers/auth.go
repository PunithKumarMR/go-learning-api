package handlers

import (
	"Module6/database"
	"Module6/models"
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hashedPassword, err1 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	user.Password = string(hashedPassword)
	result := database.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var dbUser models.User
	result := database.DB.Where("email=?", user.Email).First(&dbUser)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err1 := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err1 != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": dbUser.ID,
		"email":   dbUser.Email,
	})
	tokenString, err2 := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
