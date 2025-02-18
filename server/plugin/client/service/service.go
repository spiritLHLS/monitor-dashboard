package service

import (
	"context"
	"errors"
	"fmt"
	gvaGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ecsusers"
	"github.com/flipped-aurora/gin-vue-admin/server/model/invite_codes"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/client/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/telegram_bot/service"
	gvaService "github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"strings"
	"time"
)

type RegisterService struct{}

var eusrService = gvaService.ServiceGroupApp.EcsusersServiceGroup.EcsUsersService

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

func (e *RegisterService) Code(tgid string) (err error) {
	// 参数验证
	if tgid == "" {
		return errors.New("tgid 不能为空")
	}
	// 检查 ServiceGroupApp
	if service.ServiceGroupApp == nil {
		return errors.New("ServiceGroupApp 未初始化")
	}
	if model.ConfigCodeLength == 0 {
		return errors.New("model.ConfigCodeLength 未初始化")
	}
	if model.ConfigTgBotToken == "" {
		return errors.New("model.ConfigTgBotToken 未初始化")
	}
	// 制作验证码
	code := utils.RandomString(model.ConfigCodeLength)
	if code == "" {
		return errors.New("TG验证码未初始化")
	}
	// 发送验证码
	_, err = service.ServiceGroupApp.SendTgMessage(model.ConfigTgBotToken, tgid,
		fmt.Sprintf("验证码：<code>%v</code>", code), "html")
	if err != nil {
		return fmt.Errorf("发送TG验证码错误：%v", err)
	}
	// 存储验证码
	ctx := context.Background()
	if gvaGlobal.GVA_REDIS == nil {
		return errors.New("Redis 客户端未初始化")
	}
	err = gvaGlobal.GVA_REDIS.Set(ctx, tgid, code, 5*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("存储验证码到 Redis 失败: %v", err)
	}
	return nil
}

func (e *RegisterService) commonRegisterLogic(register model.RegisterReq) (string, request.CustomClaims, system.SysUser, error) {
	ctx := context.Background()
	if model.EnableTGRegister {
		// 验证TG码
		code, err := gvaGlobal.GVA_REDIS.Get(ctx, register.Tgid).Result()
		if err != nil {
			return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errTGCodeRetrieval, err)
		}
		if register.Code != code {
			return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errTGCodeMismatch, register.Code)
		}
		// 检查用户是否在特定频道中
		_, err = service.ServiceGroupApp.IsTgMember(model.ConfigTgBotToken, register.Tgid, model.ConfigChannelId)
		if err != nil {
			return "", request.CustomClaims{}, system.SysUser{}, errors.New(errNotInChannel)
		}
	}
	// 验证登录要求符合的格式
	if err := utils.Verify(register, utils.LoginVerify); err != nil {
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errLoginStatus, err)
	}
	// 验证图片验证码
	if !base64Captcha.DefaultMemStore.Verify(register.CaptchaId, register.Captcha, true) {
		return "", request.CustomClaims{}, system.SysUser{}, errors.New(errCaptcha)
	}
	// 检查用户名和密码是否为空
	if register.Username == "" {
		return "", request.CustomClaims{}, system.SysUser{}, errors.New(errEmptyUsername)
	}
	if register.Password == "" {
		return "", request.CustomClaims{}, system.SysUser{}, errors.New(errEmptyPassword)
	}
	// 创建用户的订阅账户(表)
	euser := &ecsusers.EcsUsers{}
	euser.UUID, _ = uuid.NewUUID()
	euser.Username = register.Username
	euser.Password = utils.BcryptHash(register.Password)
	euser.Nickname = register.Username + " 订阅用户"
	if register.Tgid != "" {
		euser.TGID = register.Tgid
	}
	euser.AuthorityID = model.ConfigAuthorityId
	// 创建用户的订阅账户
	if err := eusrService.CreateEcsUsers(euser); err != nil {
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errAccountCreation, err)
	}
	// 同时创建对应的系统账户：用户的ID、用户的用户名、密码、uuid相对应 - 避免casbin无法鉴权
	sysUser := &system.SysUser{
		UUID:        euser.UUID,
		Username:    register.Username,
		Password:    euser.Password,
		NickName:    register.Username + " 订阅用户",
		Phone:       "",
		AuthorityId: model.ConfigAuthorityId,
		Authority: system.SysAuthority{
			DefaultRouter: model.ConfigDefaultRouter,
			AuthorityId:   model.ConfigAuthorityId,
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
	if err := tx.Create(sysUser).Select("ID").Error; err != nil {
		tx.Rollback()
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errAccountCreation, err)
	}
	// 创建用户角色
	userAuthority := &system.SysUserAuthority{
		SysUserId:               sysUser.ID,
		SysAuthorityAuthorityId: model.ConfigAuthorityId,
	}
	if err := tx.Create(userAuthority).Error; err != nil {
		tx.Rollback()
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errRoleCreation, err)
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return "", request.CustomClaims{}, system.SysUser{}, err
	}
	// 登录
	token, claims, err := utils.LoginToken(euser)
	if err != nil {
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errLoginStatus, err)
	}
	return token, claims, *sysUser, err
}

func (e *RegisterService) Register(register model.RegisterReq) (string, request.CustomClaims, system.SysUser, error) {
	if !ecsusers.EcsUserPublicRegisterStatus {
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf("管理员未开放公共注册权限，请联系管理员或使用邀请码注册")
	}
	return e.commonRegisterLogic(register)
}

func (e *RegisterService) RegisterWithInvite(register model.RegisterWithInviteReq) (string, request.CustomClaims, system.SysUser, error) {
	if register.InviteCode == "" && len(register.InviteCode) != 8 {
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf("无效邀请码，请使用有效的邀请码")
	}
	var dbCode invite_codes.InviteCodes
	codeErr := gvaGlobal.GVA_DB.Model(&invite_codes.InviteCodes{}).Where("code = ?", register.InviteCode).First(&dbCode).Error
	if codeErr != nil {
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf("数据库查无此邀请码")
	}
	if *dbCode.Status > 0 {
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf("此邀请码已被使用")
	}
	if *dbCode.Status < 0 {
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf("此邀请码已过期")
	}
	token, claims, sysUser, err := e.commonRegisterLogic(register.RegisterReq)
	if err == nil {
		// 标记邀请码已被使用，同时绑定对应用户的uuid
		currentTime := time.Now()
		if err := gvaGlobal.GVA_DB.Model(&invite_codes.InviteCodes{}).
			Where("code = ?", register.InviteCode).
			Updates(map[string]interface{}{
				"status":    1,
				"used_at":   currentTime,
				"user_uuid": sysUser.UUID,
			}).Error; err != nil {
			return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf("邀请码标记使用操作失败")
		}

	}
	return token, claims, sysUser, err
}

func (e *RegisterService) ChangePassword(changer model.ChangePasswordReq) error {
	if changer.Tgid == "" {
		return fmt.Errorf("由于当前用户未绑定TGID，无法通过TG进行密码重置，请联系管理员手动重置密码")
	}
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

func (e *RegisterService) Login(loginUser model.LoginReq, key string) (string, request.CustomClaims, system.SysUser, error) {
	// 定义错误消息常量
	const (
		errAdminNotFound     = "检测不到管理员账户"
		errUserNotFound      = "检测不到该用户：%v"
		errNotInChannel      = "检测到用户不在频道中"
		errLoginVerification = "登录验证失败: %v"
		errLoginFailed       = "用户名不存在或者密码错误: %v"
		errUserDisabled      = "用户被禁止登录"
		errCaptcha           = "验证码错误"
		errSysLoginFailed    = "系统用户通过uuid查询失败"
	)
	// 检查用户是否为管理员
	admins := strings.Split(model.ConfigAdmins, ",")
	if len(admins) == 0 {
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errAdminNotFound)
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
			return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errUserNotFound, err)
		}
		if model.EnableTGLogin {
			if _, err := service.ServiceGroupApp.IsTgMember(model.ConfigTgBotToken, ecsUser.TGID,
				model.ConfigChannelId); err != nil {
				return "", request.CustomClaims{}, system.SysUser{}, errors.New(errNotInChannel)
			}
		}
	}
	// 验证登录信息是否符合逻辑
	if err := utils.Verify(loginUser, utils.LoginVerify); err != nil {
		return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errLoginVerification, err)
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
			return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errLoginFailed, err)
		}
		if *u.IsFrozen {
			gvaGlobal.GVA_LOG.Error("登陆失败! 用户被冻结，禁止登录!")
			gvaGlobal.BlackCache.Increment(key, 1)
			return "", request.CustomClaims{}, system.SysUser{}, errors.New(errUserDisabled)
		}
		token, claims, err := utils.LoginToken(u)
		if err != nil {
			gvaGlobal.BlackCache.Increment(key, 1)
			return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errLoginFailed, err)
		}
		var sysUser system.SysUser
		err = gvaGlobal.GVA_DB.Model(&system.SysUser{}).Where("uuid = ?", u.UUID).First(&sysUser).Error
		if err != nil {
			gvaGlobal.BlackCache.Increment(key, 1)
			return "", request.CustomClaims{}, system.SysUser{}, fmt.Errorf(errSysLoginFailed, err)
		}
		if !isAdmin {
			sysUser.Authority = system.SysAuthority{
				DefaultRouter: model.ConfigDefaultRouter,
				AuthorityId:   model.ConfigAuthorityId,
			}
		} else {
			sysUser.Authority = system.SysAuthority{
				DefaultRouter: "home",
				AuthorityId:   888,
			}
		}
		return token, claims, sysUser, err
	}
	// 验证码错误，增加计数
	gvaGlobal.BlackCache.Increment(key, 1)
	return "", request.CustomClaims{}, system.SysUser{}, errors.New(errCaptcha)
}
