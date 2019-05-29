package url_shortener

import (
	"log"
	"net/url"
)

// Shortener interface
type Shortener interface {
	Shorten(string) string
	Resolve(string) string
}

// URL is representation of short or long url
type URL string

var data map[string]string

// Shorten url, and stores it in map
func (u *URL) Shorten(s string) string {
	url, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	path := url.EscapedPath()

	return "tets"
}

// Resolve url from short to normal
func (u *URL) Resolve(s string) string {
	return ""
}
