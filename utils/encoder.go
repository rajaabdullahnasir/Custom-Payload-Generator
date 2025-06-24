package utils

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

// EncodeURL encodes a payload using URL encoding
func EncodeURL(s string) string {
	return url.QueryEscape(s)
}

// EncodeBase64 encodes a payload to Base64
func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// EncodeHex converts each character to hex representation (e.g., \x41)
func EncodeHex(s string) string {
	var result string
	for _, c := range s {
		result += fmt.Sprintf("\\x%X", c)
	}
	return result
}

// EncodeUnicode escapes each character to Unicode (e.g., \u0041)
func EncodeUnicode(s string) string {
	var result string
	for _, c := range s {
		result += fmt.Sprintf("\\u%04X", c)
	}
	return result
}
