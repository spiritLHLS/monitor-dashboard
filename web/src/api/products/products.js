import service from '@/utils/request'
// @Tags Products
// @Summary 创建products表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Products true "创建products表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pd/createProducts [post]
export const createProducts = (data) => {
  return service({
    url: '/pd/createProducts',
    method: 'post',
    data
  })
}

// @Tags Products
// @Summary 删除products表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Products true "删除products表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pd/deleteProducts [delete]
export const deleteProducts = (params) => {
  return service({
    url: '/pd/deleteProducts',
    method: 'delete',
    params
  })
}

// @Tags Products
// @Summary 批量删除products表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除products表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pd/deleteProducts [delete]
export const deleteProductsByIds = (params) => {
  return service({
    url: '/pd/deleteProductsByIds',
    method: 'delete',
    params
  })
}

// @Tags Products
// @Summary 更新products表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Products true "更新products表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pd/updateProducts [put]
export const updateProducts = (data) => {
  return service({
    url: '/pd/updateProducts',
    method: 'put',
    data
  })
}

// @Tags Products
// @Summary 用id查询products表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Products true "用id查询products表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pd/findProducts [get]
export const findProducts = (params) => {
  return service({
    url: '/pd/findProducts',
    method: 'get',
    params
  })
}

// @Tags Products
// @Summary 分页获取products表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取products表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pd/getProductsList [get]
export const getProductsList = (params) => {
  return service({
    url: '/pd/getProductsList',
    method: 'get',
    params
  })
}

// @Tags Products
// @Summary 不需要鉴权的products表接口
// @accept application/json
// @Produce application/json
// @Param data query productsReq.ProductsSearch true "分页获取products表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pd/getProductsPublic [get]
export const getProductsPublic = (params) => {
  return service({
    url: '/pd/getProductsPublic',
    method: 'get',
    params
  })
}