package entities

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	Filename      uuid.UUID     `gorm:"primaryKey;column:filename;type:varchar(36);unique"`
	Filetype      string        `gorm:"column:filetype;type:varchar(10)"`
	Filesize      int           `gorm:"column:filesize;type:int4"`
	Binary        []byte        `gorm:"column:binary;type:bytea"`
	CreatedAt     time.Time     `gorm:"column:created_at"`
	AssetHasHotel []HotelAssets `gorm:"foreignKey:file_id;references:filename"`
	AssetHasUser  []UserAssets  `gorm:"foreignKey:file_id;references:filename"`
}

type AssetType struct {
	Id        uint16 `gorm:"primaryKey;column:id;autoIncrement;not null;unique"`
	AssetType string `gorm:"column:asset_type;type:varchar(100);not null;unique"`
	UserAsset UserAssets `gorm:"foreignKey:asset_type_id"`
}

type HotelAssets struct {
	HotelId uuid.UUID `gorm:"column:hotel_id;primaryKey"`
	FileId  uuid.UUID `gorm:"column:file_id;primaryKey"`
}

type UserAssets struct {
	UserId      uuid.UUID `gorm:"column:user_id;primaryKey"`
	FileId      uuid.UUID `gorm:"column:file_id;primaryKey"`
	AssetTypeId uint16    `gorm:"column:asset_type_id;primaryKey"`
}
