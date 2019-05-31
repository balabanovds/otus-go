package url

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

// IShortener interface
type IShortener interface {
	Shorten(string) string
	Resolve(string) string
}

// URLShortener main data struct
type URLShortener struct {
	Data map[string]string
}

// Shorten url, and stores it in struct
func (s *URLShortener) Shorten(url string) string {
	if s.Data == nil {
		s.Data = make(map[string]string)
	}
	u, err := parse(url)
	if err != nil {
		log.Fatal(err)
	}

	path := hash(u.EscapedPath())
	hostname := u.Hostname()

	shortURL := fmt.Sprintf("%s/%s", hostname, path)

	s.Data[shortURL] = url

	return shortURL
}

// Resolve url from short to normal or empty if not found
func (s *URLShortener) Resolve(url string) string {
	if s.Data == nil {
		panic("Storage is not initialized yet")
	}
	if res, ok := s.Data[url]; ok {
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
