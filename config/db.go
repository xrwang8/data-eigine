package conf

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(conf Config, logger Logger) Database {
	username := conf.Database.User
	password := conf.Database.Password
	host := conf.Database.Host
	port := conf.Database.Port
	dbname := conf.Database.DBName

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})
	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}
	logger.Info("Database connection established")
	return Database{
		DB: db,
	}
}
