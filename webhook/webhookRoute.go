package webhook

import (
	"github.com/gin-gonic/gin"
)

type WebhookRoute struct {
	Route *gin.RouterGroup
}

func (r WebhookRoute) New() {
	service := webhookService{}
	handler := NewWebhookController(&service)
	route := r.Route.Group("webhook")
	route.POST("/", handler.WebHook)
}
