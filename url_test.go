package url

import (
	"testing"
)

func TestURLWithoutHttpPrefix(t *testing.T) {
	url := "otus.ru/link go"
	u := URLShortener{}

	short := u.Shorten(url)
	long := u.Resolve(short)

	if long != url {
		t.Errorf("expected %v, got %v", url, long)
	}
}

func TestURLWithHttpPrefix(t *testing.T) {
	url := "https://otus.ru/link go"
	u := URLShortener{}

	short := u.Shorten(url)
	long := u.Resolve(short)

	if long != url {
		t.Errorf("expected %v, got %v", url, long)
	}
}

// this test should panic as URLShortener.Data did not initialized
func TestNotFoundURL(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("TestNotFoundURL should panic")
		}
	}()

	u := URLShortener{}

	long := u.Resolve("some.ru/unknown")

	if long != "" {
		t.Error("expected empty string")
	}
}
