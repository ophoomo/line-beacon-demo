package webhook

type webhookService struct{}

func NewWebhookService() WebhookService {
	return &webhookService{}
}

const LINE_MESSAGING_API = "https://api.line.me/v2/bot"

func (s webhookService) SendMessage(token string) error {
	return nil
}
