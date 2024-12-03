package service

import (
	"crypto/tls"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
	"net/smtp"
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
