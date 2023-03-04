package taskpub

import (
	"encoding/json"
	"fmt"
)

// TaskPublisher is responsible for publishing tasks to be completed later
type TaskPublisher struct {
	notifier iNotifier
}

type iNotifier interface {
	PublishTask(message string) error
}

func New(notifier iNotifier) *TaskPublisher {
	return &TaskPublisher{
		notifier: notifier,
	}
}

func (s *TaskPublisher) PublishEmailVerificationTask(email, token string) error {
	link := fmt.Sprintf("https://dev.saintspace.app/saintspace/verify-email?email=%s&token=%s", email, token)
	emailSendTask := EmailSendTask{
		TemplateName:  "email-subscription-verification",
		SenderAddress: "noreply@dev-messages.saintspace.app",
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
