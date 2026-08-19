package commands


import(
	"fmt"
	"net/url"
	"path"
	"strings"
)

// Helper to parse character slugs for Fetch Command
func characterSlugParser(rawURL string) (string, error) {
	// Parse the input URL into a URL struct
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("parse character URL: %w", err)
	}

	// Trim the URL path and parse with the path package to extract the slug
	slug := strings.Trim(path.Base(parsed.Path), "/")
	if slug == "" || slug == "." {
		return "", fmt.Errorf("character URL has no slug: %s", rawURL)
	}
	
	return slug, nil
}
