package models

type Person struct {
	Username  string `json:"username" validate:"min=0"`
	Password  string `json:"password" validate:"min=0,max=15"`
	Firstname string `json:"firstname" validate:"min=0,max=15"`
	Lastname  string `json:"lastname" validate:"min=0,max=15"`
	Age       int    `json:"age" validate:"min=0,max=130"`
	Gender    string `json:"gender" validate:"min=0,max=15"`
	City      string `json:"city" validate:"min=0,max=15"`
	Country   string `json:"country" validate:"min=0"`
	Phone     string `json:"phone" validate:"min=0,max=10"`
}
