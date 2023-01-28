package entities

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	Id        uint16           `gorm:"primaryKey;column:id;autoIncrement;not null;unique"`
	Profile   string           `gorm:"column:profile;type:varchar(20);not null"`
	CreatedAt time.Time        `gorm:"column:created_at"`
	PHU       []UserHasProfile `gorm:"foreignKey:profile_id"`
}

type UserHasProfile struct {
	UserId    uuid.UUID `gorm:"column:user_id;primaryKey"`
	ProfileId uint16    `gorm:"column:profile_id;primaryKey"`
	IsActive  bool      `gorm:"column:is_active;type:bool;not null;default:true"`
}
