package service

import (
	"context"
	"errors"
	"fmt"
	gvaGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	plugGlobal "github.com/flipped-aurora/gin-vue-admin/server/plugin/register/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/register/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/telegram_bot/service"
	userServiceSystem "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
	"github.com/mojocn/base64Captcha"
	"time"
)

type RegisterService struct{}

func (e *RegisterService) Code(tgid string) (err error) {
	// 制作四位数code
	code := utils.RandomString(plugGlobal.GlobalConfig.CodeLength)
	// 发送code
	_, err = service.ServiceGroupApp.SendTgMessage(plugGlobal.GlobalConfig.TgBotToken, tgid,
		fmt.Sprintf("验证码：<code>%v</code>", code), "html")
	if err != nil {
		return errors.New(fmt.Sprintf("发送TG验证码错误：%v", err))
	}
	// 存储code
	ctx := context.Background()
	gvaGlobal.GVA_REDIS.Set(ctx, tgid, code, 5*time.Minute)
	return nil
}

func (e *RegisterService) Register(register model.RegisterReq) (res *system.SysUser, err error) {
	res = &system.SysUser{}
	// 检测tgcode是否正确
	ctx := context.Background()
	code, err := gvaGlobal.GVA_REDIS.Get(ctx, register.Tgid).Result()
	if err != nil {
		return res, errors.New(fmt.Sprintf("存储的TG验证码获取错误：%v", err))
	} else if register.Code != code {
		return res, errors.New(fmt.Sprintf("验证码填写错误：%v", register.Code))
	}
	// 检测用户是否在特定的频道中
	_, err = service.ServiceGroupApp.IsTgMember(plugGlobal.GlobalConfig.TgBotToken, register.Tgid,
		plugGlobal.GlobalConfig.ChannelId)
	if err != nil {
		return res, errors.New(fmt.Sprintf("检测到用户不在频道中：%v", err))
	}
	// 获取注册的信息
	if err := utils.Verify(register, utils.LoginVerify); err != nil {
		return res, errors.New(fmt.Sprintf("获取登录状态错误：%v", err))
	}
	var (
		store = base64Captcha.DefaultMemStore
		user  system.SysUser
		us    *userServiceSystem.UserService
	)
	if !store.Verify(register.CaptchaId, register.Captcha, true) {
		return res, errors.New(fmt.Sprintf("图片验证码错误"))
	}
	// 加密密码
	plain_password := register.Password
	register.Password = utils.BcryptHash(register.Password)
	// 创建用户需要传入的信息
	// 用 Phone 字段存用户的 TGID 了
	UUID, _ := uuid.NewV4()
	u := &system.SysUser{UUID: UUID, Username: register.Username, Password: register.Password, NickName: "注册用户",
		Phone: register.Tgid, AuthorityId: plugGlobal.GlobalConfig.AuthorityId,
		Authority: system.SysAuthority{DefaultRouter: "dashboard", AuthorityId: plugGlobal.GlobalConfig.AuthorityId}}
	// 检测传入信息是否为空
	if u.Username == "" {
		return res, errors.New(fmt.Sprintf("用户名为空：%v", err))
	}
	if u.Password == "" {
		return res, errors.New(fmt.Sprintf("密码为空：%v", err))
	}
	// 检测账户是否存在
	err = gvaGlobal.GVA_DB.Where("phone = ?", u.Phone).First(&user).Error
	if err == nil {
		return res, errors.New(fmt.Sprintf("该TGID已注册，用户名为：%v", user.Username))
	}
	// 创建用户账户
	err = gvaGlobal.GVA_DB.Create(&u).Error
	if err != nil {
		return res, errors.New(fmt.Sprintf("注册账户失败：%v", err))
	}
	// 创建用户角色
	gvaGlobal.GVA_DB.Where("phone = ?", u.Phone).First(&user) // 查询对应tgid的用户id存到user里
	a := &system.SysUserAuthority{SysUserId: user.ID, SysAuthorityAuthorityId: plugGlobal.GlobalConfig.AuthorityId}
	err = gvaGlobal.GVA_DB.Create(&a).Error
	if err != nil {
		return res, errors.New(fmt.Sprintf("注册角色失败：%v", err))
	}
	// 登录前密码需要解密，使用明文密码
	u.Password = plain_password
	if res, err = us.Login(u); err != nil {
		return res, errors.New("登陆失败!")
	}
	return res, nil
}

func (e *RegisterService) ChangePassword(changer model.ChangePasswordReq) (err error) {
	var user system.SysUser
	// 检测tgcode是否正确
	ctx := context.Background()
	code, err := gvaGlobal.GVA_REDIS.Get(ctx, changer.Tgid).Result()
	if err != nil {
		return errors.New(fmt.Sprintf("存储的TG验证码获取错误：%v", err))
	} else if changer.Code != code {
		return errors.New(fmt.Sprintf("验证码填写错误：%v", changer.Code))
	}
	// 检测账户是否存在
	err = gvaGlobal.GVA_DB.Where("phone = ?", changer.Tgid).First(&user).Error
	if err != nil {
		return errors.New(fmt.Sprintf("查询不到该TGID的用户：%v", changer.Tgid))
	}
	// 修改密码
	u := &system.SysUser{GVA_MODEL: gvaGlobal.GVA_MODEL{ID: user.ID}}
	if err = gvaGlobal.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return errors.New(fmt.Sprintf("查询对应ID的用户失败：%v", err))
	}
	user.Password = utils.BcryptHash(changer.NewPassword)
	err = gvaGlobal.GVA_DB.Save(&user).Error
	return err
}

// 类型转换
// server/api/v1/system/sys_captcha.go
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

func (e *RegisterService) Login(loginUser model.LoginReq, key string) (res *system.SysUser, err error) {
	res = &system.SysUser{}
	var (
		us    *userServiceSystem.UserService
		sysus system.SysUser
	)
	var store = base64Captcha.DefaultMemStore // server/api/v1/system/sys_captcha.go
	// 如果用户名是管理员，则不检测
	if loginUser.Username != "admin" && loginUser.Username != "a303176530" {
		// 检测普通注册用户是否在特定的频道中
		err = gvaGlobal.GVA_DB.Where("username = ?", loginUser.Username).First(&sysus).Error
		if err != nil {
			return res, errors.New(fmt.Sprintf("检测不到该用户：%v", err))
		}
		_, err = service.ServiceGroupApp.IsTgMember(plugGlobal.GlobalConfig.TgBotToken, sysus.Phone,
			plugGlobal.GlobalConfig.ChannelId)
		if err != nil {
			return res, errors.New(fmt.Sprintf("检测到用户不在频道中：%v", err))
		}
	}
	// 后续内容修改自 server/api/v1/system/sys_user.go
	if err := utils.Verify(loginUser, utils.LoginVerify); err != nil {
		return res, errors.New(fmt.Sprintf("登录验证失败: %v", err))
	}
	// 判断验证码是否开启
	openCaptcha := gvaGlobal.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := gvaGlobal.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := gvaGlobal.BlackCache.Get(key)
	if !ok {
		gvaGlobal.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)
	if !oc || (loginUser.CaptchaId != "" && loginUser.Captcha != "" && store.Verify(loginUser.CaptchaId, loginUser.Captcha, true)) {
		u := &system.SysUser{Username: loginUser.Username, Password: loginUser.Password}
		res, err := us.Login(u)
		if err != nil {
			// 验证码次数+1
			gvaGlobal.BlackCache.Increment(key, 1)
			return res, errors.New(fmt.Sprintf("用户名不存在或者密码错误: %v", err))
		}
		if res.Enable != 1 {
			gvaGlobal.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			gvaGlobal.BlackCache.Increment(key, 1)
			return res, errors.New("用户被禁止登录")
		}
		return res, nil
	}
	// 验证码次数+1
	gvaGlobal.BlackCache.Increment(key, 1)
	return res, errors.New("验证码错误")
}
