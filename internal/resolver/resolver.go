package resolver

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/sapirrior/blink/internal/validate"
)

// Resolve resolves raw command line arguments to a valid URL or file URI.
func Resolve(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("no input provided")
	}

	firstArg := args[0]

	// 1. Check for Alias
	if alias, ok := FindAlias(firstArg); ok {
		query := strings.Join(args[1:], " ")
		return BuildAliasURL(alias, query), nil
	}

	if len(args) > 1 {
		return "", fmt.Errorf("too many arguments: did you mean to use an alias?")
	}

	// 2. Check for File Path
	isPathPattern := strings.HasPrefix(firstArg, "./") ||
		strings.HasPrefix(firstArg, "../") ||
		strings.HasPrefix(firstArg, "/") ||
		strings.HasPrefix(firstArg, "~/")

	resolvedPath := firstArg
	if strings.HasPrefix(firstArg, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			resolvedPath = filepath.Join(home, firstArg[2:])
		}
	}

	// Perform Stat check to see if the file exists
	_, statErr := os.Stat(resolvedPath)
	exists := statErr == nil

	if exists || isPathPattern {
		if !exists {
			return "", fmt.Errorf("file not found: %s", firstArg)
		}
		absPath, err := filepath.Abs(resolvedPath)
		if err != nil {
			return "", fmt.Errorf("failed to get absolute path: %w", err)
		}
		u := url.URL{
			Scheme: "file",
			Path:   filepath.ToSlash(absPath),
		}
		return u.String(), nil
	}

	// 3. Treat as URL
	targetURL := firstArg
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "https://" + targetURL
	}

	if err := validate.IsValidURL(targetURL); err != nil {
		return "", err
	}

	return targetURL, nil
}
