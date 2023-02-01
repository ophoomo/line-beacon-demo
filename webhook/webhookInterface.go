package webhook

type WebhookService interface {
	SendMessage(token string) error
}
