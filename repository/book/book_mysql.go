package book

import "C"
import (
	"containerization/driver"
	"containerization/models"
	"database/sql"
	"log"
)

type BookRepository struct{}

const (
	tableName        = "table:bookmanagement"
	executionSuccess = "status:success"
	executionFailure = "status:failed"
)

func (b BookRepository) GetBooks(db *sql.DB, book models.BookManagement, books []models.BookManagement) ([]models.BookManagement, error) {
	datastore, _ := driver.ConnectDB()
	result, err := datastore.Db.Query("SELECT id, name, author, prices,available,pagequality,lauchedyear, isbn,stock from bookmanagement")
	if err != nil {
		log.Println("err", err)
	}
	for result.Next() {
		err = result.Scan(&book.ID, &book.Name, &book.Author, &book.Prices, &book.Available, &book.PageQuality, &book.LaunchedYear, &book.Isbn, &book.Stock)
		books = append(books, book)
	}
	if err != nil {
		log.Println("err:", err)
	}
	return books, nil
}

func (b BookRepository) GetBookById(db *sql.DB, book models.BookManagement, id int) (models.BookManagement, error) {
	datastore, _ := driver.ConnectDB()
	rows := datastore.Db.QueryRow("SELECT id, name, author, prices, available, pagequality, lauchedyear, isbn,stock FROM bookmanagement where id = ?", id)
	err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Prices, &book.Available, &book.PageQuality, &book.LaunchedYear, &book.Isbn, &book.Stock)

	return book, err
}

func (b BookRepository) GetBookByAuthor(db *sql.DB, book models.BookManagement, author string) (models.BookManagement, error) {
	datastore, _ := driver.ConnectDB()
	rows := datastore.Db.QueryRow("SELECT  id,name, author, prices, available, pagequality, lauchedyear, isbn,stock FROM bookmanagement where author = ?", author)
	err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Prices, &book.Available, &book.PageQuality, &book.LaunchedYear, &book.Isbn, &book.Stock)
	return book, err
}

func (b BookRepository) CreateBook(db *sql.DB, book models.BookManagement) (int, error) {
	datastore, _ := driver.ConnectDB()
	err := datastore.Db.QueryRow("INSERT INTO bookmanagement(id,name,author,prices,available,pagequality,lauchedyear,isbn,stock) VALUES(?,?,?,?,?,?,?,?,?)", book.ID, book.Name, book.Author, book.Prices, book.Available, book.PageQuality, book.LaunchedYear, book.Isbn, book.Stock)
	if err == nil {
		log.Println("err occurred", err)
	}
	return book.ID, nil
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.BookManagement) (int64, error) {
	datastore, _ := driver.ConnectDB()
	result, err := datastore.Db.Exec("UPDATE bookmanagement SET id=?,name =?, author=?,prices=?,available=?,pagequality=?,lauchedyear=?,stock=? WHERE isbn= ?",
		&book.ID, &book.Name, &book.Author, &book.Prices, &book.Available, &book.PageQuality, &book.LaunchedYear, &book.Stock, &book.Isbn)
	if err != nil {
		log.Println("err", err)
	}
	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("err", err)
	}
	return rowsUpdated, nil
}

func (b BookRepository) DeleteBook(db *sql.DB, id int) (int64, error) {
	datastore, _ := driver.ConnectDB()
	result, err := datastore.Db.Exec("DELETE FROM bookmanagement WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, nil
}
