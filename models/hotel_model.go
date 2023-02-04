package models

import "github.com/google/uuid"

type HotelRequestModel struct {
	NameHotel          string `json:"nameHotel" validate:"required"`
	Address            string `json:"address" validate:"required"`
	ServicePhoneNumber string `json:"servicePhoneNumber" validate:"required"`
	CountryID          uint16 `json:"countryId" validate:"required"`
	State              string `json:"state" validate:"required"`
	Province           string `json:"province" validate:"required"`
	OwnerId            string `json:"ownerId" validate:"required"`
}

type HotelResponseModel struct {
	HotelId            uuid.UUID
	NameHotel          string
	Address            string
	ServicePhoneNumber string
	Country            string
	State              string
	Province           string
}
