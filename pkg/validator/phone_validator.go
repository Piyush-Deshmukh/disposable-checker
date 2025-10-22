package validator

import (
	"regexp"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

// ValidatePhone parses, formats, and checks if number is disposable
func ValidatePhone(raw string) (bool, string) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return false, "empty_number"
	}

	// Parse phone number (no default region: infer from +)
	num, err := phonenumbers.Parse(raw, "")
	if err != nil {
		return false, "invalid_format"
	}

	// Check valid number
	if !phonenumbers.IsValidNumber(num) {
		return false, "invalid_number"
	}

	// Format to E.164
	formatted := phonenumbers.Format(num, phonenumbers.E164)

	// Check disposable
	if CheckDisposablePhone(normalizeNumber(formatted)) {
		return false, "disposable_number"
	}

	return true, "valid"
}

func normalizeNumber(num string) string {
	re := regexp.MustCompile(`\D`) // remove all non-digit chars
	return re.ReplaceAllString(num, "")
}
