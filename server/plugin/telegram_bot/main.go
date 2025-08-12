package telegram_bot

import (
	"github.com/gin-gonic/gin"
	"server/plugin/telegram_bot/router"
)

type Telegram_botPlugin struct{}

func CreateTelegram_botPlug() *Telegram_botPlugin {
	return &Telegram_botPlugin{}
}

func (*Telegram_botPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitTelegram_botRouter(group)
}

func (*Telegram_botPlugin) RouterPath() string {
	return "telegram_bot"
}
