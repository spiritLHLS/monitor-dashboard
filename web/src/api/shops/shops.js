import service from '@/utils/request'
// @Tags Shops
// @Summary 创建shops表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shops true "创建shops表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sp/createShops [post]
export const createShops = (data) => {
  return service({
    url: '/sp/createShops',
    method: 'post',
    data
  })
}

// @Tags Shops
// @Summary 删除shops表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shops true "删除shops表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sp/deleteShops [delete]
export const deleteShops = (params) => {
  return service({
    url: '/sp/deleteShops',
    method: 'delete',
    params
  })
}

// @Tags Shops
// @Summary 批量删除shops表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除shops表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sp/deleteShops [delete]
export const deleteShopsByIds = (params) => {
  return service({
    url: '/sp/deleteShopsByIds',
    method: 'delete',
    params
  })
}

// @Tags Shops
// @Summary 更新shops表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shops true "更新shops表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sp/updateShops [put]
export const updateShops = (data) => {
  return service({
    url: '/sp/updateShops',
    method: 'put',
    data
  })
}

// @Tags Shops
// @Summary 用id查询shops表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Shops true "用id查询shops表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sp/findShops [get]
export const findShops = (params) => {
  return service({
    url: '/sp/findShops',
    method: 'get',
    params
  })
}

// @Tags Shops
// @Summary 分页获取shops表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取shops表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sp/getShopsList [get]
export const getShopsList = (params) => {
  return service({
    url: '/sp/getShopsList',
    method: 'get',
    params
  })
}

// @Tags Shops
// @Summary 不需要鉴权的shops表接口
// @accept application/json
// @Produce application/json
// @Param data query shopsReq.ShopsSearch true "分页获取shops表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sp/getShopsPublic [get]
export const getShopsPublic = () => {
  return service({
    url: '/sp/getShopsPublic',
    method: 'get',
  })
}
