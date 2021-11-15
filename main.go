package main

import (
	"containerization/controllers"
	"containerization/driver"
	"containerization/routers"
	"containerization/viper"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	db, err := driver.ConnectDB()
	if err != nil {
		log.Fatal("not able connect with server database: application terminated")
	}
	log.Println("Db connected", db)
	controllers := controllers.Controllers{}
	router := routers.Router(controllers)
	port := viper.GetPort()
	fmt.Println("Server is on port 9099:")
	log.Fatal(http.ListenAndServe(port, router))
}
