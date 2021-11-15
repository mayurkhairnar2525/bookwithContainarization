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
	router.HandleFunc("/books", auth.IsAuthorised(controllers.GetBooks(db))).Methods("GET")
	router.HandleFunc("/books", auth.IsAuthorised(controllers.CreateBook(db))).Methods("POST")
	router.HandleFunc("/books/{id}", auth.IsAuthorised(controllers.GetBookById(db))).Methods("GET")
	router.HandleFunc("/books/author/{author}", auth.IsAuthorised(controllers.GetBookByAuthor(db))).Methods("GET")
	router.HandleFunc("/books", auth.IsAuthorised(controllers.UpdateBook(db))).Methods("PUT")
	router.HandleFunc("/books/{id}", auth.IsAuthorised(controllers.DeleteBook(db))).Methods("DELETE")
	router.HandleFunc("/login", controllers.Login)
	router.HandleFunc("/register", controllers.Registration).Methods("POST")

	return router
}

// 	r.HandleFunc("/register", handler.Registration).Methods("POST")
