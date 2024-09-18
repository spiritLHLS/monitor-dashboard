import service from '@/utils/request'
// @Tags EncryptedLink
// @Summary 创建加密链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EncryptedLink true "创建加密链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /EL/createEncryptedLink [post]
export const createEncryptedLink = (data) => {
  return service({
    url: '/EL/createEncryptedLink',
    method: 'post',
    data
  })
}

// @Tags EncryptedLink
// @Summary 删除加密链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EncryptedLink true "删除加密链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EL/deleteEncryptedLink [delete]
export const deleteEncryptedLink = (params) => {
  return service({
    url: '/EL/deleteEncryptedLink',
    method: 'delete',
    params
  })
}

// @Tags EncryptedLink
// @Summary 批量删除加密链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除加密链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EL/deleteEncryptedLink [delete]
export const deleteEncryptedLinkByIds = (params) => {
  return service({
    url: '/EL/deleteEncryptedLinkByIds',
    method: 'delete',
    params
  })
}

// @Tags EncryptedLink
// @Summary 更新加密链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EncryptedLink true "更新加密链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /EL/updateEncryptedLink [put]
export const updateEncryptedLink = (data) => {
  return service({
    url: '/EL/updateEncryptedLink',
    method: 'put',
    data
  })
}

// @Tags EncryptedLink
// @Summary 用id查询加密链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EncryptedLink true "用id查询加密链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /EL/findEncryptedLink [get]
export const findEncryptedLink = (params) => {
  return service({
    url: '/EL/findEncryptedLink',
    method: 'get',
    params
  })
}

// @Tags EncryptedLink
// @Summary 分页获取加密链接列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取加密链接列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EL/getEncryptedLinkList [get]
export const getEncryptedLinkList = (params) => {
  return service({
    url: '/EL/getEncryptedLinkList',
    method: 'get',
    params
  })
}

// BuildEncryptedUrl 生成加密链接
// @Tags EncryptedLink
// @Summary 生成加密链接
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EL/buildEncryptedUrl [POST]
export const buildEncryptedUrl = (data) => {
  return service({
    url: '/EL/buildEncryptedUrl',
    method: 'POST',
    data
  })
}

// HandleRedirect 处理短代码对应重定向
// @Tags EncryptedLink
// @Summary 处理短代码对应重定向
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EL/handleRedirect [GET]
export const handleRedirect = async (params) => {
  const response = await service({
    url: '/EL/handleRedirect',
    method: 'GET',
    params
  })
  if (response.code === 0 && response.data) {
    return response.data
  } else {
    throw new Error(response.message || '请求失败')
  }
}

// @Tags EncryptedLink
// @Summary 不需要鉴权的加密链接接口
// @accept application/json
// @Produce application/json
// @Param data query request.EncryptedLinkSearch true "分页获取加密链接列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EL/getEncryptedLinkPublic [get]
export const getEncryptedLinkPublic = () => {
  return service({
    url: '/EL/getEncryptedLinkPublic',
    method: 'get',
  })
}