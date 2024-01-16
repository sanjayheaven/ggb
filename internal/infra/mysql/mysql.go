package mysql

import (
	"fmt"
	"go-gin-boilerplate/configs"
	"go-gin-boilerplate/internal/models"
	"go-gin-boilerplate/internal/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *configs.Mysql) *gorm.DB {

	logger := logger.Logger

	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DataBase,
	)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               address,
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{})

	// Migrate the schema
	migrateErr := db.AutoMigrate(&models.Example{}, &models.User{})

	if migrateErr != nil {
		panic(`😫: Auto migrate failed, check your Mysql with ` + address)
	}

	// export DB
	DB = db

	logger.Printf(`🍟: Successfully connected to Mysql at ` + address)

	if err != nil {
		panic(`😫: Connected failed, check your Mysql with ` + address)
	}

	return db

}
