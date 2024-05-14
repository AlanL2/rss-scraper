package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from the headers of an http request
// Example: Authorization: ApiKey {insert apikey here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("error: no authentication info found")
	}
	
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("error: malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("error: malformed first part of auth")
	}

	return vals[1], nil
}