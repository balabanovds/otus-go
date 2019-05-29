package otus_go

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

// Shortener interface
type Shortener interface {
	Shorten(string) string
	Resolve(string) string
}

// URL is representation of short or long url
type URL string

var data map[string]string

func init() {
	data = make(map[string]string)
}

// Shorten url, and stores it in map
func (*URL) Shorten(s string) string {
	u, err := parse(s)
	if err != nil {
		log.Fatal(err)
	}

	path := hash(u.EscapedPath())
	hostname := u.Hostname()

	shortURL := fmt.Sprintf("%s/%s", hostname, path)

	data[shortURL] = s

	return shortURL
}

// Resolve url from short to normal or empty if not found
func (*URL) Resolve(s string) string {
	if res, ok := data[s]; ok {
		return res
	}
	return ""
}

func hash(str string) string {
	hash := getMD5(str)
	var sum uint64
	for _, b := range []byte(hash) {
		sum += uint64(b)
	}
	return strconv.FormatUint(sum+1e5, 16)
}

func parse(s string) (*url.URL, error) {
	addHttp(&s)
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func addHttp(str *string) {
	if !strings.Contains(*str, "http") {
		*str = "http://" + *str
	}
}

func getMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
