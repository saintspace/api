package handler

type RouteHandler struct {
	emailService iEmailService
}

func New(emailService iEmailService) *RouteHandler {
	return &RouteHandler{
		emailService: emailService,
	}
}

type iEmailService interface {
	CreateEmailSubscription(email string) error
	EmailSubscriptionExists(email string) (bool, error)
	IsValidEmail(email string) bool
	VerifyEmailwithSubscriptionToken(token string) error
}
