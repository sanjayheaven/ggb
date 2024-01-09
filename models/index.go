package models

import (
	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(address string) {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               address,
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{})

	// Migrate the schema
	db.AutoMigrate(&Example{})

	// export DB
	DB = db

	log.Printf(`🍟: Successfully connected to Mysql at ` + address)
	if err != nil {
		panic(`😫: Connected failed, check your Mysql with ` + address)
	}

}
