package entities

import (
	"time"

	"github.com/google/uuid"
)

type Hotel struct {
	Id                 uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36)"`
	NameHotel          string    `gorm:"index;column:name_hotel;type:varchar(100);unique;not null"`
	Address            string    `gorm:"column:address;type:varchar(200);not null"`
	ServicePhoneNumber string    `gorm:"column:service_phone_number;type:varchar(20);not null"`
	Country            string    `gorm:"column:country;type:varchar(100);not null"`
	State              string    `gorm:"column:state;type:varchar(200)"`
	IsActive           bool      `gorm:"column:is_active;type:bool;not null;default:true"`
	Province           string    `gorm:"column:province;type:varchar(200);not null"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	OwnerId            uuid.UUID `gorm:"column:owner_id;primaryKey"`
}
