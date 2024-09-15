package config

type Register struct {
	Name        string // 用户名
	AuthorityId uint   // 权限ID
	TgBotToken  string // tg的bot的token
	CodeLength  int    // tg的验证码的长度
	ChannelId   string // 上面telegram的bot所在的频道的chat_id
}
