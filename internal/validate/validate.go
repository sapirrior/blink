package validate

import (
	"fmt"
	"net/url"
)

// IsValidURL checks if a URL is strictly formatted and has a valid scheme (http, https, file).
func IsValidURL(raw string) error {
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	if u.Scheme == "" {
		return fmt.Errorf("invalid URL: must include scheme")
	}
	if u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "file" {
		return fmt.Errorf("unsupported scheme %q: use http, https, or a file path", u.Scheme)
	}
	if u.Scheme != "file" && u.Host == "" {
		return fmt.Errorf("invalid URL: remote URLs must include host")
	}
	return nil
}
