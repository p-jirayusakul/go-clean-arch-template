package middleware

import (
	"database/sql/driver"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
)

// CustomValidator holds the validator instance.
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator creates a new instance of CustomValidator.
func NewCustomValidator() *CustomValidator {
	v := validator.New()

	// register all pgtype* types to use the ValidateValuer CustomTypeFunc
	v.RegisterCustomTypeFunc(ValidateValuer, pgtype.Text{})
	return &CustomValidator{validator: v}
}

// Validate validates the given struct using the validator instance.
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}

// ValidateValuer implements validator.CustomTypeFunc
func ValidateValuer(field reflect.Value) interface{} {

	if valuer, ok := field.Interface().(driver.Valuer); ok {

		val, err := valuer.Value()
		if err == nil {
			return val
		}
		// handle the error how you want
	}

	return nil
}
