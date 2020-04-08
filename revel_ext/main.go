package revel_ext

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/revel/revel"
	"github.com/runner-mei/validation"
)

func TryToRevelValidateErrors(v *revel.Validation, err error, prefix string) bool {
	ok, errList := validation.ToValidationErrors(err)
	if !ok {
		return ok
	}
	for idx := range errList {
		v.Errors = append(v.Errors, &revel.ValidationError{
			Key:     prefix + errList[idx].Key,
			Message: errList[idx].Message,
		})
	}
	return true
}

func ToRevelErrors(errList []validation.ValidationError, prefix string) []revel.ValidationError {
	results := make([]revel.ValidationError, len(errList))
	for idx := range errList {
		results[idx].Key = prefix + errList[idx].Key
		results[idx].Message = errList[idx].Message
	}
	return results
}

func ToValidation(v *revel.Validation) *validation.Validation {
	return &validation.Validation{
		Locale:     v.Request.Locale,
		Translator: v.Translator,
	}
}

func Validate(value interface{}, v *revel.Validation, prefix string) bool {
	v1, ok := value.(interface {
		Validate(*validation.Validation) bool
	})
	if ok {
		iv := ToValidation(v)
		if v1.Validate(iv) {
			for idx := range iv.Errors {
				v.Errors = append(v.Errors, &revel.ValidationError{
					Key:     prefix + iv.Errors[idx].Key,
					Message: iv.Errors[idx].Message,
				})
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
				v.Errors = append(v.Errors, &revel.ValidationError{
					Key:     prefix + iv.Errors[idx].Key,
					Message: iv.Errors[idx].Message,
				})
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

	rValue := reflect.ValueOf(value)
	if rValue.CanAddr() {
		return Validate(rValue.Addr().Interface(), v, prefix)
	} else if rValue.Kind() != reflect.Ptr {
		panic(errors.New("请试试用指针不要用值对象"))
	}

	v.Errors = append(v.Errors, &revel.ValidationError{Message: fmt.Sprintf("revel_ext.Validate: %T", value)})
	return v.HasErrors()
}
