package database

import (
	"auth_microservice/internal/config"
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Gorm *gorm.DB

func GORMConnect() {
	dsn := url.URL{
		User:     url.UserPassword(config.Env.Database.User, config.Env.Database.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", config.Env.Database.Host, config.Env.Database.Port),
		Path:     config.Env.Database.DatabaseName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Panic("failed to connect database:", err.Error())
	}

	Gorm = db
}
