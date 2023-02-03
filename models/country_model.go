package models

import "github.com/jackc/pgtype"

type Country struct {
	Id          uint16       `json:"id"`
	Name        string       `json:"name"`
	Capital     string       `json:"capital"`
	Cca3        string       `json:"cca3"`
	CallingCode string       `json:"callingcode"`
	TimeZones   pgtype.JSONB `json:"timezones"`
	States      pgtype.JSONB `json:"states"`
	Latitude    string       `json:"latitude"`
	Longitude   string       `json:"longitude"`
	FlagPng     string       `json:"flagpng"`
	FlagSvg     string       `json:"flagsvg"`
	CurrCode    string       `json:"currcode"`
	CurrName    string       `json:"currname"`
	CurrSymbol  string       `json:"currsymbol"`
}
