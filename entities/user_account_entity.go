package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserAccount struct {
	Id         uuid.UUID        `gorm:"primaryKey;column:id;type:varchar(36)"`
	Email      string           `gorm:"column:email;type:varchar(200);not null;unique"`
	Password   string           `gorm:"column:password;type:varchar(255);not null"`
	IsActive   bool             `gorm:"column:is_active;type:bool;not null;default:true"`
	CreatedAt  time.Time        `gorm:"column:created_at"`
	UpdatedAt  time.Time        `gorm:"column:updated_at"`
	UserKYC    UserInfo         `gorm:"foreignKey:user_id"`
	OwnerHotel []Hotel          `gorm:"foreignKey:owner_id"`
	UHP        []UserHasProfile `gorm:"foreignKey:user_id"`
}

type UserInfo struct {
	Id        uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36)"`
	FirstName string    `gorm:"column:first_name;type:varchar(100);not null"`
	LastName  string    `gorm:"column:last_name;type:varchar(100);not null"`
	UserId    uuid.UUID `gorm:"column:user_id;primaryKey"`
}
