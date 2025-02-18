import service from '@/utils/request'
// @Tags InviteCodes
// @Summary 创建邀请码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InviteCodes true "创建邀请码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /invcode/createInviteCodes [post]
export const createInviteCodes = (data) => {
  return service({
    url: '/invcode/createInviteCodes',
    method: 'post',
    data
  })
}

// @Tags InviteCodes
// @Summary 删除邀请码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InviteCodes true "删除邀请码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /invcode/deleteInviteCodes [delete]
export const deleteInviteCodes = (params) => {
  return service({
    url: '/invcode/deleteInviteCodes',
    method: 'delete',
    params
  })
}

// @Tags InviteCodes
// @Summary 批量删除邀请码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除邀请码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /invcode/deleteInviteCodes [delete]
export const deleteInviteCodesByIds = (params) => {
  return service({
    url: '/invcode/deleteInviteCodesByIds',
    method: 'delete',
    params
  })
}

// @Tags InviteCodes
// @Summary 更新邀请码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InviteCodes true "更新邀请码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /invcode/updateInviteCodes [put]
export const updateInviteCodes = (data) => {
  return service({
    url: '/invcode/updateInviteCodes',
    method: 'put',
    data
  })
}

// @Tags InviteCodes
// @Summary 用id查询邀请码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InviteCodes true "用id查询邀请码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /invcode/findInviteCodes [get]
export const findInviteCodes = (params) => {
  return service({
    url: '/invcode/findInviteCodes',
    method: 'get',
    params
  })
}

// @Tags InviteCodes
// @Summary 分页获取邀请码列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取邀请码列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /invcode/getInviteCodesList [get]
export const getInviteCodesList = (params) => {
  return service({
    url: '/invcode/getInviteCodesList',
    method: 'get',
    params
  })
}

// @Tags InviteCodes
// @Summary 不需要鉴权的邀请码接口
// @accept application/json
// @Produce application/json
// @Param data query invite_codesReq.InviteCodesSearch true "分页获取邀请码列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /invcode/getInviteCodesPublic [get]
export const getInviteCodesPublic = () => {
  return service({
    url: '/invcode/getInviteCodesPublic',
    method: 'get',
  })
}

// BatchBuildCodes 批量生成邀请码
// @Tags InviteCodes
// @Summary 批量生成邀请码
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /invcode/batchBuildCodes [POST]
export const batchBuildCodes = (data) => {
  return service({
    url: '/invcode/batchBuildCodes',
    method: 'POST',
    data
  })
}

// BatchExportCodes 批量导出邀请码
// @Tags InviteCodes
// @Summary 批量导出邀请码
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /invcode/batchExportCodes [POST]
export const batchExportCodes = (data) => {
  return service({
    url: '/invcode/batchExportCodes',
    method: 'POST',
    data
  })
}

