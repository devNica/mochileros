package response

import "github.com/jackc/pgtype"

type CountryResponseModel struct {
	Id          uint16
	Name        string
	Capital     string
	Cca3        string
	CallingCode string
	TimeZones   pgtype.JSONB
	States      pgtype.JSONB
	Latitude    string
	Longitude   string
	FlagPng     string
	FlagSvg     string
	CurrCode    string
	CurrName    string
	CurrSymbol  string
}
