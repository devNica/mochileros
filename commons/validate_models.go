package commons

import (
	"encoding/json"

	"github.com/devNica/mochileros/exceptions"
	"github.com/go-playground/validator/v10"
)

func ValidateModel(modelToValidate interface{}) {
	validate := validator.New()
	err := validate.Struct(modelToValidate)

	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": "this field is" + err.Tag(),
			})
		}

		jsonMessage, errJson := json.Marshal(messages)
		exceptions.PanicLogging(errJson)

		panic(exceptions.ValidationError{
			Message: string(jsonMessage),
		})
	}
}
