package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/allpdspiders"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/client"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/spiders"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/telegram_bot"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/telegrampush"
	"github.com/gin-gonic/gin"
)

func InstallPlugin(PrivateGroup *gin.RouterGroup, PublicGroup *gin.RouterGroup, engine *gin.Engine) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Info("项目暂未初始化，无法安装插件，初始化后重启项目即可完成插件安装")
		return
	}
	bizPluginV1(PrivateGroup, PublicGroup)
	bizPluginV2(engine)
	// 注册TG发信插件
	PluginInit(PrivateGroup, telegram_bot.CreateTelegram_botPlug())
	// 注册用户登录注册校验插件
	tgrAdmins := global.GVA_VP.GetString("tgr.admins")
	defaultRouter := global.GVA_VP.GetString("tgr.default_router")
	authorityID := global.GVA_VP.GetInt("tgr.authority_id")
	tgbotToken := global.GVA_VP.GetString("tgr.tgbot_token")
	codeLength := global.GVA_VP.GetInt("tgr.verification_code_length")
	channelChatID := global.GVA_VP.GetString("tgr.channel_chat_id")
	PluginInit(PublicGroup, client.CreateRegisterPlug(uint(authorityID), defaultRouter, tgbotToken, codeLength, channelChatID, tgrAdmins))
	// Telegram频道信息推送插件注册
	PluginInit(PrivateGroup, telegrampush.CreateTelegramPushPlug())
	// 多节点爬虫插件注册
	PluginInit(PublicGroup, spiders.CreateSpidersPlug())
	// 商家所有商品爬虫推送信息插件注册
	PluginInit(PublicGroup, allpdspiders.CreateAllPdSpidersPlug())
}
