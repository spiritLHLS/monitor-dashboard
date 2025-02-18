import service from '@/utils/request'


// CheckEmail 检查并发送邮件的接口
// @Tags      System
// @Summary   测试邮件推送是否成功
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      email_response.Email  true  "发送邮件必须的参数"
// @Success   200   {string}  string                "{"success":true,"data":{},"msg":"发送成功"}"
// @Router    /email/checkEmail [post]
export const CheckEmail = () => {
  return service({
    url: '/email/checkEmail',
    method: 'post'
  })
}

// CheckTgBot
// @Tags Telegram_bot
// @Summary 检测TG推送是否可用
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param    token user_id channel_id  "检测TG推送是否可用的必要参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /telegram_bot/checkTgBot [post]
export const CheckTgBot = () => {
    return service({
      url: '/telegram_bot/checkTgBot',
      method: 'post'
    })
  }