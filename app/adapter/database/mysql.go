package database

import (
	model "swag-tutor/app/models/sql"
	"swag-tutor/app/shared/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	DB *gorm.DB
}

func SetupDatabase() *MySQL {

	url := config.DB.MySQL.DataSourceName
	db, err := gorm.Open(mysql.Open(url))

	db.AutoMigrate(&model.Person{})

	if err != nil {
		panic("")
	}

	Db, err := db.DB()

	if err != nil {
		panic("")
	}

	Db.Ping()

	return &MySQL{
		DB: db,
	}
}
