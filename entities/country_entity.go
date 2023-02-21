package entities

import "github.com/jackc/pgtype"

type Country struct {
	Id          uint16       `gorm:"primaryKey;autoIncrement;unique;not null"`
	Name        string       `gorm:"column:name;type:varchar(255);not null;unique"`
	Capital     string       `gorm:"column:capital;type:varchar(200);not null"`
	Cca3        string       `gorm:"column:cca3;type:varchar(10);not null"`
	CallingCode string       `gorm:"column:callingcode;type:varchar(10);not null"`
	TimeZones   pgtype.JSONB `gorm:"column:timezones;type:jsonb;default:'{}'"`
	States      pgtype.JSONB `gorm:"column:states;type:jsonb;default:'[]'"`
	Latitude    string       `gorm:"column:latitude;type:varchar(15);not null"`
	Longitude   string       `gorm:"column:longitude;type:varchar(15);not null"`
	FlagPng     string       `gorm:"column:flagpng;type:varchar(255);not null"`
	FlagSvg     string       `gorm:"column:flagsvg;type:varchar(255);not null"`
	CurrCode    string       `gorm:"column:currcode;type:varchar(6);not null"`
	CurrName    string       `gorm:"column:currname;type:varchar(50);not null"`
	CurrSymbol  string       `gorm:"column:currsymbol;type:varchar(10);not null"`
	HotelInfo   []Hotel      `gorm:"foreignKey:country_id"`
	KYCInfo     UserInfo     `gorm:"foreingkey:country_id"`
}
