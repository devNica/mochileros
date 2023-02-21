package request

type UserAccounRequestModel struct {
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

type KYCRequestModel struct {
	UserId    string `json:"userId" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}
