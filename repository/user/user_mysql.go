package user

import (
	"containerization/models"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Register interface {
	AddUser(user models.Person) error
	userAlreadyExists(username string) bool
}

var alreadyExistsError = "user is already present please choose another username"

type DataStore struct {
	Db *sqlx.DB
}

//AddUser will register user's details in userRepository table and return error if occurs
func (repository DataStore) AddUser(p *models.Person) error {

	if repository.userAlreadyExists(p.Username) {
		return errors.New(fmt.Sprintf("%s %s", p.Username, alreadyExistsError))
	}
	query := `INSERT INTO user (username,password,firstname,lastname,age,gender,city,country,phone) VALUES
			(?,?,?,?,?,?,?,?,?)`
	//insert user's registration details into userRepository
	_, err := repository.Db.Exec(query, p.Username, p.Password, p.Firstname, p.Lastname, p.Age, p.Gender, p.City, p.Country, p.Phone)
	return err
}

//userAlreadyExists checks if user already present or not
func (repository DataStore) userAlreadyExists(username string) bool {
	query := `SELECT username FROM user WHERE username =?`
	var returnedUser string
	err := repository.Db.QueryRowx(query, username).Scan(&returnedUser)
	if err != nil && returnedUser == "" {
		return false
	}
	if username == returnedUser {
		return true
	}
	return false
}
