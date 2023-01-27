package configurations

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DatabaseConnect(config Config) *gorm.DB {
	username := config.Get("DATABASE_USERNAME")
	password := config.Get("DATABASE_PASSWORD")
	host := config.Get("DATABASE_HOST")
	port := config.Get("DATABASE_PORT")
	dbName := config.Get("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("Failed to conect to database")
	}

	//AutoMigrate

	return db

}
