package resolver

import (
	"net/url"
	"strings"
)

// Alias defines a site shortcut and search engine mapping.
type Alias struct {
	Keys      []string
	BaseURL   string
	SearchURL string
}

// BuiltInAliases is the default set of shortcuts.
var BuiltInAliases = []Alias{
	{
		Keys:      []string{"g", "google"},
		BaseURL:   "https://google.com",
		SearchURL: "https://google.com/search?q=",
	},
	{
		Keys:      []string{"yt", "youtube"},
		BaseURL:   "https://youtube.com",
		SearchURL: "https://youtube.com/results?search_query=",
	},
	{
		Keys:      []string{"gh", "github"},
		BaseURL:   "https://github.com",
		SearchURL: "https://github.com/search?q=",
	},
	{
		Keys:      []string{"gist"},
		BaseURL:   "https://gist.github.com",
		SearchURL: "https://gist.github.com/search?q=",
	},
	{
		Keys:      []string{"reddit", "r"},
		BaseURL:   "https://reddit.com",
		SearchURL: "https://reddit.com/search?q=",
	},
	{
		Keys:      []string{"wiki", "w", "wikipedia"},
		BaseURL:   "https://en.wikipedia.org",
		SearchURL: "https://en.wikipedia.org/wiki/Special:Search?search=",
	},
	{
		Keys:      []string{"dict", "wiktionary"},
		BaseURL:   "https://en.wiktionary.org",
		SearchURL: "https://en.wiktionary.org/wiki/Special:Search?search=",
	},
	{
		Keys:      []string{"npm"},
		BaseURL:   "https://npmjs.com",
		SearchURL: "https://npmjs.com/search?q=",
	},
	{
		Keys:      []string{"pkg", "godoc"},
		BaseURL:   "https://pkg.go.dev",
		SearchURL: "https://pkg.go.dev/search?q=",
	},
	{
		Keys:      []string{"mdn"},
		BaseURL:   "https://developer.mozilla.org",
		SearchURL: "https://developer.mozilla.org/en-US/search?q=",
	},
	{
		Keys:      []string{"cargo", "crates"},
		BaseURL:   "https://crates.io",
		SearchURL: "https://crates.io/search?q=",
	},
	{
		Keys:      []string{"docker", "hub"},
		BaseURL:   "https://hub.docker.com",
		SearchURL: "https://hub.docker.com/search?q=",
	},
	{
		Keys:      []string{"aw", "archwiki"},
		BaseURL:   "https://wiki.archlinux.org",
		SearchURL: "https://wiki.archlinux.org/index.php?search=",
	},
	{
		Keys:      []string{"tw", "twitter", "x"},
		BaseURL:   "https://x.com",
		SearchURL: "https://x.com/search?q=",
	},
	{
		Keys:      []string{"maps"},
		BaseURL:   "https://maps.google.com",
		SearchURL: "https://maps.google.com/?q=",
	},
	{
		Keys:      []string{"translate"},
		BaseURL:   "https://translate.google.com",
		SearchURL: "https://translate.google.com/?text=",
	},
	{
		Keys:      []string{"ddg", "duckduckgo"},
		BaseURL:   "https://duckduckgo.com",
		SearchURL: "https://duckduckgo.com/?q=",
	},
	{
		Keys:      []string{"so", "stackoverflow"},
		BaseURL:   "https://stackoverflow.com",
		SearchURL: "https://stackoverflow.com/search?q=",
	},
	{
		Keys:      []string{"pypi"},
		BaseURL:   "https://pypi.org",
		SearchURL: "https://pypi.org/search/?q=",
	},
	{
		Keys:      []string{"mvn"},
		BaseURL:   "https://search.maven.org",
		SearchURL: "https://search.maven.org/search?q=",
	},
}

// FindAlias resolves a key to a specific Alias structure, if matching.
func FindAlias(key string) (Alias, bool) {
	lowerKey := strings.ToLower(key)
	for _, a := range BuiltInAliases {
		for _, k := range a.Keys {
			if k == lowerKey {
				return a, true
			}
		}
	}
	return Alias{}, false
}

// BuildAliasURL generates the final URL for an alias and its search query.
func BuildAliasURL(a Alias, query string) string {
	if query == "" {
		return a.BaseURL
	}
	return a.SearchURL + url.QueryEscape(query)
}
