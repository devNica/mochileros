package response

import "github.com/google/uuid"

type HotelResponseModel struct {
	HotelId            uuid.UUID
	NameHotel          string
	Address            string
	ServicePhoneNumber string
	Country            string
	State              string
	Province           string
}
