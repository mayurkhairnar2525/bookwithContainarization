package controllers

import (
	"containerization/auth"
	"containerization/config"
	"containerization/models"
	users "containerization/repository/user"
	"containerization/utils"
	"containerization/validation"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Controllers struct {
	Repository users.UserRepository
	Db         *sql.DB
}

var userss []models.Person

var JwtKey = config.Config{}

//var JwtKey = config.Config{}

func (u Controllers) Login(w http.ResponseWriter, request *http.Request) {
	var credentials auth.Credentials
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	v := validator.New()
	if err = v.Struct(credentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPassword, ok := auth.USers[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//expirationTime := time.Now().Add(viper.GetCookieExpiryTime())
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &auth.Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	log.Println("user successfully logged in")
}

func (u Controllers) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.Person
		var userUsername string
		var error models.Error
		err := json.NewDecoder(r.Body).Decode(&user)
		validationError := validation.ValidateCreateUser(&user) //validate inputs of user and display errors if any
		if validationError != nil {
			validation.DisplayError(w, validationError)
			return
		}
		userRepo := users.UserRepository{}
		userUsername, err = userRepo.CreateUser(user)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, userUsername)
	}
}

func (u Controllers) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.Person
		userss = []models.Person{}
		userRepo := users.UserRepository{}
		users, err := userRepo.GetUsers(user, userss)

		v := validator.New()
		if err = v.Struct(user); err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			log.Println("err", err)
			return
		}
	}
}

func (u Controllers) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.Person
		params := mux.Vars(r)
		userRepo := users.UserRepository{}
		username, _ := params["username"]
		rowsDeleted, err := userRepo.DeleteUser(username)
		v := validator.New()

		if err = v.Struct(user); err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
		fmt.Fprint(w, "new user Deleted successfully")
	}
}
