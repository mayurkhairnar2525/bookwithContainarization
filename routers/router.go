package routers

import (
	"containerization/auth"
	"containerization/controllers"
	"database/sql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func Router(controllers controllers.Controllers) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", auth.IsAuthorised(controllers.GetBooks())).Methods("GET")
	router.HandleFunc("/books", auth.IsAuthorised(controllers.CreateBook())).Methods("POST")
	router.HandleFunc("/books/{id}", auth.IsAuthorised(controllers.GetBookById())).Methods("GET")
	router.HandleFunc("/books/author/{author}", auth.IsAuthorised(controllers.GetBookByAuthor())).Methods("GET")
	router.HandleFunc("/books", auth.IsAuthorised(controllers.UpdateBook())).Methods("PUT")
	router.HandleFunc("/books/{name}", auth.IsAuthorised(controllers.DeleteBook())).Methods("DELETE")

	router.HandleFunc("/login", controllers.Login)

	router.HandleFunc("/register", auth.IsAuthorised(controllers.CreateUser())).Methods("POST")
	router.HandleFunc("/users", auth.IsAuthorised(controllers.GetUsers())).Methods("GET")
	router.HandleFunc("/users/{username}", auth.IsAuthorised(controllers.DeleteUser())).Methods("DELETE")
	return router
}
