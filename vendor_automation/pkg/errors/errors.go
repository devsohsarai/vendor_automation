package errors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gowaves/vendor_automation/internal/providers/validation"
)

var errorList = make(map[string]string)

func Init() {
	errorList = map[string]string{}
}

func SetFromErrors(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			Add(fieldError.Field(), GetErrorMsg(fieldError.Tag()))
		}
	}
}

func Add(key string, value string) {
	errorList[strings.ToLower(key)] = value
}

func GetErrorMsg(tag string) string {
	return validation.ErrorMessages()[tag]
}

func Get() map[string]string {
	return errorList
}
