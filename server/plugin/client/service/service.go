package service

import (
	"context"
	"errors"
	"fmt"
	gvaGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ecsusers"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	tgrGlobal "github.com/flipped-aurora/gin-vue-admin/server/plugin/client/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/client/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/telegram_bot/service"
	gvaService "github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
	"github.com/mojocn/base64Captcha"
	"strings"
	"time"
)

type RegisterService struct{}

var eusrService = gvaService.ServiceGroupApp.EcsusersServiceGroup.EcsUsersService

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

func (e *RegisterService) Register(register model.RegisterReq) (string, request.CustomClaims, ecsusers.EcsUsers, error) {
	// 定义错误消息常量
	const (
		errTGCodeRetrieval = "存储的TG验证码获取错误: %v"
		errTGCodeMismatch  = "验证码填写错误: %v"
		errNotInChannel    = "检测到用户不在频道中: %v"
		errLoginStatus     = "获取登录状态错误: %v"
		errCaptcha         = "图片验证码错误"
		errEmptyUsername   = "用户名为空"
		errEmptyPassword   = "密码为空"
		errAccountCreation = "注册账户失败: %v"
		errRoleCreation    = "注册角色失败: %v"
	)
	ctx := context.Background()
	// 验证TG码
	code, err := gvaGlobal.GVA_REDIS.Get(ctx, register.Tgid).Result()
	if err != nil {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errTGCodeRetrieval, err)
	}
	if register.Code != code {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errTGCodeMismatch, register.Code)
	}
	// 检查用户是否在特定频道中
	_, err = service.ServiceGroupApp.IsTgMember(tgrGlobal.GlobalConfig.TgBotToken, register.Tgid, tgrGlobal.GlobalConfig.ChannelId)
	if err != nil {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errNotInChannel, err)
	}
	// 验证登录要求符合的格式
	if err := utils.Verify(register, utils.LoginVerify); err != nil {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errLoginStatus, err)
	}
	// 验证图片验证码
	if !base64Captcha.DefaultMemStore.Verify(register.CaptchaId, register.Captcha, true) {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, errors.New(errCaptcha)
	}
	// 检查用户名和密码是否为空
	if register.Username == "" {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, errors.New(errEmptyUsername)
	}
	if register.Password == "" {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, errors.New(errEmptyPassword)
	}
	// 创建用户的订阅账户(表)
	var euser *ecsusers.EcsUsers
	euser.UUID, _ = uuid.NewV4()
	euser.Username = register.Username
	euser.Password = utils.BcryptHash(register.Password)
	euser.Nickname = "注册用户"
	euser.TGID = register.Tgid
	euser.AuthorityID = tgrGlobal.GlobalConfig.AuthorityId
	// 创建用户的订阅账户
	if err := eusrService.CreateEcsUsers(euser); err != nil {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errAccountCreation, err)
	}
	// 同时创建对应的系统账户：用户的ID、用户的用户名、密码、uuid相对应 - 避免casbin无法鉴权
	sysUser := &system.SysUser{
		UUID:        euser.UUID,
		Username:    register.Username,
		Password:    euser.Password,
		NickName:    "注册用户",
		Phone:       "",
		AuthorityId: tgrGlobal.GlobalConfig.AuthorityId,
		Authority: system.SysAuthority{
			DefaultRouter: "dashboard",
			AuthorityId:   tgrGlobal.GlobalConfig.AuthorityId,
		},
	}
	// 开始数据库事务
	tx := gvaGlobal.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 创建用户账户
	if err := tx.Create(sysUser).Error; err != nil {
		tx.Rollback()
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errAccountCreation, err)
	}
	// 创建用户角色
	userAuthority := &system.SysUserAuthority{
		SysUserId:               sysUser.ID,
		SysAuthorityAuthorityId: tgrGlobal.GlobalConfig.AuthorityId,
	}
	if err := tx.Create(userAuthority).Error; err != nil {
		tx.Rollback()
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errRoleCreation, err)
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, err
	}
	// 登录
	token, claims, err := utils.LoginToken(euser)
	if err != nil {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errLoginStatus, err)
	}
	return token, claims, *euser, err
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
	var user ecsusers.EcsUsers
	if err := tx.Where("tg_id = ?", changer.Tgid).First(&user).Error; err != nil {
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

func (e *RegisterService) Login(loginUser model.LoginReq, key string) (string, request.CustomClaims, ecsusers.EcsUsers, error) {
	// 定义错误消息常量
	const (
		errAdminNotFound     = "检测不到管理员账户"
		errUserNotFound      = "检测不到该用户：%v"
		errNotInChannel      = "检测到用户不在频道中：%v"
		errLoginVerification = "登录验证失败: %v"
		errLoginFailed       = "用户名不存在或者密码错误: %v"
		errUserDisabled      = "用户被禁止登录"
		errCaptcha           = "验证码错误"
	)
	// 检查用户是否为管理员
	admins := strings.Split(tgrGlobal.GlobalConfig.Admins, ",")
	if len(admins) == 0 {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errAdminNotFound)
	}
	var isAdmin bool
	for _, a := range admins {
		if a == loginUser.Username {
			isAdmin = true
			break
		}
	}
	// 如果不是管理员，检查用户是否在特定频道中，如果是管理员，不需要检测
	if !isAdmin {
		var ecsUser ecsusers.EcsUsers
		if err := gvaGlobal.GVA_DB.Model(&ecsusers.EcsUsers{}).Where("username = ?", loginUser.Username).
			First(&ecsUser).Error; err != nil {
			return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errUserNotFound, err)
		}
		if _, err := service.ServiceGroupApp.IsTgMember(tgrGlobal.GlobalConfig.TgBotToken, ecsUser.TGID,
			tgrGlobal.GlobalConfig.ChannelId); err != nil {
			return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errNotInChannel, err)
		}
	}
	// 验证登录信息是否符合逻辑
	if err := utils.Verify(loginUser, utils.LoginVerify); err != nil {
		return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errLoginVerification, err)
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
		// 验证用户名和密码进行登录，得到token和用户信息
		u, err := eusrService.Login(loginUser.Username, loginUser.Password)
		if err != nil {
			gvaGlobal.BlackCache.Increment(key, 1)
			return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errLoginFailed, err)
		}
		if *u.IsFrozen {
			gvaGlobal.GVA_LOG.Error("登陆失败! 用户被冻结，禁止登录!")
			gvaGlobal.BlackCache.Increment(key, 1)
			return "", request.CustomClaims{}, ecsusers.EcsUsers{}, errors.New(errUserDisabled)
		}
		token, claims, err := utils.LoginToken(u)
		if err != nil {
			gvaGlobal.BlackCache.Increment(key, 1)
			return "", request.CustomClaims{}, ecsusers.EcsUsers{}, fmt.Errorf(errLoginFailed, err)
		}
		return token, claims, u, err
	}
	// 验证码错误，增加计数
	gvaGlobal.BlackCache.Increment(key, 1)
	return "", request.CustomClaims{}, ecsusers.EcsUsers{}, errors.New(errCaptcha)
}