package validation

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func getJsonFieldName(fieldName string, obj interface{}) (tag string) {
	structType := reflect.TypeOf(obj)
	tag = "unknown"

	field, ok := structType.FieldByName(fieldName)
	if !ok {
		return
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag != "" {
		jsonTagParts := strings.Split(jsonTag, ",")
		tag = jsonTagParts[0]
	}
	return
}

func validateCardNumber(fl validator.FieldLevel) bool {
	cardNumber := fl.Field().String()

	re := regexp.MustCompile(`^[A-Za-z]\d{10}$`)
	return re.MatchString(cardNumber)
}

func validateNotEmpty(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

func BuyerValidator() *validator.Validate {
	validate := validator.New()
	if err := validate.RegisterValidation("cardnumber", validateCardNumber); err != nil {
		panic(fmt.Sprint("Error registering cardnumber validation:", err.Error()))
	}

	if err := validate.RegisterValidation("str_not_empty", validateNotEmpty); err != nil {
		panic(fmt.Sprint("Error registering str_not_empty validation:", err.Error()))
	}

	return validate
}

func MapValidatorErrors(err error, structToValidate interface{}) (messages map[string]string) {
	messages = make(map[string]string)
	for _, errItem := range err.(validator.ValidationErrors) {
		tag := errItem.Tag()
		field := errItem.StructField()
		jsonTag := getJsonFieldName(field, structToValidate)
		switch tag {
		case "cardnumber":
			messages[jsonTag] = "Must start with a letter and be followed by 10 digits"
		case "str_not_empty":
			messages[jsonTag] = "Field cannot be empty"
		case "required":
			messages[jsonTag] = "Field is required"
		}
	}

	return
}
