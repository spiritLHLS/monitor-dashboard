package initialize

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/plugin/allpdspiders"
	"server/plugin/client"
	"server/plugin/cryptourl"
	"server/plugin/obopush"
	"server/plugin/spiders"
	"server/plugin/telegram_bot"
	"server/plugin/telegrampush"
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
	PluginInit(PublicGroup, client.CreateRegisterPlug())
	// Telegram频道信息推送插件注册
	PluginInit(PrivateGroup, telegrampush.CreateTelegramPushPlug())
	// 多节点爬虫插件注册
	PluginInit(PublicGroup, spiders.CreateSpidersPlug())
	// 商家最新商品爬虫推送信息插件注册
	PluginInit(PublicGroup, allpdspiders.CreateAllPdSpidersPlug())
	// 生成加密链接的插件注册
	PluginInitV2(engine, cryptourl.Plugin)
	// 一对一商品消息推送插件注册
	PluginInitV2(engine, obopush.Plugin)
}
