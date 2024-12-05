package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string]string {

	var customMessages = map[string]string{
		"required": "The %s field is required",
		"email":    "The %s field must be a valid email address",
		"min":      "The %s field must be at least %s characters long",
		"gte":      "The %s field must be greater than or equal to %s",
		"lte":      "The %s field must be less than or equal to %s",
	}

	errors := make(map[string]string)

	if _, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			tag := fieldErr.Tag()
			field := fieldErr.Field()
			param := fieldErr.Param()

			if message, exists := customMessages[tag]; exists {
				// Ensure all placeholders are filled
				if param != "" {
					errors[field] = fmt.Sprintf(message, field, param)
				} else {
					errors[field] = fmt.Sprintf(message, field)
				}
			} else {
				// Fallback for undefined rules
				errors[field] = fmt.Sprintf("%s failed on the '%s' validation rule", field, tag)
			}
		}
	} else {
		errors["error"] = "Unknown error occurred"
	}

	return errors
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// Match URL with dynamic parts (e.g., :id)
func MatchURL(pattern, url string) bool {
	// Split the pattern and the actual URL into segments
	patternParts := strings.Split(strings.TrimSuffix(pattern, "/"), "/")
	urlParts := strings.Split(strings.TrimSuffix(url, "/"), "/")

	// Check if the number of segments matches
	if len(patternParts) != len(urlParts) {
		return false
	}

	// Compare each segment
	for i := 0; i < len(patternParts); i++ {
		if strings.HasPrefix(patternParts[i], ":") {
			// Dynamic segment matches any value
			continue
		}
		if patternParts[i] != urlParts[i] {
			// Segment mismatch
			return false
		}
	}

	// All segments matched
	return true
}
