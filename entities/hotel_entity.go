package entities

import (
	"time"

	"github.com/google/uuid"
)

type Hotel struct {
	Id                 uuid.UUID     `gorm:"primaryKey;column:id;type:varchar(36);unique"`
	NameHotel          string        `gorm:"index;column:name_hotel;type:varchar(100);unique;not null"`
	Address            string        `gorm:"column:address;type:varchar(200);not null"`
	ServicePhoneNumber string        `gorm:"column:service_phone_number;type:varchar(20);not null"`
	State              string        `gorm:"column:state;type:varchar(200)"`
	Province           string        `gorm:"column:province;type:varchar(200);not null"`
	CreatedAt          time.Time     `gorm:"column:created_at"`
	OwnerId            uuid.UUID     `gorm:"column:owner_id;primaryKey"`
	CountryID          uint16        `gorm:"column:country_id;primaryKey"`
	StatusID           uint16        `gorm:"column:status_id;primaryKey;default:1"`
	HotelHasAsset      []HotelAssets `gorm:"foreignKey:hotel_id"`
}

type HotelStatus struct {
	Id        uint16  `gorm:"primaryKey;column:id;autoIncrement;not null;unique"`
	Status    string  `gorm:"column:status;type:varchar(50);not null; unique"`
	HotelInfo []Hotel `gorm:"foreignKey:status_id"`
}
