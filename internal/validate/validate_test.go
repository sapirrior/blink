package validate

import (
	"testing"
)

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name    string
		rawURL  string
		wantErr bool
	}{
		{"valid http", "http://google.com", false},
		{"valid https", "https://google.com", false},
		{"valid file", "file:///home/user/test.txt", false},
		{"empty string", "", true},
		{"missing scheme", "google.com", true},
		{"invalid scheme", "ftp://google.com", true},
		{"missing host for remote", "https://", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsValidURL(tt.rawURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidURL(%q) error = %v, wantErr %v", tt.rawURL, err, tt.wantErr)
			}
		})
	}
}
