package controllers

import (
	"containerization/models"
	book2 "containerization/repository/book"
	"containerization/utils"
	"containerization/validation"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var books []models.BookManagement

func (u Controllers) GetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		bookRepo := book2.BookRepository{}
		books, err := bookRepo.GetBooks(book, books)

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

func (u Controllers) CreateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		var bookCreate int
		var error models.Error
		err := json.NewDecoder(r.Body).Decode(&book)

		validationError := validation.ValidateCreateBook(&book)
		if validationError != nil {
			validation.DisplayError(w, validationError)
			return
		}
		bookRepo := book2.BookRepository{}
		bookCreate, err = bookRepo.CreateBook(book)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, bookCreate)
	}
}

func (u Controllers) GetBookById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		params := mux.Vars(r)
		books = []models.BookManagement{}
		bookRepo := book2.BookRepository{}

		id, _ := strconv.Atoi(params["id"])
		book, err := bookRepo.GetBookById(book, id)

		v := validator.New()
		if err = v.Struct(book); err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

func (u Controllers) GetBookByAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		params := mux.Vars(r)
		bookRepo := book2.BookRepository{}

		author, _ := params["author"]
		book, err := bookRepo.GetBookByAuthor(book, author)
		v := validator.New()
		if err = v.Struct(book); err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)

	}
}

func (u Controllers) UpdateBook() http.HandlerFunc {
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
		rowsUpdated, err := bookRepo.UpdateBook(book)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (u Controllers) DeleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		params := mux.Vars(r)
		bookRepo := book2.BookRepository{}
		name, _ := params["name"]
		rowsDeleted, err := bookRepo.DeleteBook(name)
		v := validator.New()

		if err = v.Struct(book); err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}
