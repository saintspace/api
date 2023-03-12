package handler

type RouteHandler struct {
	emailService  iEmailService
	loggerService iLoggerService
}

func New(emailService iEmailService, loggerService iLoggerService) *RouteHandler {
	return &RouteHandler{
		emailService:  emailService,
		loggerService: loggerService,
	}
}

type iEmailService interface {
	CreateEmailSubscription(email string) error
	EmailSubscriptionExists(email string) (bool, error)
	IsValidEmail(email string) bool
	VerifyEmailwithSubscriptionToken(token string) error
}

type iLoggerService interface {
	InfoWithContext(message string, keysAndValues ...interface{})
	ErrorWithContext(message string, keysAndValues ...interface{})
	DebugWithContext(message string, keysAndValues ...interface{})
}
