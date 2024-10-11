// Copyright (c) 2012-2016 The Revel Framework Authors, All rights reserved.
// Revel Framework source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package validation

import (
	boovalidation "github.com/boo-admin/boo/validation"
)

// var (
// 	En   ut.Translator
// 	Zh   ut.Translator
// 	UtZh *ut.UniversalTranslator
// )

type StructValidator = boovalidation.StructValidator

var DefaultStructValidator = boovalidation.DefaultStructValidator

type ValidatableError = boovalidation.ValidatableError

func ToValidationErrors(err error) (bool, []ValidationError) {
	return boovalidation.ToValidationErrors(err)
}

type ValidationErrors = boovalidation.ValidationErrors

// ValidationError simple struct to store the Message & Key of a validation error
type ValidationError = boovalidation.ValidationError

func NewValidationError(field, message string) error {
	return boovalidation.NewValidationError(field, message)
}

type TranslateFunc = boovalidation.TranslateFunc

var Default = boovalidation.Default

func New(locale string, translator TranslateFunc) *Validation {
	return boovalidation.New(locale, translator)
}

// Validation context manages data validation and error messages.
type Validation = boovalidation.Validation

// ValidationResult is returned from every validation method.
// It provides an indication of success, and a pointer to the Error (if any).
type ValidationResult = boovalidation.ValidationResult
