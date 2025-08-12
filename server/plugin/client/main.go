package client

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/plugin/client/model"
	"server/plugin/client/router"
)

type RegisterPlugin struct {
}

func CreateRegisterPlug() *RegisterPlugin {
	model.ConfigAuthorityId = uint(global.GVA_VP.GetInt("tgr.authority_id"))
	model.ConfigTgBotToken = global.GVA_VP.GetString("tgr.tgbot_token")
	model.ConfigCodeLength = global.GVA_VP.GetInt("tgr.verification_code_length")
	model.ConfigChannelId = global.GVA_VP.GetString("tgr.channel_chat_id")
	model.ConfigAdmins = global.GVA_VP.GetString("tgr.admins")
	model.ConfigDefaultRouter = global.GVA_VP.GetString("tgr.default_router")
	return &RegisterPlugin{}
}

func (*RegisterPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitRegisterRouter(group)
}

func (*RegisterPlugin) RouterPath() string {
	return "client"
}
