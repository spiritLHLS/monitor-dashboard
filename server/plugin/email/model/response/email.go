package response

type Email struct {
	To      string `json:"to"`      // 邮件发送给谁
	Subject string `json:"subject"` // 邮件标题
	Body    string `json:"body"`    // 邮件内容
}

type EmailCheckResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
