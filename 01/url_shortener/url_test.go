package url_shortener

import (
	"fmt"
	"testing"
)

func TestURLWithoutHttpPrefix(t *testing.T) {
	url := "otus.ru/link"
	var u URL

	short := u.Shorten(url)

	fmt.Println(short)
}
