package util

import (
	"github.com/go-openapi/swag"
)

// IsEmpty returns the value has a valid value or not
func IsEmpty(candidate *string) bool {
	return candidate == nil || swag.StringValue(candidate) == ""
}
