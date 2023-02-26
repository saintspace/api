package main

import (
	"net/mail"
)

func emailIsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
