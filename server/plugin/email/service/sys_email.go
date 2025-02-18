package service

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ecsusers"
	"github.com/flipped-aurora/gin-vue-admin/server/model/privmsg"
	email_response "github.com/flipped-aurora/gin-vue-admin/server/plugin/email/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

// EmailService 邮件服务结构体
type EmailService struct{}

// SendEmailWithSMTPDetails 使用指定的SMTP配置发送邮件
// 参数:
// - smtpHost: SMTP服务器地址
// - smtpPort: SMTP服务器端口
// - from: 发件人邮箱地址
// - password: 发件人邮箱密码
// - to: 收件人邮箱地址
// - subject: 邮件主题
// - body: 邮件正文
func (e *EmailService) SendEmailWithSMTPDetails(smtpHost string, smtpPort int, from, password string, to string, subject, body string) error {
	if smtpHost == "" || smtpPort == 0 || from == "" || password == "" {
		err := utils.Email(to, subject, body)
		return err
	}
	// 构建SMTP服务器地址
	smtpServerAddr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)
	// 创建SMTP认证
	auth := smtp.PlainAuth("", from, password, smtpHost)
	// 准备邮件消息
	message := []byte(fmt.Sprintf("Subject: %s\r\n"+
		"From: %s\r\n"+
		"To: %s\r\n"+
		"\r\n"+
		"%s",
		subject,
		from,
		to,
		body))
	// 创建TLS配置
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         smtpHost,
	}
	// 建立安全连接
	conn, err := tls.Dial("tcp", smtpServerAddr, tlsConfig)
	if err != nil {
		return fmt.Errorf("创建TLS连接失败: %v", err)
	}
	// 创建新的SMTP客户端
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return fmt.Errorf("创建SMTP客户端失败: %v", err)
	}
	defer client.Close()
	// 进行身份认证
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("认证失败: %v", err)
	}
	// 设置发件人
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("设置发件人失败: %v", err)
	}
	// 设置收件人
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("设置收件人 %s 失败: %v", to, err)
	}
	// 发送邮件正文
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("创建数据写入器失败: %v", err)
	}
	_, err = w.Write(message)
	if err != nil {
		return fmt.Errorf("写入邮件正文失败: %v", err)
	}
	err = w.Close()
	if err != nil {
		return fmt.Errorf("关闭数据写入器失败: %v", err)
	}
	// 发送邮件
	err = client.Quit()
	if err != nil {
		return fmt.Errorf("发送邮件失败: %v", err)
	}
	return nil
}

// EmailTest 邮件测试函数
func (e *EmailService) EmailTest() (err error) {
	subject := "测试"
	body := "这是一封测试邮件"
	err = utils.EmailTest(subject, body)
	return err
}

// CheckEmailLogic 处理邮件检查的业务逻辑
func (s *EmailService) CheckEmailLogic(clientIP string, uuid uuid.UUID) email_response.EmailCheckResult {
	// 检查Redis客户端是否初始化
	if global.GVA_REDIS == nil {
		global.GVA_LOG.Error("Redis客户端未初始化")
		return email_response.EmailCheckResult{Success: false, Message: "服务配置错误"}
	}
	// 创建跟踪IP邮件测试次数的Redis键
	rateLimitKey := fmt.Sprintf("email_test_rate_limit:%s", clientIP)
	// 检查当前尝试次数
	ctx := context.Background()
	currentAttempts, err := global.GVA_REDIS.Get(ctx, rateLimitKey).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		global.GVA_LOG.Error("查询Redis出错!", zap.Error(err))
		return email_response.EmailCheckResult{Success: false, Message: "服务错误"}
	}
	// 如果尝试次数超过3次，返回限流错误
	if currentAttempts >= 3 {
		ttl, _ := global.GVA_REDIS.TTL(ctx, rateLimitKey).Result()
		return email_response.EmailCheckResult{
			Success: false,
			Message: fmt.Sprintf("超过测试限制，请在 %d 分钟后重试", int(ttl.Minutes())+1),
		}
	}
	// 获取邮件推送配置
	var pushers []privmsg.PusherConfig
	err = global.GVA_DB.Where("push_type = ?", "email").Find(&pushers).Error
	if err != nil {
		global.GVA_LOG.Error("查询配置时出错!", zap.Error(err))
		return email_response.EmailCheckResult{Success: false, Message: "查询配置时出错"}
	}
	// 发送邮件
	var user ecsusers.EcsUsers
	result := global.GVA_DB.Model(&ecsusers.EcsUsers{}).Where("uuid = ?", uuid.String()).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			global.GVA_LOG.Error("未找到对应的用户", zap.String("uuid", uuid.String()))
			return email_response.EmailCheckResult{Success: false, Message: "未找到对应的用户"}
		}
		global.GVA_LOG.Error("查询配置时出错!", zap.Error(result.Error))
		return email_response.EmailCheckResult{Success: false, Message: "查询配置时出错"}
	}
	if user.Email == "" {
		return email_response.EmailCheckResult{Success: false, Message: "用户未配置Email，无法发信"}
	}
	var emailSent bool
	for _, pusher := range pushers {
		// 配置值格式：域名地址:端口:发件邮件:密码
		values := strings.Split(pusher.ConfigValue, ":")
		if len(values) < 4 {
			global.GVA_LOG.Error("邮件配置格式错误")
			continue
		}
		smtpHost, smtpPortString, from, password := values[0], values[1], values[2], values[3]
		smtpPort, err := strconv.Atoi(smtpPortString)
		if err != nil {
			global.GVA_LOG.Error("smtp端口转换失败", zap.Error(err))
			continue
		}
		err = s.SendEmailWithSMTPDetails(
			smtpHost,
			smtpPort,
			from,
			password,
			user.Email,
			"订阅成功",
			"恭喜你订阅成功，本邮件无需回复。",
		)
		if err != nil {
			global.GVA_LOG.Error("发送邮件失败!", zap.Error(err))
			continue
		}
		emailSent = true
		break
	}
	// 如果没有成功发送邮件
	if !emailSent {
		return email_response.EmailCheckResult{Success: false, Message: "邮件发送失败，请检查配置"}
	}
	// 在Redis中增加尝试次数
	err = global.GVA_REDIS.Incr(ctx, rateLimitKey).Err()
	if err != nil {
		global.GVA_LOG.Error("Redis计数器增加失败!", zap.Error(err))
	}
	// 如果是第一次尝试，设置15分钟过期
	if currentAttempts == 0 {
		err = global.GVA_REDIS.Expire(ctx, rateLimitKey, 15*time.Minute).Err()
		if err != nil {
			global.GVA_LOG.Error("设置Redis过期时间失败!", zap.Error(err))
		}
	}
	return email_response.EmailCheckResult{Success: true, Message: "发送成功"}
}
