package auth

import (
	"containerization/config"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

var tokenStore = make(map[string]string)

type Credentials struct {
	Username string `json:"username" validate:"required,min=4,max=15"`
	Password string `json:"password" validate:"required,min=4"`
}

type Claims struct {
	Username string `json:"username" validate:"required"`
	jwt.StandardClaims
}

var jwtKey = config.GetJwtKey()

var USers = map[string]string{
	"mayur": "password1",
}

func IsAuthorised(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("Authenticating user")
		cookie, err := req.Cookie("token")
		if err != nil {
			log.Fatal("err", err)
			return
		}
		tokenStr := cookie.Value
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims,
			func(t *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
		if err != nil {
			log.Fatal("err", err)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		endpoint(w, req)
	}
}
