package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/register"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/telegram_bot"
	"github.com/gin-gonic/gin"
)

func InstallPlugin(PrivateGroup *gin.RouterGroup, PublicRouter *gin.RouterGroup, engine *gin.Engine) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Info("项目暂未初始化，无法安装插件，初始化后重启项目即可完成插件安装")
		return
	}
	bizPluginV1(PrivateGroup, PublicRouter)
	bizPluginV2(engine)
	// 注册TG发信插件
	PluginInit(PrivateGroup, telegram_bot.CreateTelegram_botPlug())
	// 注册用户登录注册校验插件
	tgrAdmins := global.GVA_VP.GetString("tgr.admins")
	authorityID := global.GVA_VP.GetInt("tgr.authority_id")
	tgbotToken := global.GVA_VP.GetString("tgr.tgbot_token")
	codeLength := global.GVA_VP.GetInt("tgr.verification_code_length")
	channelChatID := global.GVA_VP.GetString("tgr.channel_chat_id")
	PluginInit(PublicRouter, register.CreateRegisterPlug(uint(authorityID), tgbotToken, codeLength, channelChatID, tgrAdmins))
}
