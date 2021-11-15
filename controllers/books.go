package controllers

import (
	"containerization/models"
	book2 "containerization/repository/book"
	"containerization/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Controllers struct{}

var books []models.BookManagement

func (c Controllers) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		books = []models.BookManagement{}
		bookRepo := book2.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		v := validator.New()
		if err = v.Struct(book); err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(books)
		if err != nil {
			log.Println("err", err)
			return
		}
	}
}

func (c Controllers) CreateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		var bookID int
		var error models.Error
		err := json.NewDecoder(r.Body).Decode(&book)

		v := validator.New()
		if err = v.Struct(book); err != nil {
			fmt.Println(err)
			return
		}
		bookRepo := book2.BookRepository{}
		bookID, err = bookRepo.CreateBook(db, book)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controllers) GetBookById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		params := mux.Vars(r)
		books = []models.BookManagement{}
		bookRepo := book2.BookRepository{}

		id, _ := strconv.Atoi(params["id"])
		book, err := bookRepo.GetBookById(db, book, id)

		v := validator.New()
		if err = v.Struct(book); err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

func (c Controllers) GetBookByAuthor(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		params := mux.Vars(r)
		bookRepo := book2.BookRepository{}

		author, _ := params["author"]
		book, err := bookRepo.GetBookByAuthor(db, book, author)
		v := validator.New()
		if err = v.Struct(book); err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)

	}
}

func (c Controllers) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		var error models.Error

		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		if book.ID == 0 || book.Name == "" || book.Author == "" || book.Prices == 0 || book.Available == "" || book.PageQuality == "" || book.LaunchedYear == "" || book.Isbn == "" || book.Stock == 0 {
			error.Message = "All fields are required."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		bookRepo := book2.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c Controllers) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		//	var error models.Error
		params := mux.Vars(r)
		bookRepo := book2.BookRepository{}
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Println("err", err)
			return
		}
		rowsDeleted, err := bookRepo.DeleteBook(db, id)
		v := validator.New()

		if err = v.Struct(book); err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}
