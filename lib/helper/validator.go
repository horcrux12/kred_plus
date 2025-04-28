package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"kredi-plus.com/be/lib/exception"
	"net/http"
	"regexp"
	"strings"
)

func Validate(data interface{}) (err error) {
	validate := validator.New()

	validate = RegisterCustomValidation(validate)

	err = validate.Struct(data)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			// just return the first error
			return exception.CustomError{
				Message:  TagMessage(e),
				HttpCode: http.StatusBadRequest,
			}
		}
	}

	return
}

func TagMessage(e validator.FieldError) string {
	messages := map[string]string{
		"required":             "%s can't be empty",
		"required_if":          "%s can't be empty",
		"required_with":        "%s can't be empty",
		"required_without":     "%s can't be empty",
		"required_without_all": "%s can't be empty",
		"email":                "Please enter your email address in format: yourname@example.com",
		"gt":                   "%s must be greater than %s",
		"nik":                  "NIK should be 16 digits number",
		"phone":                "%s incorrect",
		"eqfield":              "%s mismatch",
		"nefield":              "%s cannot match with %s",
		"oneof":                "%s must be one of %s",
		"eq":                   "%s must be equal to %s",
		"number":               "%s must be number values only",
	}

	if CheckDataOnSlice(e.Tag(), []string{"nik", "email"}) {
		return messages[e.Tag()]
	}

	if msg, ok := messages[e.Tag()]; ok {
		return formatMessage(msg, e)
	}

	return e.Error()
}

func formatMessage(msg string, e validator.FieldError) string {
	switch e.Tag() {
	case "nefield", "eq":
		return fmt.Sprintf(msg, CapitalizedEachWords(e.Field()), e.Param())
	case "gt":
		param := strings.Split(e.Param(), ";")
		return fmt.Sprintf(msg, CapitalizedEachWords(e.Field()), param[0])
	case "oneof":
		param := strings.ReplaceAll(e.Param(), " ", ", ")
		return fmt.Sprintf(msg, CapitalizedEachWords(e.Field()), param)
	case "nik":
		return msg
	default:
		return fmt.Sprintf(msg, e.Field())
	}
}

func RegisterCustomValidation(v *validator.Validate) *validator.Validate {
	_ = v.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		regexPattern := `^[0-9A-Za-z]{6,16}$`
		matched, err := regexp.MatchString(regexPattern, fl.Field().String())
		if err != nil {
			return false
		}
		return matched
	})

	// nik
	_ = v.RegisterValidation("nik", func(fl validator.FieldLevel) bool {
		regexPattern := `^[0-9]{16}$`
		matched, err := regexp.MatchString(regexPattern, fl.Field().String())
		if err != nil {
			return false
		}
		return matched
	})
	return v
}
