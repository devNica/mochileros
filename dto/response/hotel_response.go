package response

import "github.com/google/uuid"

type HotelRepositoryResponseModel struct {
	HotelId            uuid.UUID
	NameHotel          string
	Address            string
	ServicePhoneNumber string
	Country            string
	State              string
	Province           string
	Filename           uuid.UUID
	Filetype           string
}

type HotelResponseModel struct {
	HotelId            uuid.UUID
	NameHotel          string
	Address            string
	ServicePhoneNumber string
	Country            string
	State              string
	Province           string
	Url                string
}

type FileResponseModel struct {
	Filetype string
	Binary   []byte
}
