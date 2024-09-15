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
