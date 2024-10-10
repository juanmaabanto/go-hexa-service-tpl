package valueobject

import (
	"errors"
	"net/url"
)

type URL string

func NewURL(value string) (URL, error) {
	parsedURL, err := url.Parse(value)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", errors.New("format error")
	}

	return URL(value), nil
}

func (u URL) String() string {
	return string(u)
}
