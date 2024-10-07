import service from '@/utils/request'
// @Tags PusherConfig
// @Summary 创建推送配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PusherConfig true "创建推送配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pc/createPusherConfig [post]
export const createPusherConfig = (data) => {
  return service({
    url: '/pc/createPusherConfig',
    method: 'post',
    data
  })
}

// @Tags PusherConfig
// @Summary 删除推送配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PusherConfig true "删除推送配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pc/deletePusherConfig [delete]
export const deletePusherConfig = (params) => {
  return service({
    url: '/pc/deletePusherConfig',
    method: 'delete',
    params
  })
}

// @Tags PusherConfig
// @Summary 批量删除推送配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除推送配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pc/deletePusherConfig [delete]
export const deletePusherConfigByIds = (params) => {
  return service({
    url: '/pc/deletePusherConfigByIds',
    method: 'delete',
    params
  })
}

// @Tags PusherConfig
// @Summary 更新推送配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PusherConfig true "更新推送配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pc/updatePusherConfig [put]
export const updatePusherConfig = (data) => {
  return service({
    url: '/pc/updatePusherConfig',
    method: 'put',
    data
  })
}

// @Tags PusherConfig
// @Summary 用id查询推送配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PusherConfig true "用id查询推送配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pc/findPusherConfig [get]
export const findPusherConfig = (params) => {
  return service({
    url: '/pc/findPusherConfig',
    method: 'get',
    params
  })
}

// @Tags PusherConfig
// @Summary 分页获取推送配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取推送配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pc/getPusherConfigList [get]
export const getPusherConfigList = (params) => {
  return service({
    url: '/pc/getPusherConfigList',
    method: 'get',
    params
  })
}

// @Tags PusherConfig
// @Summary 不需要鉴权的推送配置接口
// @accept application/json
// @Produce application/json
// @Param data query privmsgReq.PusherConfigSearch true "分页获取推送配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pc/getPusherConfigPublic [get]
export const getPusherConfigPublic = () => {
  return service({
    url: '/pc/getPusherConfigPublic',
    method: 'get',
  })
}