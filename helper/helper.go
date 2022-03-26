package helper

import (
	"triadmoko-be-golang/translation/id"

	indonesia "github.com/go-playground/locales/id"
	uni "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

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

func FormatValidationError(structs interface{}) []string {
	validate := validator.New()
	idn := indonesia.New()
	uni := uni.New(idn, idn)
	trans, _ := uni.GetTranslator("id")
	_ = id.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(structs)
	errs := translateError(err, trans)

	return errs
}

func translateError(err error, trans uni.Translator) (errs []string) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := e.Translate(trans)
		errs = append(errs, translatedErr)
	}
	return errs
}
