package router

import (
	"github.com/gin-gonic/gin"
	"server/plugin/telegram_bot/api"
)

type Telegram_botRouter struct{}

func (s *Telegram_botRouter) InitTelegram_botRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.Telegram_botApi
	{
		plugRouter.POST("sendMessage", plugApi.SendMessage)
		plugRouter.POST("isMember", plugApi.IsMember)
		plugRouter.POST("checkTgBot", plugApi.CheckTgBot)
	}
}
