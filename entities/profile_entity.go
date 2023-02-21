package entities

import (
	"github.com/google/uuid"
)

type Profile struct {
	Id             uint16         `gorm:"primaryKey;column:id;autoIncrement;not null;unique"`
	Profile        string         `gorm:"column:profile;type:varchar(20);not null"`
	ProfileHasUser []UserProfiles `gorm:"foreignKey:profile_id"`
}

type UserProfiles struct {
	UserId    uuid.UUID `gorm:"column:user_id;primaryKey"`
	ProfileId uint16    `gorm:"column:profile_id;primaryKey"`
	IsActive  bool      `gorm:"column:is_active;type:bool;not null;default:true"`
}
