package taskpub

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// TaskPublisher is responsible for publishing tasks to be completed later
type TaskPublisher struct {
	notifier iNotifier
	config   iConfig
}

type iNotifier interface {
	PublishTask(message string) error
}

type iConfig interface {
	WebAppDomainName() string
	MainTransactionalSendingAddress() string
}

func New(notifier iNotifier, config iConfig) *TaskPublisher {
	return &TaskPublisher{
		notifier: notifier,
		config:   config,
	}
}

func (s *TaskPublisher) PublishEmailVerificationTask(email, token string) error {
	escapedToken := url.QueryEscape(token)
	linkTemplate := "https://%s/saintspace/verify-email?token=%s"
	link := fmt.Sprintf(linkTemplate, s.config.WebAppDomainName(), escapedToken)
	emailSendTask := EmailSendTask{
		TemplateName:  "email-subscription-verification",
		SenderAddress: s.config.MainTransactionalSendingAddress(),
		SubjectLine:   "SaintSpace: Confirm Your Subscription",
		ToAddresses:   []string{email},
		Parameters: EmailSendTaskParameters{
			VerificationLink: link,
		},
	}
	emailSendTaskBytes, err := json.Marshal(emailSendTask)
	if err != nil {
		return fmt.Errorf("error while marshaling email send task details => %v", err.Error())
	} else {
		task := Task{
			TaskName:      "email-send",
			CorrelationId: "",
			TaskDetails:   string(emailSendTaskBytes),
		}
		taskBytes, err := json.Marshal(task)
		if err != nil {
			return fmt.Errorf("error while marshaling email send task => %v", err.Error())
		} else {
			if err := s.notifier.PublishTask(string(taskBytes)); err != nil {
				return fmt.Errorf("error while publishing email send task => %v", err.Error())
			}
		}
	}
	return nil
}
