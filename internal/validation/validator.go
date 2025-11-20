package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	//Validate.RegisterValidation("myCustom", customValidation)

}

func ValidatingOrder(s interface{}) []string {
	if err := Validate.Struct(s); err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(
				errors,
				fmt.Sprintf(
					"Field '%s' failed validation '%s', value: '%s'",
					e.Field(), e.Tag(), e.Value(),
				),
			)
		}
		return errors
	}
	return nil
}
