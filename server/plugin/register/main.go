package register

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/register/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/register/router"
	"github.com/gin-gonic/gin"
)

type RegisterPlugin struct {
}

func CreateRegisterPlug(AuthorityId uint, TgBotToken string, CodeLength int, ChannelId string) *RegisterPlugin {
	global.GlobalConfig.AuthorityId = AuthorityId
	global.GlobalConfig.TgBotToken = TgBotToken
	global.GlobalConfig.CodeLength = CodeLength
	global.GlobalConfig.ChannelId = ChannelId
	return &RegisterPlugin{}
}

func (*RegisterPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitRegisterRouter(group)
}

func (*RegisterPlugin) RouterPath() string {
	return "register"
}
