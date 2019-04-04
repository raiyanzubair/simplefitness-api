package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"simplefitnessApi/auth"
	"simplefitnessApi/model"
	"simplefitnessApi/usecase"
	"strconv"
	"time"
)

//Auth handler interface
type Auth struct {
	uc *usecase.User
}

// NewAuthHandler returns a new handler for auth requests
func NewAuthHandler(uc *usecase.User) *Auth {
	return &Auth{uc}
}

// GetUser returns details for a specific user
func (au *Auth) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, "Sorry something went wrong!", http.StatusInternalServerError)
		return
	}
	user, err := au.uc.GetByID(userID)
	if err != nil {
		http.Error(w, "Sorry something went wrong!", http.StatusInternalServerError)
		return
	}
	js, _ := json.Marshal(user)
	w.Write(js)
}

// HandleSignIn does user sign in
func (au *Auth) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	inputUser := model.User{}
	json.NewDecoder(r.Body).Decode(&inputUser)

	if inputUser.Email == "" || inputUser.Password == "" {
		http.Error(w, "Email and password required", http.StatusBadRequest)
		return
	}

	retrievedUser, err := au.uc.GetByEmail(inputUser.Email)
	if err != nil {
		http.Error(w, "Authorization error", http.StatusUnauthorized)
		return
	}
	log.Print(retrievedUser)
	err = bcrypt.CompareHashAndPassword([]byte(retrievedUser.HashedPassword), []byte(inputUser.Password))
	if err != nil {
		http.Error(w, "Authorization error", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(retrievedUser.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(5 * time.Minute),
	})
	json.NewEncoder(w).Encode(retrievedUser)
}

// HandleSignUp does user sign up/registration
func (au *Auth) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	var inputUser model.User
	json.NewDecoder(r.Body).Decode(&inputUser)
	if inputUser.Email == "" || inputUser.Password == "" {
		http.Error(w, "Email and password required", http.StatusBadRequest)
		return
	}

	retrievedUser, err := au.uc.GetByEmail(inputUser.Email)
	if retrievedUser != nil {
		http.Error(w, "Authorization error", http.StatusUnauthorized)
		return
	}

	newUser := &model.User{
		Email:     inputUser.Email,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(inputUser.Password), bcrypt.DefaultCost)
	newUser.HashedPassword = string(hashed)

	created, err := au.uc.Create(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := auth.GenerateJWT(created.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(5 * time.Minute),
	})
	js, _ := json.Marshal(created)
	w.Write(js)
}

func (au *Auth) handle(w http.ResponseWriter, r *http.Request) {

}
