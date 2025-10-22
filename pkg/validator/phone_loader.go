package validator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// In-memory cache of disposable numbers
var DisposablePhones = map[string]struct{}{}

// URL for public disposable phone numbers
const DisposablePhoneURL = "https://raw.githubusercontent.com/iP1SMS/disposable-phone-numbers/refs/heads/master/number-list.json"

// LoadDisposablePhones fetches and caches numbers
func LoadDisposablePhones() error {
	resp, err := http.Get(DisposablePhoneURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var numbersMap map[string]string
	if err := json.Unmarshal(body, &numbersMap); err != nil {
		return err
	}

	for num := range numbersMap {
		DisposablePhones[strings.TrimSpace(num)] = struct{}{}
	}

	fmt.Printf("✅ Loaded %d disposable phone numbers\n", len(DisposablePhones))
	return nil
}

// CheckDisposablePhone returns true if the number exists in the cache
func CheckDisposablePhone(number string) bool {
	_, exists := DisposablePhones[strings.TrimSpace(number)]
	return exists
}
