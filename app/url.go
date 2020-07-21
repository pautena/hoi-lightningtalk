package app

import (
	"net/url"
	"strings"
)

// isValidUrl tests a string to determine if it is a well-structured url or not.
func IsValidUrl(str string) bool {
	u, err := url.Parse(strings.ReplaceAll(str," ",""))
	return err == nil && u.Scheme != "" && u.Host != ""
}
