import service from '@/utils/request'
// @Tags DigitalProducts
// @Summary 创建数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DigitalProducts true "创建数字商品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dpd/createDigitalProducts [post]
export const createDigitalProducts = (data) => {
  return service({
    url: '/dpd/createDigitalProducts',
    method: 'post',
    data
  })
}

// @Tags DigitalProducts
// @Summary 删除数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DigitalProducts true "删除数字商品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dpd/deleteDigitalProducts [delete]
export const deleteDigitalProducts = (params) => {
  return service({
    url: '/dpd/deleteDigitalProducts',
    method: 'delete',
    params
  })
}

// @Tags DigitalProducts
// @Summary 批量删除数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数字商品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dpd/deleteDigitalProducts [delete]
export const deleteDigitalProductsByIds = (params) => {
  return service({
    url: '/dpd/deleteDigitalProductsByIds',
    method: 'delete',
    params
  })
}

// @Tags DigitalProducts
// @Summary 更新数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DigitalProducts true "更新数字商品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dpd/updateDigitalProducts [put]
export const updateDigitalProducts = (data) => {
  return service({
    url: '/dpd/updateDigitalProducts',
    method: 'put',
    data
  })
}

// @Tags DigitalProducts
// @Summary 用id查询数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DigitalProducts true "用id查询数字商品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dpd/findDigitalProducts [get]
export const findDigitalProducts = (params) => {
  return service({
    url: '/dpd/findDigitalProducts',
    method: 'get',
    params
  })
}

// @Tags DigitalProducts
// @Summary 分页获取数字商品表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数字商品表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dpd/getDigitalProductsList [get]
export const getDigitalProductsList = (params) => {
  return service({
    url: '/dpd/getDigitalProductsList',
    method: 'get',
    params
  })
}

// @Tags DigitalProducts
// @Summary 不需要鉴权的数字商品表接口
// @Accept application/json
// @Produce application/json
// @Param data query digitalproductsReq.DigitalProductsSearch true "分页获取数字商品表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dpd/getDigitalProductsPublic [get]
export const getDigitalProductsPublic = () => {
  return service({
    url: '/dpd/getDigitalProductsPublic',
    method: 'get',
  })
}

// @Tags DigitalProducts
// @Summary 将products表数据转换为数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "转换成功"
// @Router /dpd/convertProductsToDigital [get]
export const convertProductsToDigital = () => {
  return service({
    url: '/dpd/convertProductsToDigital',
    method: 'get',
  })
}