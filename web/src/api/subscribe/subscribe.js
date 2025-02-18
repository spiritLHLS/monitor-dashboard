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

// @Tags selfGetSub
// @Summary 前端用户获取自己已订阅的商品
// @accept application/json
// @Produce application/json
// @Param data query subscribeReq.SubscribeSearch true "分页获取订阅列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sub/selfGetSub [get]
export const selfGetSub = (params) => {
  return service({
    url: '/sub/selfGetSub',
    method: 'get',
    params
  })
}

// SelfCreateSub 前端用户自行选择订阅商品
// @Tags Subscribe
// @Summary 仅当前用户创建当前用户关联的商品推送记录
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /sub/selfCreateSub [POST]
export const selfCreateSub = (data) => {
  return service({
    url: '/sub/selfCreateSub',
    method: 'POST',
    data
  })
}

// selfDeleteSub 前端用户删除自己已订阅的商品
// @Tags Subscribe
// @Summary 前端用户删除自己已订阅的商品
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /sub/selfDeleteSub [POST]
export const selfDeleteSub = (data) => {
  return service({
    url: '/sub/selfDeleteSub',
    method: 'POST',
    data
  })
}

// SelfGetAllPd 前端用户获取所有的商品
// @Tags Subscribe
// @Summary 前端用户获取所有的商品
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /sub/selfGetAllPd [GET]
export const selfGetAllPd = (params) => {
  return service({
    url: '/sub/selfGetAllPd',
    method: 'GET',
    params
  })
}

// SelfUpdateSub 前端用户更新自己已订阅的商品信息
// @Tags Subscribe
// @Summary 前端用户更新自己已订阅的商品信息
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /sub/selfUpdateSub [POST]
export const selfUpdateSub = (data) => {
  return service({
    url: '/sub/selfUpdateSub',
    method: 'POST',
    data
  })
}

// SelfBatchUpdateStatus 前端用户修改订阅状态
// @Tags Subscribe
// @Summary 前端用户修改订阅状态
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /sub/selfBatchUpdateStatus [POST]
export const selfBatchUpdateStatus = (data) => {
  return service({
    url: '/sub/selfBatchUpdateStatus',
    method: 'POST',
    data
  })
}
