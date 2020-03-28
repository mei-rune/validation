package revel_ext

import (
	"fmt"

	"github.com/revel/revel"
	"github.com/runner-mei/validation"
)

func ToRevelErrors(errList []validation.ValidationError) []revel.ValidationError {
	results := make(revel.ValidationError, len(errList))
	for idx := range errList {
		results[idx] = errList[idx]
	}
	return results
}

func ToValidation(v *revel.Validation) *validation.Validation {
	return &validation.Validation{
		Locale:     v.Request.Locale,
		Translator: v.Translator,
	}
}

func Validate(value interface{}, v *revel.Validation) bool {
	v1, ok := value.(interface {
		Validate(*validation.Validation) bool
	})
	if ok {
		iv := ToValidation(v)
		if v1.Validate(iv) {
			for idx := range iv.Errors {
				v.Errors = append(v.Errors, iv.Errors[idx])
			}
			return true
		}
		return false
	}
	v2, ok := value.(func(*validation.Validation) bool)
	if ok {
		iv := ToValidation(v)
		if v2(iv) {
			for idx := range iv.Errors {
				v.Errors = append(v.Errors, iv.Errors[idx])
			}
			return true
		}
		return false
	}

	v3, ok := value.(interface {
		Validate(*revel.Validation) bool
	})
	if ok {
		return v3.Validate(v)
	}
	v4, ok := value.(func(*revel.Validation) bool)
	if ok {
		return v4(v)
	}
	v.Errors = append(v.Errors, revel.ValidationError{Message: fmt.Sprintf("revel_ext.Validate: %T", value)})
	return validation.HasErrors()
}
