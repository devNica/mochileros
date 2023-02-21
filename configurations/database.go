package configurations

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DatabaseConnect(config Config) *gorm.DB {
	username := config.Get("DATABASE_USERNAME")
	password := config.Get("DATABASE_PASSWORD")
	host := config.Get("DATABASE_HOST")
	port := config.Get("DATABASE_PORT")
	dbName := config.Get("DATABASE_NAME")
	sslMode := config.Get("DATABASE_SSL_MODE")

	// for mysql connection
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// })

	// for postgres connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, dbName, sslMode)

	fmt.Println("DSN", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("Failed to conect to database")
	}

	//AutoMigrate
	// err = db.AutoMigrate(&entities.Country{})
	// err = db.AutoMigrate(&entities.Profile{})
	// err = db.AutoMigrate(&entities.AssetType{})
	// err = db.AutoMigrate(&entities.File{})
	// err = db.AutoMigrate(&entities.UserAccount{})
	// err = db.AutoMigrate(&entities.UserInfo{})
	// err = db.AutoMigrate(&entities.UserAssets{})
	// err = db.AutoMigrate(&entities.UserProfiles{})
	// err = db.AutoMigrate(&entities.HotelStatus{})
	// err = db.AutoMigrate(&entities.Hotel{})
	// err = db.AutoMigrate(&entities.HotelAssets{})

	return db

}
