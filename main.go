package main

import (
	"containerization/auth"
	"containerization/controllers"
	"containerization/driver"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"

	"net/http"
)

var db *sql.DB

func main() {
	db, _ := driver.ConnectDB()
	log.Println("Db connected", db)
	controllers := controllers.Controllers{}
	router := initRouter(controllers)

	fmt.Println("Server is on port 9099:")
	http.ListenAndServe(":9099", router)
}

func initRouter(controllers controllers.Controllers) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", auth.IsAuthorised(controllers.GetBooks(db))).Methods("GET")
	router.HandleFunc("/books", auth.IsAuthorised(controllers.CreateBook(db))).Methods("POST")
	router.HandleFunc("/books/{id}", auth.IsAuthorised(controllers.GetBookById(db))).Methods("GET")
	router.HandleFunc("/books/author/{author}", auth.IsAuthorised(controllers.GetBookByAuthor(db))).Methods("GET")
	router.HandleFunc("/books", auth.IsAuthorised(controllers.UpdateBook(db))).Methods("PUT")
	router.HandleFunc("/books/{id}", auth.IsAuthorised(controllers.DeleteBook(db))).Methods("DELETE")
	router.HandleFunc("/login", controllers.Login)
	return router
}
