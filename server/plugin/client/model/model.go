package model

type CodeReq struct {
	TgId string `json:"tg_id"` // 绑定用户的TGID
}

type RegisterReq struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Tgid      string `json:"tg_id"`     // 用户的TGID
	Code      string `json:"code"`      // tg验证码
	Captcha   string `json:"captcha"`   // 图片验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type RegisterWithInviteReq struct {
	RegisterReq
	InviteCode string `json:"invite_code"` // 邀请码
}

type ChangePasswordReq struct {
	Tgid        string `json:"tg_id"`        // 用户的TGID
	Code        string `json:"code"`         // tg验证码
	NewPassword string `json:"new_password"` // 新密码
}

type LoginReq struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// 是否启用TG验证进行登录
var EnableTGLogin = false

// 是否启用TG验证进行注册
var EnableTGRegister = false

type TGLoginControl struct {
	EnableTGLogin bool `json:"enable_tg_login"`
}

type TGRegisterControl struct {
	EnableTGRegister bool `json:"enable_tg_register"`
}

// 设置默认的校验注册和登录的配置
var (
	ConfigAuthorityId   uint
	ConfigTgBotToken    string
	ConfigCodeLength    int
	ConfigChannelId     string // 上面telegram的bot所在的频道的chat_id
	ConfigAdmins        string
	ConfigDefaultRouter string
)
