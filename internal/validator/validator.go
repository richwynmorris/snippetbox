package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Validator struct captures the field errors that may be encountered when a user submits a form.
type Validator struct {
	NonFieldErrors []string
	FieldErrors    map[string]string
}

// Valid returns whether any FieldErrors have been encountered.
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
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

// NotBlank checks if the given string is empty.
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars checks whether the given string is less than or equal to the provided argument.
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedInt checks if the value argument is containted within the permitted values collection.
func PermittedInt(val int, permittedValues ...int) bool {
	for i := range permittedValues {
		if val == permittedValues[i] {
			return true
		}
	}
	return false
}

// MinChars returns a boolean based on whether the given string argument is greater than the passed in integer.
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

// Matches returns a boolean based on whether the string argument matches the regex pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// addNonFieldError appends a new message to the NonFieldError field on the Validator struct.
func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}
