package models

import (
	"log"
	"time"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type BasicModel struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

var DB *gorm.DB

func Connect(address string) {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               address,
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{})

	// Migrate the schema
	migrateErr := db.AutoMigrate(&Example{})
	if migrateErr != nil {
		panic(`😫: Auto migrate failed, check your Mysql with ` + address)
	}

	// export DB
	DB = db

	log.Printf(`🍟: Successfully connected to Mysql at ` + address)
	if err != nil {
		panic(`😫: Connected failed, check your Mysql with ` + address)
	}

}
