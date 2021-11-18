package driver

import (
	"containerization/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type DataStore struct {
	Db *sqlx.DB
}

func ConnectDB() (*DataStore, error) {
	//defconfig,err := vipers.LoadConfigFromFile()
	dbConf, err := config.LoadConfigFromFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	dbInstance := &DataStore{}
	//	dbInstance.Db, err = sqlx.Connect(dbConf.Drivername, dbConf.Username+":"+dbConf.Password+"@tcp("+dbConf.Host+":"+dbConf.Port+")/"+dbConf.DbName)
	dbInstance.Db, err = sqlx.Open(dbConf.DBConfig.Drivername, dbConf.DBConfig.Username+":"+dbConf.DBConfig.Password+"@tcp("+dbConf.DBConfig.Host+":"+dbConf.DBConfig.Port+")/"+dbConf.DBConfig.Dbname+"?parseTime=true")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Println("Database : connected successfully")
	return dbInstance, nil
}
