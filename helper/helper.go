package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func ResponseApi(message string, code int, status string, data interface{}) Response {
	response := Response{
		Message: message,
		Code:    code,
		Status:  status,
		Data:    data,
	}
	return response
}

// func FormatError(err error) []string {
// 	var errors []string
// 	for _, e := range err.(validator.ValidationErrors) {
// 		errors = append(errors, e.Error())
// 	}
// 	return errors
// }
func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
