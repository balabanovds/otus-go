package url_shortener

import (
	"testing"
)

func TestURLWithoutHttpPrefix(t *testing.T) {
	url := "otus.ru/link go"
	var u URL

	short := u.Shorten(url)
	long := u.Resolve(short)

	if long != url {
		t.Errorf("expected %v, got %v", url, long)
	}
}

func TestURLWithHttpPrefix(t *testing.T) {
	url := "https://otus.ru/link go"
	var u URL

	short := u.Shorten(url)
	long := u.Resolve(short)

	if long != url {
		t.Errorf("expected %v, got %v", url, long)
	}
}

func TestNotFoundURL(t *testing.T) {
	var u URL

	long := u.Resolve("some.ru/unknown")

	if long != "" {
		t.Error("expected empty string")
	}
}
