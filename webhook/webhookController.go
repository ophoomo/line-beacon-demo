package webhook

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type WebhookController struct {
	webhookService WebhookService
}

func NewWebhookController(service WebhookService) *WebhookController {
	return &WebhookController{
		webhookService: service,
	}
}

func (c WebhookController) WebHook(ctx *gin.Context) {
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		ctx.JSON(500, "")
		return
	}
	events, err := bot.ParseRequest(ctx.Request)
	if err != nil {
		ctx.JSON(500, "")
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			replyToken := event.ReplyToken
			var messages []linebot.SendingMessage
			messages = append(messages, linebot.NewTextMessage("hello phoom"))
			if _, err := bot.ReplyMessage(replyToken, messages...).Do(); err != nil {
				ctx.JSON(500, "")
				return
			}
		}
	}
	ctx.JSON(200, "")
}
