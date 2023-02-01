package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ophoomo/line-beacon-demo/webhook"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("api")

	webhookRoute := webhook.WebhookRoute{Route: api}
	webhookRoute.New()

	return r
}
