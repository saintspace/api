package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"
)

func GenerateEmailSubscriptionToken(email string) (string, error) {
	// Generate a 32-byte random token
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Encode the token and email address as a base64 string
	token := base64.StdEncoding.EncodeToString(tokenBytes)
	emailToken := base64.StdEncoding.EncodeToString([]byte(email))

	// Add the email token to the token as a prefix, separated by a colon
	tokenWithPrefix := emailToken + ":" + token

	return tokenWithPrefix, nil
}

func ParseEmailSubscriptionToken(tokenWithPrefix string) (string, error) {
	// Split the token into email token and token parts
	parts := strings.Split(tokenWithPrefix, ":")
	if len(parts) != 2 {
		return "", errors.New("invalid token")
	}

	// Decode the email token and return the email address
	emailBytes, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return "", err
	}

	return string(emailBytes), nil
}
