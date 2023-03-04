package taskpub

type Task struct {
	TaskName      string `json:"taskName"`
	CorrelationId string `json:"correlationId"`
	TaskDetails   string `json:"taskDetails"`
}

type EmailSendTask struct {
	TemplateName  string                  `json:"templateName"`
	SenderAddress string                  `json:"senderAddress"`
	SubjectLine   string                  `json:"subjectLine"`
	ToAddresses   []string                `json:"toAddresses"`
	Parameters    EmailSendTaskParameters `json:"parameters"`
}

type EmailSendTaskParameters struct {
	VerificationLink string `json:"verificationLink"`
}
