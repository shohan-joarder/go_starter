package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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

// ParseUserID converts the userID to a uint and validates it
func ParseUserID(ctx *gin.Context) (uint, error) {

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID not found in context",
		})

		return 0, errors.New("user ID not found in context")
	}

	switch v := userID.(type) {
	case float64:
		// Validate that the float64 is a whole number
		if v != float64(int(v)) {
			return 0, errors.New("user ID must be a whole number")
		}
		return uint(v), nil
	case int:
		if v < 0 {
			return 0, errors.New("user ID must be a positive number")
		}
		return uint(v), nil
	case string:
		parsedID, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return 0, errors.New("user ID should be a valid numeric string")
		}
		return uint(parsedID), nil
	default:
		return 0, fmt.Errorf("unsupported user ID type: %T", v)
	}
}

// UniqueValidator checks if a field value is unique in the database.
func UniqueValidator(db *gorm.DB, isUpdate bool) validator.Func {
	return func(fl validator.FieldLevel) bool {
		// Get the param from the validation tag (e.g., "users_user_name")
		param := fl.Param()
		fmt.Println("Params", param)
		// Split the param by underscore to extract table name and column name
		parts := strings.Split(param, "_")
		if len(parts) != 2 {
			// Invalid format for the param; should be "table_column"
			fmt.Println("Invalid param format. Expected 'table_column'. Got:", param)
			return false
		}

		tableName := parts[0]  // Extract table name, e.g., "users"
		columnName := parts[1] // Extract column name, e.g., "user_name"

		// Get the field value being validated
		value := fl.Field().Interface()

		// Prepare a query to check for uniqueness
		query := db.Table(tableName).Where(columnName+" = ?", value)

		fmt.Println("query", query)

		// Handle update-specific logic (if `isUpdate` is true)
		if isUpdate {
			// Retrieve the current record ID to exclude it from the uniqueness check
			idField := fl.Parent().FieldByName("ID")
			if idField.IsValid() && idField.Interface() != nil {
				id := idField.Interface()
				query = query.Where("id != ?", id)
			}
		}

		// Execute the query and check the count
		var count int64
		query.Count(&count)

		// Return true if the count is 0 (unique), false otherwise
		return count == 0
	}
}

func PhoneValidator(fl validator.FieldLevel) bool {
	// Define a regex pattern for valid phone numbers
	phoneRegex := `^\+?[0-9]{10,15}$`

	// Compile the regex
	re := regexp.MustCompile(phoneRegex)

	// Validate the field value
	phone := fl.Field().String()
	return re.MatchString(phone)
}
