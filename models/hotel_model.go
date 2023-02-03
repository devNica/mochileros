package models

type HotelRequestModel struct {
	NameHotel          string `json:"nameHotel" validate:"required"`
	Address            string `json:"address" validate:"required"`
	ServicePhoneNumber string `json:"servicePhoneNumber" validate:"required"`
	CountryID          uint16 `json:"countryId" validate:"required"`
	State              string `json:"state" validate:"required"`
	Province           string `json:"province" validate:"required"`
	OwnerId            string `json:"ownerId" validate:"required"`
}
