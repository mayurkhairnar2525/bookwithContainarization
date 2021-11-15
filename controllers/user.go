package controllers

import (
	"containerization/auth"
	"containerization/models"
	vipers "containerization/viper"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

var JwtKey = vipers.GetJwtKey()

func newUser() *models.Person {
	return &models.Person{}
}

func (c Controllers) Login(w http.ResponseWriter, request *http.Request) {
	var credentials models.Credentials
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
	claims := &models.Claims{
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

func (c Controllers) Registration(w http.ResponseWriter, req *http.Request) {
	log.Println("new user registering to the system")

	w.Header().Set("Content-Type", "application/json")
	if req.Body == nil {
		fmt.Fprintln(w, "nil body passed")
		return
	}
	person := newUser()                              //initialize the person
	err := json.NewDecoder(req.Body).Decode(&person) //Decode person from json
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprint(w, "new user registered successfully")
	log.Println("new user registered successfully")
}
