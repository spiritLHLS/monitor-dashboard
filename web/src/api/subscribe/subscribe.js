import service from '@/utils/request'
// @Tags Subscribe
// @Summary 创建订阅
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Subscribe true "创建订阅"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sub/createSubscribe [post]
export const createSubscribe = (data) => {
  return service({
    url: '/sub/createSubscribe',
    method: 'post',
    data
  })
}

// @Tags Subscribe
// @Summary 删除订阅
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Subscribe true "删除订阅"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sub/deleteSubscribe [delete]
export const deleteSubscribe = (params) => {
  return service({
    url: '/sub/deleteSubscribe',
    method: 'delete',
    params
  })
}

// @Tags Subscribe
// @Summary 批量删除订阅
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订阅"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sub/deleteSubscribe [delete]
export const deleteSubscribeByIds = (params) => {
  return service({
    url: '/sub/deleteSubscribeByIds',
    method: 'delete',
    params
  })
}

// @Tags Subscribe
// @Summary 更新订阅
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Subscribe true "更新订阅"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sub/updateSubscribe [put]
export const updateSubscribe = (data) => {
  return service({
    url: '/sub/updateSubscribe',
    method: 'put',
    data
  })
}

// @Tags Subscribe
// @Summary 用id查询订阅
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Subscribe true "用id查询订阅"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sub/findSubscribe [get]
export const findSubscribe = (params) => {
  return service({
    url: '/sub/findSubscribe',
    method: 'get',
    params
  })
}

// @Tags Subscribe
// @Summary 分页获取订阅列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订阅列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sub/getSubscribeList [get]
export const getSubscribeList = (params) => {
  return service({
    url: '/sub/getSubscribeList',
    method: 'get',
    params
  })
}

// @Tags Subscribe
// @Summary 不需要鉴权的订阅接口
// @accept application/json
// @Produce application/json
// @Param data query subscribeReq.SubscribeSearch true "分页获取订阅列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sub/getSubscribePublic [get]
export const getSubscribePublic = () => {
  return service({
    url: '/sub/getSubscribePublic',
    method: 'get',
  })
}
// SelfCreateSub 仅当前用户创建当前用户关联的商品推送记录
// @Tags Subscribe
// @Summary 仅当前用户创建当前用户关联的商品推送记录
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /sub/selfCreateSub [POST]
export const selfCreateSub = () => {
  return service({
    url: '/sub/selfCreateSub',
    method: 'POST'
  })
}
// DeleteSubscribePublic 前端用户删除自己已订阅的商品
// @Tags Subscribe
// @Summary 前端用户删除自己已订阅的商品
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /sub/deleteSubscribePublic [POST]
export const deleteSubscribePublic = () => {
  return service({
    url: '/sub/deleteSubscribePublic',
    method: 'POST'
  })
}
