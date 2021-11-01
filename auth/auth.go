package auth

import (
	"containerization/models"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

var JwtKey = []byte("secretKey")

var USers = map[string]string{
	"mayur": "password1",
	"user2": "password2",
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
		claims := &models.Claims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims,
			func(t *jwt.Token) (interface{}, error) {
				return JwtKey, nil
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
