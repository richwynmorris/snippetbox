package validator

import (
	"strings"
	"unicode/utf8"
)

// Validator struct captures the field errors that may be encountered when a user submits a form.
type Validator struct {
	FieldErrors map[string]string
}

// Valid returns whether any FieldErrors have been encountered.
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

// AddFieldError adds the error to the FieldErrors map if not already previously encountered.
func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// CheckField uses the return value from exported functions to determine is an error should be
// captured.
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// Exported functions return boolean values to validate whether the error should be captured:

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func PermittedInt(val int, permittedValues ...int) bool {
	for i := range permittedValues {
		if val == i {
			return true
		}
	}
	return false
}
