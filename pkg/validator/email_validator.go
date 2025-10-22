package validator

import (
	"net"
	"net/mail"
	"strings"
)

// ValidateEmail runs format + disposable + MX checks
func ValidateEmail(email string) (bool, string) {
	// 1. Format validation
	if _, err := mail.ParseAddress(email); err != nil {
		return false, "invalid_format"
	}

	// 2. Extract domain
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false, "invalid_format"
	}
	domain := strings.ToLower(strings.TrimSpace(parts[1]))

	// 3. Check disposable
	if CheckDisposable(domain) {
		return false, "disposable_domain"
	}

	// 4. MX record
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return false, "no_mx_record"
	}

	return true, "valid"
}
