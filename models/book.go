package models

import "github.com/dgrijalva/jwt-go"

type BookManagement struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Author       string `json:"author"`
	Prices       int    `json:"prices"`
	Available    string `json:"available"`
	PageQuality  string `json:"pagequality"`
	LaunchedYear string `json:"launchedyear"`
	Isbn         string `json:"isbn"`
	Stock        int    `json:"stock"`
}

type Credentials struct {
	Username string `json:"username" validate:"required,min=4,max=15"`
	Password string `json:"password" validate:"required,min=4"`
}

type Claims struct {
	Username string `json:"username" validate:"required"`
	jwt.StandardClaims
}
