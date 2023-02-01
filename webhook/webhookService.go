package webhook

type webhookService struct{}

func NewWebhookService() WebhookService {
	return &webhookService{}
}
