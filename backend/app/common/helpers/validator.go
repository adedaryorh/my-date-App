package helpers

import (
	"fmt"
	"mime/multipart"
	"regexp"
	"strings"
	"time"
	"unicode"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var AllowedFiles = map[string]bool{
	"jpeg": true,
	"pdf":  true,
	"jpg":  true,
	"docx": true,
	"png":  true,
	"csv":  true,
	"xlsx": true,
}

type Enum interface {
	IsValid() bool
}

func ValidateInput(input interface{}) error {
	var errors []string
	v := validator.New()
	v.RegisterValidation("is_enum", ValidateEnum)
	v.RegisterValidation("is_amount", ValidateAmount)
	v.RegisterValidation("is_phone", ValidatePhoneNumber)
	v.RegisterValidation("is_uuid", ValidateUUID)
	v.RegisterValidation("is_password", ValidatePassword)
	v.RegisterValidation("is_date", ValidateDate)
	v.RegisterValidation("after_now", ValidateAfterNow)
	v.RegisterValidation("before_now", ValidateBeforeNow)
	v.RegisterValidation("is_base64_file", ValidateBase64File)
	v.RegisterValidation("is_file", ValidateFile)

	err := v.Struct(input)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			switch e.ActualTag() {
			case "required":
				errors = append(errors, fmt.Sprintf("%s field is required", e.Field()))
			case "required_if":
				errors = append(errors, fmt.Sprintf("%s field is required", e.Field()))
			case "email":
				errors = append(errors, fmt.Sprintf("%s must be a valid email", e.Field()))
			case "url":
				errors = append(errors, fmt.Sprintf("%s must be a valid url", e.Field()))
			case "gt":
				errors = append(errors, fmt.Sprintf("%s array cannot be empty", e.Field()))
			case "is_enum":
				errors = append(errors, fmt.Sprintf("%s is not a valid %v", e.Value(), e.Type()))
			case "is_amount":
				errors = append(errors, fmt.Sprintf("%s is not a valid amount", e.Value()))
			case "is_base64_file":
				errors = append(errors, fmt.Sprintf("%s is not a valid file", e.Value()))
			case "is_file":
				errors = append(errors, fmt.Sprintf("%s is not a valid file", e.Field()))
			case "is_uuid":
				errors = append(errors, fmt.Sprintf("%s is not a valid uuid", e.Value()))
			case "is_date":
				errors = append(errors, fmt.Sprintf("%s is not a valid date", e.Value()))
			case "after_now":
				errors = append(errors, fmt.Sprintf("%s cannot be in the past", e.Value()))
			case "before_now":
				errors = append(errors, fmt.Sprintf("%s cannot be in the future", e.Value()))
			case "is_phone":
				errors = append(errors, fmt.Sprintf("%s is not a valid phone number", e.Value()))
			case "is_password":
				errors = append(errors, fmt.Sprintf("%s is not a valid password", e.Value()))
			case "min":
				errors = append(errors, fmt.Sprintf("%s must be at least %s letters", e.Field(), e.Param()))
			case "max":
				errors = append(errors, fmt.Sprintf("%s cannot be more than %s letters", e.Field(), e.Param()))
			case "len":
				errors = append(errors, fmt.Sprintf("%s must be %s in length", e.Field(), e.Param()))
			case "numeric":
				errors = append(errors, fmt.Sprintf("%s must be a numeric string", e.Field()))
			default:
				errors = append(errors, "an error occurred")
			}
		}
	}
	var s string
	for _, errorString := range errors {
		s += fmt.Sprintf("%s \n", errorString)
	}
	if len(s) > 0 {
		return fmt.Errorf("%s", s)
	}
	return nil
}

func ValidateFile(field validator.FieldLevel) bool {
	value := field.Field().Interface().(multipart.FileHeader)
	split := strings.Split(value.Filename, ".")

	fileType := split[len(split)-1]

	_, ok := AllowedFiles[fileType]

	return ok
}

func ValidateBase64File(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)
	if !strings.HasPrefix(value, "data") {
		return false
	}
	splittedBase64 := strings.Split(value, ":")
	return len(splittedBase64) >= 2
}

func ValidateEnum(field validator.FieldLevel) bool {
	value := field.Field().Interface().(Enum)
	return value.IsValid()
}

func ValidateDate(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)
	layout := "2006-01-02"
	_, err := time.Parse(layout, value)
	return err == nil
}

func ValidateAfterNow(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)
	layout := "2006-01-02"
	t, _ := time.Parse(layout, value)
	return t.After(time.Now())
}

func ValidateBeforeNow(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)
	layout := "2006-01-02"
	t, _ := time.Parse(layout, value)
	return t.Before(time.Now())
}

func ValidateAmount(field validator.FieldLevel) bool {
	value := field.Field().Interface().(int64)
	return value > 0
}

func ValidateUUID(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)
	_, err := uuid.Parse(value)
	return err == nil
}

func ValidatePassword(field validator.FieldLevel) bool {
	var number, upper, special, lenghtOrMore bool
	password := field.Field().Interface().(string)
	letters := 0
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			letters++
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			letters++
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
		}
	}
	lenghtOrMore = letters >= 8
	if !number || !upper || !special || !lenghtOrMore {
		return false
	}
	return true
}

func ValidatePhoneNumber(field validator.FieldLevel) bool {
	value := field.Field().Interface().(string)
	// starts with +
	if !strings.HasPrefix(value, "+") {
		return false
	}
	// must be at least 11 characters
	if len(value) < 11 {
		return false
	}
	return regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`).MatchString(value)
}
