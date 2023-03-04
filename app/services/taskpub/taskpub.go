package taskpub

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

func (s *TaskPublisher) PublishEmailVerificationTask(email string) error {

	return nil
}
