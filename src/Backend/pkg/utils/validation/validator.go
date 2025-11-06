// validation/validator.go
package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Turgho/Aluguei/internal/errors"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	// Registrar função personalizada para validar UUID
	validate.RegisterValidation("uuid4", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		if len(value) != 36 {
			return false
		}
		return strings.Count(value, "-") == 4
	})
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(s interface{}) *errors.AppError {
	err := validate.Struct(s)
	if err != nil {
		var fieldErrors []FieldError

		for _, err := range err.(validator.ValidationErrors) {
			field := getJSONFieldName(s, err.StructField())
			message := getValidationMessage(err)

			fieldErrors = append(fieldErrors, FieldError{
				Field:   field,
				Message: message,
			})
		}

		return errors.NewValidationError("Dados de entrada inválidos", fieldErrors)
	}

	return nil
}

func getJSONFieldName(s interface{}, fieldName string) string {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	field, found := t.FieldByName(fieldName)
	if !found {
		return fieldName
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" || jsonTag == "-" {
		return fieldName
	}

	// Remover opções como omitempty
	parts := strings.Split(jsonTag, ",")
	return parts[0]
}

func getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "Este campo é obrigatório"
	case "email":
		return "Email inválido"
	case "min":
		return fmt.Sprintf("Deve ter no mínimo %s caracteres", err.Param())
	case "max":
		return fmt.Sprintf("Deve ter no máximo %s caracteres", err.Param())
	case "len":
		return fmt.Sprintf("Deve ter exatamente %s caracteres", err.Param())
	case "uuid4":
		return "UUID inválido"
	case "oneof":
		return fmt.Sprintf("Deve ser um dos: %s", strings.ReplaceAll(err.Param(), " ", ", "))
	case "gtfield":
		return fmt.Sprintf("Deve ser maior que %s", err.Param())
	default:
		return fmt.Sprintf("Campo inválido: %s", err.Tag())
	}
}
