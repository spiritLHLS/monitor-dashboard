import service from '@/utils/request'
// @Tags EcsUsers
// @Summary 创建订阅用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EcsUsers true "创建订阅用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /eusr/createEcsUsers [post]
export const createEcsUsers = (data) => {
  return service({
    url: '/eusr/createEcsUsers',
    method: 'post',
    data
  })
}

// @Tags EcsUsers
// @Summary 删除订阅用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EcsUsers true "删除订阅用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eusr/deleteEcsUsers [delete]
export const deleteEcsUsers = (params) => {
  return service({
    url: '/eusr/deleteEcsUsers',
    method: 'delete',
    params
  })
}

// @Tags EcsUsers
// @Summary 批量删除订阅用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订阅用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eusr/deleteEcsUsers [delete]
export const deleteEcsUsersByIds = (params) => {
  return service({
    url: '/eusr/deleteEcsUsersByIds',
    method: 'delete',
    params
  })
}

// @Tags EcsUsers
// @Summary 更新订阅用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EcsUsers true "更新订阅用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eusr/updateEcsUsers [put]
export const updateEcsUsers = (data) => {
  return service({
    url: '/eusr/updateEcsUsers',
    method: 'put',
    data
  })
}

// @Tags EcsUsers
// @Summary 用id查询订阅用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EcsUsers true "用id查询订阅用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eusr/findEcsUsers [get]
export const findEcsUsers = (params) => {
  return service({
    url: '/eusr/findEcsUsers',
    method: 'get',
    params
  })
}

// @Tags EcsUsers
// @Summary 分页获取订阅用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订阅用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eusr/getEcsUsersList [get]
export const getEcsUsersList = (params) => {
  return service({
    url: '/eusr/getEcsUsersList',
    method: 'get',
    params
  })
}

// @Tags EcsUsers
// @Summary 不需要鉴权的订阅用户接口
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "分页获取订阅用户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /eusr/getEcsUsersPublic [get]
export const getEcsUsersPublic = () => {
  return service({
    url: '/eusr/getEcsUsersPublic',
    method: 'get',
  })
}

// AdminChangePassword 方法介绍
// @Tags EcsUsers
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Param data body ecsusersReq.EcsUsersSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/adminChangePassword [PUT]
export const adminChangePassword = (data) => {
  return service({
    url: '/eusr/adminChangePassword',
    method: 'PUT',
    data
  })
}

// selfGetUserInfo 方法介绍
// @Tags EcsUsers
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/selfGetUserInfo [GET]
export const selfGetUserInfo = () => {
  return service({
    url: '/eusr/selfGetUserInfo',
    method: 'GET'
  })
}

// SelfModifyInfo 仅认证用户修改自己的信息
// @Tags EcsUsers
// @Summary 仅认证用户修改自己的信息
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/selfModifyInfo [POST]
export const selfModifyInfo = (data) => {
  return service({
    url: '/eusr/selfModifyInfo',
    method: 'POST',
    data
  })
}