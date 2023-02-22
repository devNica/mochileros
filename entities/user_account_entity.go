package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserAccount struct {
	Id             uuid.UUID      `gorm:"primaryKey;column:id;type:varchar(36)"`
	Email          string         `gorm:"column:email;type:varchar(200);not null;unique"`
	Password       string         `gorm:"column:password;type:varchar(255);not null"`
	PhoneNumber    string         `gorm:"column:phone_number;type:varchar(20);not null;unique"`
	IsActive       bool           `gorm:"column:is_active;type:bool;not null;default:true"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	UserKYC        UserInfo       `gorm:"foreignKey:user_id"`
	OwnerHotel     []Hotel        `gorm:"foreignKey:owner_id"`
	UserHasProfile []UserProfiles `gorm:"foreignKey:user_id"`
	UserHasAsset   []UserAssets   `gorm:"foreignKey:user_id"`
	StatusId       uint8          `gorm:"column:status_id;primaryKey"`
}

type UserInfo struct {
	Id        uint32    `gorm:"primaryKey;column:id;autoIncrement;not null;unique"`
	FirstName string    `gorm:"column:first_name;type:varchar(100);not null"`
	LastName  string    `gorm:"column:last_name;type:varchar(100);not null"`
	Address   string    `gorm:"column:address;type:varchar(255);not null"`
	Birthdate string    `gorm:"column:birthdate;type:varchar(10);not null"`
	DNI       string    `gorm:"column:dni;type:varchar(100);not null;unique"`
	CountryID uint16    `gorm:"column:country_id;primaryKey"`
	UserId    uuid.UUID `gorm:"column:user_id;primaryKey"`
}

type AccountStatus struct {
	Id         uint8         `gorm:"primaryKey;column:id;autoIncrement;not null;unique"`
	Status     string        `gorm:"column:status;type:varchar(100);not null"`
	UserStatus []UserAccount `gorm:"foreignKey:status_id"`
}
