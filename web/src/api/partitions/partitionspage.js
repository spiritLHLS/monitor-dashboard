import service from '@/utils/request'
// @Tags Partitionspage
// @Summary 创建partitionspage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Partitionspage true "创建partitionspage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /psp/createPartitionspage [post]
export const createPartitionspage = (data) => {
  return service({
    url: '/psp/createPartitionspage',
    method: 'post',
    data
  })
}

// @Tags Partitionspage
// @Summary 删除partitionspage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Partitionspage true "删除partitionspage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /psp/deletePartitionspage [delete]
export const deletePartitionspage = (params) => {
  return service({
    url: '/psp/deletePartitionspage',
    method: 'delete',
    params
  })
}

// @Tags Partitionspage
// @Summary 批量删除partitionspage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除partitionspage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /psp/deletePartitionspage [delete]
export const deletePartitionspageByIds = (params) => {
  return service({
    url: '/psp/deletePartitionspageByIds',
    method: 'delete',
    params
  })
}

// @Tags Partitionspage
// @Summary 更新partitionspage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Partitionspage true "更新partitionspage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /psp/updatePartitionspage [put]
export const updatePartitionspage = (data) => {
  return service({
    url: '/psp/updatePartitionspage',
    method: 'put',
    data
  })
}

// @Tags Partitionspage
// @Summary 用id查询partitionspage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Partitionspage true "用id查询partitionspage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /psp/findPartitionspage [get]
export const findPartitionspage = (params) => {
  return service({
    url: '/psp/findPartitionspage',
    method: 'get',
    params
  })
}

// @Tags Partitionspage
// @Summary 分页获取partitionspage表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取partitionspage表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /psp/getPartitionspageList [get]
export const getPartitionspageList = (params) => {
  return service({
    url: '/psp/getPartitionspageList',
    method: 'get',
    params
  })
}

// @Tags Partitionspage
// @Summary 不需要鉴权的partitionspage表接口
// @accept application/json
// @Produce application/json
// @Param data query partitionsReq.PartitionspageSearch true "分页获取partitionspage表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /psp/getPartitionspagePublic [get]
export const getPartitionspagePublic = () => {
  return service({
    url: '/psp/getPartitionspagePublic',
    method: 'get',
  })
}
