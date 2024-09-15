import service from '@/utils/request'

// @Summary 用户的TG验证码获取
// @Produce  application/json
// @Param data body {tgid: "string"}
// @Router /register/code [post]
export const TGRGetCode = (data) => {
  return service({
    url: '/register/code',
    method: 'post',
    data: data,
  })
}

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string", password:"string", tgid:"string", code:"string", captcha:"string", captchaId:"string"}
// @Router /register/register [post]
export const TGRRegister = (data) => {
  return service({
    url: '/register/register',
    method: 'post',
    data: data,
  })
}

// @Summary 密码修改
// @Produce  application/json
// @Param data body {tgid:"string", code:"string", password:"string", new_password:"string"}
// @Router /register/register [post]
export const TGRChangePassword = (data) => {
  return service({
    url: '/register/changePassword',
    method: 'post',
    data: data,
  })
}

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string", password:"string", captcha:"string", captchaId:"string"}
// @Router /register/login [post]
export const TGRLogin = (data) => {
  return service({
    url: '/register/login',
    method: 'post',
    data: data,
  })
}
