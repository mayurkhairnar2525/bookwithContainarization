package user

import (
	"containerization/driver"
	"containerization/models"
	"log"
)

type UserRepository struct{}

func (u UserRepository) CreateUser(user models.Person) (string, error) {
	datastore, _ := driver.ConnectDB()
	err := datastore.Db.QueryRow("INSERT INTO usermanagement(username,password,firstname,lastname,age,gender,city,country,phone) VALUES(?,?,?,?,?,?,?,?,?)", user.Username,
		user.Password, user.Firstname, user.Lastname, user.Age, user.Gender, user.City, user.Country, user.Phone)
	if err == nil {
		log.Println("err occurred", err)
	}
	return user.Username, nil
}

func (u UserRepository) GetUsers(user models.Person, users []models.Person) ([]models.Person, error) {
	datastore, _ := driver.ConnectDB()
	result, err := datastore.Db.Query("SELECT username,password,firstname,lastname,age,gender,city,country,phone from usermanagement")
	if err != nil {
		log.Println("err", err)
	}
	for result.Next() {
		err = result.Scan(&user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.Age, &user.Gender, &user.City, &user.Country, &user.Phone)
		users = append(users, user)
	}
	if err != nil {
		log.Println("err:", err)
	}
	return users, nil
}

func (u UserRepository) DeleteUser(username string) (string, error) {
	datastore, _ := driver.ConnectDB()
	result, err := datastore.Db.Exec("DELETE FROM usermanagement WHERE username  = ?", username)
	if err != nil {
		return string(0), err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return string(0), err
	}
	return string(rowsDeleted), nil
}
