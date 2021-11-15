package viper

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Host       string
	Drivername string
	Username   string
	Password   string
	DbName     string
	Port       string
}

func GetDbconfigs() (DbConfig, error) {
	// Config
	log.Println("fetching db configs")
	v := viper.New()
	v.SetConfigName("config") // config file name
	v.SetConfigType("yaml")
	v.AddConfigPath("./config") // config file path
	err := v.ReadInConfig()
	if err != nil {
		return DbConfig{}, err
	}

	drivername := v.GetString("db.drivername")
	username := v.GetString("db.username")
	password := v.GetString("db.password")
	host := v.GetString("db.host")
	port := v.GetString("db.port")
	dbname := v.GetString("db.dbName")

	return DbConfig{
		Host:       host,
		Port:       port,
		Drivername: drivername,
		Username:   username,
		Password:   password,
		DbName:     dbname,
	}, nil
}

func GetJwtKey() (key []byte) {
	log.Println("fetching jwt-key configs")
	v := viper.New()
	v.SetConfigName("config") // config file name
	v.SetConfigType("yaml")
	v.AddConfigPath("./config") // config file path

	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("unable to fetch jwtKey err :", err.Error())
	}
	key = []byte(v.GetString("key.jwtKey"))
	return key
}

func GetPort() (port string) {
	log.Println("fetching port configs")
	v := viper.New()
	v.SetConfigName("config") // config file name
	v.SetConfigType("yaml")
	v.AddConfigPath("./config") // config file path

	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("Unable to fetch port")
	}
	port = v.GetString("bookrepo.port")
	return port
}

//func GetCookieExpiryTime() (CookieTime time.Duration) {
//	log.Println("fetching Cookie configs")
//	v := viper.New()
//	v.SetConfigName("config") // config file name
//	v.SetConfigType("yaml")
//	v.AddConfigPath("./config") // config file path
//
//	err := v.ReadInConfig()
//	if err != nil {
//		log.Fatal("Unable to fetch Cookie:")
//	}
//	//CookieExpiryDuration := CookieTime
//	CookieTime = v.GetDuration("expirytime.cookieexpirytime")
//	return CookieTime
//}
