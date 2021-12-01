package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	DBConfig   DbConfig `yaml:"db"`
	jwtKey     string   `yaml:"jwtKey"`
	ServerAddr string   `yaml:"server_address"`
}

type DbConfig struct {
	Drivername string `yaml:"drivername"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Dbname     string `yaml:"dbname"`
}

func LoadConfigFromFile(FileName string) (*Config, error) {
	fileData, err := ioutil.ReadFile(FileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Read config file error: %+v", err.Error()))
	}
	cfgData := &Config{}
	err = yaml.Unmarshal(fileData, cfgData)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Config file unmarshall error: %+v", err.Error()))
	}
	fmt.Printf("%+v\n", cfgData)
	return cfgData, nil
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
