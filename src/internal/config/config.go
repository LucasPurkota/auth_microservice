package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

var Env Config

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database struct {
		Host         string `json:"host"`
		Port         int    `json:"port"`
		Driver       string `json:"driver"`
		User         string `json:"user"`
		Password     string `json:"password"`
		DatabaseName string `json:"name"`
	}
}

func LoadConfig() {
	ex, err := os.Executable()
	if err != nil {
		log.Panic("Unable to get executable path:", err)
	}
	exPath := filepath.Dir(ex)

	fileConf := filepath.Join(exPath, "auth_microservice.conf")
	cfg, err := ini.Load(fileConf)
	if err != nil {
		log.Panic("Unable to load config file:", err)
	}

	//SERVICE
	Env.Host = cfg.Section("SERVICE").Key("HOST").String()
	Env.Port = cfg.Section("SERVICE").Key("PORT").String()

	//DATABASE
	Env.Database.Host = cfg.Section("DATABASE").Key("DB_HOST").String()
	Env.Database.Port = cfg.Section("DATABASE").Key("DB_PORT").MustInt()
	Env.Database.Driver = cfg.Section("DATABASE").Key("DB_DRIVER").String()
	Env.Database.User = cfg.Section("DATABASE").Key("DB_USERNAME").String()
	Env.Database.Password = cfg.Section("DATABASE").Key("DB_PASSWORD").String()
	Env.Database.DatabaseName = cfg.Section("DATABASE").Key("DB_NAME").String()
}
