package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func Email(bearer string) (*string, error) {
	raw := strings.Split(strings.TrimSpace(bearer), " ")
	if len(raw) != 2 {
		return nil, fmt.Errorf("the 'Bearer' or 'token' string is missing. Received: %q", bearer)
	}

	tokenParts := strings.Split(raw[1], ".")
	if len(tokenParts) != 3 {
		return nil, fmt.Errorf("invalid token value. Received: %q", raw[1])
	}

	var (
		payload []byte
		err     error
	)

	for _, encoding := range []*base64.Encoding{base64.RawStdEncoding, base64.StdEncoding} {
		payload, err = encoding.DecodeString(tokenParts[1])
		if err == nil {
			break
		}
	}

	if err != nil {
		return nil, fmt.Errorf("an error occurred decoding the payload: %w", err)
	}

	var userInfo map[string]interface{}
	if err := json.Unmarshal(payload, &userInfo); err != nil {
		return nil, fmt.Errorf("an error has occurred while trying to get user data: %w", err)
	}

	email, found := userInfo["email"].(string)
	if !found {
		return nil, fmt.Errorf("email not found")
	}

	return &email, nil
}
