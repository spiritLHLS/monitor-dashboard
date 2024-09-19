package config

type Client struct {
	Admins        string // 管理员
	Name          string // 用户名
	AuthorityId   uint   // 权限ID
	DefaultRouter string // 默认路由
	TgBotToken    string // tg的bot的token
	CodeLength    int    // tg的验证码的长度
	ChannelId     string // 上面telegram的bot所在的频道的chat_id
}
