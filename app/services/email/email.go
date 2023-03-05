package email

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

// TaskPublisher is responsible for publishing tasks to be completed later
type EmailService struct {
	datastore     iDatastore
	taskPublisher iTaskPublisher
}

func New(datastore iDatastore, taskPublisher iTaskPublisher) *EmailService {
	return &EmailService{
		datastore:     datastore,
		taskPublisher: taskPublisher,
	}
}

type iDatastore interface {
	CreateEmailSubscription(email string, subscriptionToken string, isVerified bool) error
	CheckEmailSubscriptionExists(email string) (bool, error)
	VerifyEmailSubscription(email string) error
}

type iTaskPublisher interface {
	PublishEmailVerificationTask(email, token string) error
}

func (s *EmailService) IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s *EmailService) EmailSubscriptionExists(email string) (bool, error) {
	return s.datastore.CheckEmailSubscriptionExists(email)
}

func (s *EmailService) CreateEmailSubscription(email string) error {
	subscriptionToken, err := generateEmailSubscriptionToken(email)
	if err != nil {
		return fmt.Errorf("error while trying to generate a new email subscription token => %v", err.Error())
	}
	err = s.datastore.CreateEmailSubscription(email, subscriptionToken, false)
	if err != nil {
		return fmt.Errorf("error while trying to create email subscription in datastore => %v", err.Error())
	}
	err = s.taskPublisher.PublishEmailVerificationTask(email, subscriptionToken)
	if err != nil {
		fmt.Printf("error while attempting to publish email verification task {email: %s | error: %s}\n", email, err.Error())
	}
	return nil
}

func (s *EmailService) VerifyEmailwithSubscriptionToken(token string) error {
	email, err := parseEmailFromSubscriptionToken(token)
	if err != nil {
		return fmt.Errorf("error while trying to parse email from token => %v", err.Error())
	}
	subcriptionExists, err := s.datastore.CheckEmailSubscriptionExists(email)
	if err != nil {
		return fmt.Errorf("error while checking if email subscription exists => %v", err.Error())
	}
	if !subcriptionExists {
		return fmt.Errorf("email subscription doesn't exist {email: %v}", email)
	}
	err = s.datastore.VerifyEmailSubscription(email)
	if err != nil {
		return fmt.Errorf("error while attempting to verify existing subscription => %v", err.Error())
	}
	return nil
}

func generateEmailSubscriptionToken(email string) (string, error) {
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

func parseEmailFromSubscriptionToken(tokenWithPrefix string) (string, error) {
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
