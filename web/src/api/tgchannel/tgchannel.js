import service from '@/utils/request'
// @Tags Tgchannel
// @Summary 创建tgchannel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tgchannel true "创建tgchannel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tgc/createTgchannel [post]
export const createTgchannel = (data) => {
  return service({
    url: '/tgc/createTgchannel',
    method: 'post',
    data
  })
}

// @Tags Tgchannel
// @Summary 删除tgchannel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tgchannel true "删除tgchannel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tgc/deleteTgchannel [delete]
export const deleteTgchannel = (params) => {
  return service({
    url: '/tgc/deleteTgchannel',
    method: 'delete',
    params
  })
}

// @Tags Tgchannel
// @Summary 批量删除tgchannel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除tgchannel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tgc/deleteTgchannel [delete]
export const deleteTgchannelByIds = (params) => {
  return service({
    url: '/tgc/deleteTgchannelByIds',
    method: 'delete',
    params
  })
}

// @Tags Tgchannel
// @Summary 更新tgchannel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tgchannel true "更新tgchannel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tgc/updateTgchannel [put]
export const updateTgchannel = (data) => {
  return service({
    url: '/tgc/updateTgchannel',
    method: 'put',
    data
  })
}

// @Tags Tgchannel
// @Summary 用id查询tgchannel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Tgchannel true "用id查询tgchannel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tgc/findTgchannel [get]
export const findTgchannel = (params) => {
  return service({
    url: '/tgc/findTgchannel',
    method: 'get',
    params
  })
}

// @Tags Tgchannel
// @Summary 分页获取tgchannel表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取tgchannel表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tgc/getTgchannelList [get]
export const getTgchannelList = (params) => {
  return service({
    url: '/tgc/getTgchannelList',
    method: 'get',
    params
  })
}

// @Tags Tgchannel
// @Summary 不需要鉴权的tgchannel表接口
// @accept application/json
// @Produce application/json
// @Param data query tgchannelReq.TgchannelSearch true "分页获取tgchannel表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tgc/getTgchannelPublic [get]
export const getTgchannelPublic = () => {
  return service({
    url: '/tgc/getTgchannelPublic',
    method: 'get',
  })
}
