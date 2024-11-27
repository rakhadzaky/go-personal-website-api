package configs

import (
	"fmt"
	"log"

	"go-personal-website-api/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DBEngine   string `env:"DB_ENGINE"`
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
}

func InitiateDatabase(config DatabaseConfig) (dbConn *gorm.DB, err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(constants.LogTypeError, ", with error: ", err.Error())
		return nil, err
	}

	return db, nil
}
