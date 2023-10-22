package rest_api_common

import (
	"fmt"
	customError2 "gandiwa/pkg/customError"
	validator3 "gandiwa/pkg/validator"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate(validatorEngine *validator3.ValidationEngine, payload interface{}) error {
	if err := validatorEngine.Validator.Struct(payload); err != nil {
		validatorErrs, ok := err.(validator.ValidationErrors)
		if !ok || len(validatorErrs) == 0 {
			return customError2.ErrGeneral(fmt.Errorf("something went wrong in validation engine"))
		}

		errs := make(map[string]string, len(validatorErrs))
		for _, e := range validatorErrs {
			fieldName := strings.ToLower(e.Field())
			errMsg := e.Error()
			errMsg = e.Translate(validatorEngine.ENTranslator)
			errs[fieldName] = errMsg
		}

		return customError2.ErrBadRequestFields(errs)
	}

	return nil
}
