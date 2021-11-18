package validation

import (
	"containerization/models"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func ValidateCreateBook(book *models.BookManagement) error {
	log.Println("Validate CreateBook")
	v := validator.New()
	err := v.Struct(book)
	return err
}

func ValidateCreateUser(person *models.Person) error {
	log.Println("validate details entered by user")
	v := validator.New()
	err := v.Struct(person)
	return err
}

//display validation errors
func DisplayError(w http.ResponseWriter, errs error) {
	for _, e := range errs.(validator.ValidationErrors) {
		fmt.Fprintln(w, e)
	}
}
