import service from '@/utils/request'

// @Summary 用户的TG验证码获取
// @Produce  application/json
// @Param data body {tgid: "string"}
// @Router /client/code [post]
export const TGRGetCode = (data) => {
  return service({
    url: '/client/code',
    method: 'post',
    data: data,
  })
}

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string", password:"string", tgid:"string", code:"string", captcha:"string", captchaId:"string"}
// @Router /client/client [post]
export const TGRRegister = (data) => {
  return service({
    url: '/client/client',
    method: 'post',
    data: data,
  })
}

// @Summary 密码修改
// @Produce  application/json
// @Param data body {tgid:"string", code:"string", password:"string", new_password:"string"}
// @Router /client/client [post]
export const TGRChangePassword = (data) => {
  return service({
    url: '/client/changePassword',
    method: 'post',
    data: data,
  })
}

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string", password:"string", captcha:"string", captchaId:"string"}
// @Router /client/login [post]
export const TGRLogin = (data) => {
  return service({
    url: '/client/login',
    method: 'post',
    data: data,
  })
}
