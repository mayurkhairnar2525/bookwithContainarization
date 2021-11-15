package models

import "github.com/dgrijalva/jwt-go"

type BookManagement struct {
	ID           int    `json:"id" validate:"min=0,max=130"`
	Name         string `json:"name" validator:"nonzero"`
	Author       string `json:"author"Validate:"empty=true | gte=2 & lte=15"`
	Prices       int    `json:"prices" validator:"nonzero"`
	Available    string `json:"available" validator:"nonzero"`
	PageQuality  string `json:"pagequality" validator:"nonzero"`
	LaunchedYear string `json:"launchedyear" validator:"empty=true | gte=2 & lte=4"`
	Isbn         string `json:"isbn" validator:"min=21"`
	Stock        int    `json:"stock" validator:"nonzero"`
}

type Credentials struct {
	Username string `json:"username" validate:"required,min=4,max=15"`
	Password string `json:"password" validate:"required,min=4"`
}

type Claims struct {
	Username string `json:"username" validate:"required"`
	jwt.StandardClaims
}
