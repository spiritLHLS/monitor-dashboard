import service from '@/utils/request'
// @Tags Findallpd
// @Summary 创建findallpd表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Findallpd true "创建findallpd表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fapd/createFindallpd [post]
export const createFindallpd = (data) => {
  return service({
    url: '/fapd/createFindallpd',
    method: 'post',
    data
  })
}

// @Tags Findallpd
// @Summary 删除findallpd表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Findallpd true "删除findallpd表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fapd/deleteFindallpd [delete]
export const deleteFindallpd = (params) => {
  return service({
    url: '/fapd/deleteFindallpd',
    method: 'delete',
    params
  })
}

// @Tags Findallpd
// @Summary 批量删除findallpd表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除findallpd表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fapd/deleteFindallpd [delete]
export const deleteFindallpdByIds = (params) => {
  return service({
    url: '/fapd/deleteFindallpdByIds',
    method: 'delete',
    params
  })
}

// @Tags Findallpd
// @Summary 更新findallpd表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Findallpd true "更新findallpd表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fapd/updateFindallpd [put]
export const updateFindallpd = (data) => {
  return service({
    url: '/fapd/updateFindallpd',
    method: 'put',
    data
  })
}

// @Tags Findallpd
// @Summary 用id查询findallpd表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Findallpd true "用id查询findallpd表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fapd/findFindallpd [get]
export const findFindallpd = (params) => {
  return service({
    url: '/fapd/findFindallpd',
    method: 'get',
    params
  })
}

// @Tags Findallpd
// @Summary 分页获取findallpd表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取findallpd表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fapd/getFindallpdList [get]
export const getFindallpdList = (params) => {
  return service({
    url: '/fapd/getFindallpdList',
    method: 'get',
    params
  })
}

// @Tags Findallpd
// @Summary 不需要鉴权的findallpd表接口
// @accept application/json
// @Produce application/json
// @Param data query findallpdReq.FindallpdSearch true "分页获取findallpd表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /fapd/getFindallpdPublic [get]
export const getFindallpdPublic = () => {
  return service({
    url: '/fapd/getFindallpdPublic',
    method: 'get',
  })
}
