package main

import (
	"containerization/config"
	"containerization/controllers"
	"containerization/driver"
	"containerization/routers"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var configFile = flag.String("config", "config/config.yaml", "Config File")

func main() {
	db, err := driver.ConnectDB()
	if err != nil {
		log.Fatal("not able connect with server database: application terminated")
	}
	log.Println("Db connected", db)

	// fetching the config's
	cfg, err := config.LoadConfigFromFile(*configFile)
	if err != nil {
		log.Fatalf("Failed to load config: %+v\n", err.Error())
	}
	fmt.Println("Fetching the Configs", cfg)
	controllers := controllers.Controllers{}
	router := routers.Router(controllers)
	fmt.Println("Server is on port 9099:")
	log.Fatal(http.ListenAndServe(cfg.ServerAddr, router))
}
