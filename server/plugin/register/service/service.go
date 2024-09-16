package service

import (
	"context"
	"errors"
	"fmt"
	gvaGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	tgrGlobal "github.com/flipped-aurora/gin-vue-admin/server/plugin/register/global"
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
	code := utils.RandomString(tgrGlobal.GlobalConfig.CodeLength)
	// 发送code
	_, err = service.ServiceGroupApp.SendTgMessage(tgrGlobal.GlobalConfig.TgBotToken, tgid,
		fmt.Sprintf("验证码：<code>%v</code>", code), "html")
	if err != nil {
		return errors.New(fmt.Sprintf("发送TG验证码错误：%v", err))
	}
	// 存储code
	ctx := context.Background()
	gvaGlobal.GVA_REDIS.Set(ctx, tgid, code, 5*time.Minute)
	return nil
}

func (e *RegisterService) Register(register model.RegisterReq) (*system.SysUser, error) {
	// 定义错误消息常量
	const (
		errTGCodeRetrieval = "存储的TG验证码获取错误: %v"
		errTGCodeMismatch  = "验证码填写错误: %v"
		errNotInChannel    = "检测到用户不在频道中: %v"
		errLoginStatus     = "获取登录状态错误: %v"
		errCaptcha         = "图片验证码错误"
		errEmptyUsername   = "用户名为空"
		errEmptyPassword   = "密码为空"
		errUsernameExists  = "用户名（%v）已被占用，请更换用户名"
		errTGIDExists      = "该TGID已注册，用户名为：%v"
		errAccountCreation = "注册账户失败: %v"
		errRoleCreation    = "注册角色失败: %v"
		errLogin           = "登陆失败"
	)
	ctx := context.Background()
	// 验证TG码
	code, err := gvaGlobal.GVA_REDIS.Get(ctx, register.Tgid).Result()
	if err != nil {
		return nil, fmt.Errorf(errTGCodeRetrieval, err)
	}
	if register.Code != code {
		return nil, fmt.Errorf(errTGCodeMismatch, register.Code)
	}
	// 检查用户是否在特定频道中
	_, err = service.ServiceGroupApp.IsTgMember(tgrGlobal.GlobalConfig.TgBotToken, register.Tgid, tgrGlobal.GlobalConfig.ChannelId)
	if err != nil {
		return nil, fmt.Errorf(errNotInChannel, err)
	}
	// 验证登录
	if err := utils.Verify(register, utils.LoginVerify); err != nil {
		return nil, fmt.Errorf(errLoginStatus, err)
	}
	// 验证图片验证码
	if !base64Captcha.DefaultMemStore.Verify(register.CaptchaId, register.Captcha, true) {
		return nil, errors.New(errCaptcha)
	}
	// 创建用户
	plainPassword := register.Password
	hashedPassword := utils.BcryptHash(register.Password)
	UUID, _ := uuid.NewV4()
	user := &system.SysUser{
		UUID:        UUID,
		Username:    register.Username,
		Password:    hashedPassword,
		NickName:    "注册用户",
		Phone:       register.Tgid,
		AuthorityId: tgrGlobal.GlobalConfig.AuthorityId,
		Authority: system.SysAuthority{
			DefaultRouter: "dashboard",
			AuthorityId:   tgrGlobal.GlobalConfig.AuthorityId,
		},
	}
	// 检查用户名和密码是否为空
	if user.Username == "" {
		return nil, errors.New(errEmptyUsername)
	}
	if plainPassword == "" {
		return nil, errors.New(errEmptyPassword)
	}
	// 开始数据库事务
	tx := gvaGlobal.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 检查用户名或TGID是否已存在
	var count int64
	if err := tx.Model(&system.SysUser{}).Where("username = ?", user.Username).Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if count > 0 {
		tx.Rollback()
		return nil, fmt.Errorf(errUsernameExists, user.Username)
	}
	if err := tx.Model(&system.SysUser{}).Where("phone = ?", user.Phone).Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if count > 0 {
		tx.Rollback()
		return nil, fmt.Errorf(errTGIDExists, user.Username)
	}
	// 创建用户账户
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf(errAccountCreation, err)
	}
	// 创建用户角色
	userAuthority := &system.SysUserAuthority{
		SysUserId:               user.ID,
		SysAuthorityAuthorityId: tgrGlobal.GlobalConfig.AuthorityId,
	}
	if err := tx.Create(userAuthority).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf(errRoleCreation, err)
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	// 登录
	user.Password = plainPassword
	us := &userServiceSystem.UserService{}
	res, err := us.Login(user)
	if err != nil {
		return nil, errors.New(errLogin)
	}
	return res, nil
}

func (e *RegisterService) ChangePassword(changer model.ChangePasswordReq) error {
	// 定义错误消息常量
	const (
		errTGCodeRetrieval = "存储的TG验证码获取错误：%v"
		errTGCodeMismatch  = "验证码填写错误：%v"
		errUserNotFound    = "查询不到该TGID的用户：%v"
		errUpdatePassword  = "更新密码失败：%v"
	)
	ctx := context.Background()
	// 验证TG验证码
	code, err := gvaGlobal.GVA_REDIS.Get(ctx, changer.Tgid).Result()
	if err != nil {
		return fmt.Errorf(errTGCodeRetrieval, err)
	}
	if changer.Code != code {
		return fmt.Errorf(errTGCodeMismatch, changer.Code)
	}
	// 开始数据库事务
	tx := gvaGlobal.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 查找用户
	var user system.SysUser
	if err := tx.Where("phone = ?", changer.Tgid).First(&user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf(errUserNotFound, changer.Tgid)
	}
	// 更新密码
	hashedPassword := utils.BcryptHash(changer.NewPassword)
	if err := tx.Model(&user).Update("password", hashedPassword).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf(errUpdatePassword, err)
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf(errUpdatePassword, err)
	}
	return nil
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

func (e *RegisterService) Login(loginUser model.LoginReq, key string) (*system.SysUser, error) {
	// 定义错误消息常量
	const (
		errUserNotFound      = "检测不到该用户：%v"
		errNotInChannel      = "检测到用户不在频道中：%v"
		errLoginVerification = "登录验证失败: %v"
		errLoginFailed       = "用户名不存在或者密码错误: %v"
		errUserDisabled      = "用户被禁止登录"
		errCaptcha           = "验证码错误"
	)
	// 检查用户是否为管理员
	isAdmin := loginUser.Username == "admin" || loginUser.Username == "a303176530"
	// 如果不是管理员，检查用户是否在特定频道中
	if !isAdmin {
		var sysus system.SysUser
		if err := gvaGlobal.GVA_DB.Where("username = ?", loginUser.Username).First(&sysus).Error; err != nil {
			return nil, fmt.Errorf(errUserNotFound, err)
		}
		if _, err := service.ServiceGroupApp.IsTgMember(tgrGlobal.GlobalConfig.TgBotToken, sysus.Phone, tgrGlobal.GlobalConfig.ChannelId); err != nil {
			return nil, fmt.Errorf(errNotInChannel, err)
		}
	}
	// 验证登录信息
	if err := utils.Verify(loginUser, utils.LoginVerify); err != nil {
		return nil, fmt.Errorf(errLoginVerification, err)
	}
	// 处理验证码逻辑
	openCaptcha := gvaGlobal.GVA_CONFIG.Captcha.OpenCaptcha
	openCaptchaTimeOut := gvaGlobal.GVA_CONFIG.Captcha.OpenCaptchaTimeOut
	v, _ := gvaGlobal.BlackCache.Get(key)
	if v == nil {
		gvaGlobal.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
		v = 1
	}
	needCaptcha := openCaptcha != 0 && openCaptcha >= interfaceToInt(v)
	captchaVerified := loginUser.CaptchaId != "" && loginUser.Captcha != "" &&
		base64Captcha.DefaultMemStore.Verify(loginUser.CaptchaId, loginUser.Captcha, true)
	if !needCaptcha || captchaVerified {
		// 验证用户名和密码
		u := &system.SysUser{Username: loginUser.Username, Password: loginUser.Password}
		us := &userServiceSystem.UserService{}
		res, err := us.Login(u)
		if err != nil {
			gvaGlobal.BlackCache.Increment(key, 1)
			return nil, fmt.Errorf(errLoginFailed, err)
		}
		if res.Enable != 1 {
			gvaGlobal.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			gvaGlobal.BlackCache.Increment(key, 1)
			return nil, errors.New(errUserDisabled)
		}
		return res, nil
	}
	// 验证码错误，增加计数
	gvaGlobal.BlackCache.Increment(key, 1)
	return nil, errors.New(errCaptcha)
}
