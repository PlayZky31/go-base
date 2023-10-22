package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
)

type ValidationEngine struct {
	Validator    *validator.Validate
	ENTranslator ut.Translator
}

func NewEngine() (*ValidationEngine, error) {
	english := en.New()
	indonesian := id.New()
	v := &ValidationEngine{}
	uni := ut.New(english, english, indonesian, indonesian)

	enTranslator, _ := uni.GetTranslator("en")

	validate := validator.New()

	// Register English translations and customize error messages
	if err := enTranslation.RegisterDefaultTranslations(validate, enTranslator); err != nil {
		return nil, err
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	v.Validator = validate
	v.ENTranslator = enTranslator

	_ = validate.RegisterValidation("customDate", validateCustomDate)
	v.addTranslation("customDate", ErrInvalidDateMSG)

	return v, nil
}
