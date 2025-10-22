package validator

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

// Global in-memory cache
var DisposableDomains = map[string]struct{}{}

// URLs of public disposable domain lists
var DisposableURLs = []string{
	"https://raw.githubusercontent.com/7c/fakefilter/refs/heads/main/txt/data.txt",
	"https://raw.githubusercontent.com/disposable-email-domains/disposable-email-domains/refs/heads/main/disposable_email_blocklist.conf",
	"https://raw.githubusercontent.com/wesbos/burner-email-providers/refs/heads/master/emails.txt",
}

// LoadDisposableDomains fetches all domains from the URLs and caches them in memory
func LoadDisposableDomains() error {
	for _, url := range DisposableURLs {
		if err := fetchAndCache(url); err != nil {
			fmt.Printf("⚠️ Failed to fetch %s: %v\n", url, err)
		}
	}
	fmt.Printf("✅ Loaded %d disposable domains\n", len(DisposableDomains))
	return nil
}

// fetchAndCache fetches a URL and adds all domains to the global map
func fetchAndCache(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		DisposableDomains[strings.ToLower(line)] = struct{}{}
	}

	return scanner.Err()
}

// CheckDisposable checks if a domain exists in the cached map
func CheckDisposable(domain string) bool {
	_, exists := DisposableDomains[strings.ToLower(domain)]
	return exists
}
