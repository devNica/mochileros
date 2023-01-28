package entities

import "github.com/jackc/pgtype"

type Country struct {
	Id          uint16       `gorm:"primaryKey;autoIncrement;unique;not null"`
	Name        string       `gorm:"column:name;type:varchar(100);not null;unique"`
	Capital     string       `gorm:"column:capital;type:varchar(120);not null"`
	Cca3        string       `gorm:"column:cca3;type:varchar(3);not null"`
	CallingCode string       `gorm:"column:callingcode;type:varchar(10);not null"`
	TimeZones   pgtype.JSONB `gorm:"column:timezones;type:jsonb;not null"`
	State       pgtype.JSONB `gorm:"column:states;type:jsonb;not null"`
	Latitude    string       `gorm:"column:latitude;type:varchar(15);not null"`
	Longitude   string       `gorm:"column:longitude;type:varchar(15);not null"`
	FlagPng     string       `gorm:"column:flagpng;type:varchar(60);not null"`
	FlagSvg     string       `gorm:"column:flagsvg;type:varchar(60);not null"`
	CurrCode    string       `gorm:"column:currcode;type:varchar(6);not null"`
	CurrName    string       `gorm:"column:currname;type:varchar(50);not null"`
	CurrSymbol  string       `gorm:"column:currsymbol;type:varchar(3);not null"`
}
