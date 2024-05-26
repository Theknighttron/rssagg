package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API KEY from
// the Headers of an HTTP request
// Example
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
    key := headers.Get("Authorization")
    if key == "" {
        return "", errors.New("no authentication info found")
    }

    apiKey := strings.Split(key, " ")
    if len(apiKey) != 2 {
        return "", errors.New("malformed authentication header")
    }

    if apiKey[0] != "ApiKey" {
        return "", errors.New("malformed first part of the authentication header")
    }

    return apiKey[1], nil
}
