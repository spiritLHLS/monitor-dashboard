package shops

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/shops"
	shopsReq "server/model/shops/request"
)

type ShopsApi struct{}

// CreateShops 创建shops表
// @Tags Shops
// @Summary 创建shops表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shops.Shops true "创建shops表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sp/createShops [post]
func (spApi *ShopsApi) CreateShops(c *gin.Context) {
	var sp shops.Shops
	err := c.ShouldBindJSON(&sp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = spService.CreateShops(&sp)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteShops 删除shops表
// @Tags Shops
// @Summary 删除shops表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shops.Shops true "删除shops表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sp/deleteShops [delete]
func (spApi *ShopsApi) DeleteShops(c *gin.Context) {
	ID := c.Query("ID")
	err := spService.DeleteShops(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteShopsByIds 批量删除shops表
// @Tags Shops
// @Summary 批量删除shops表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sp/deleteShopsByIds [delete]
func (spApi *ShopsApi) DeleteShopsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := spService.DeleteShopsByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateShops 更新shops表
// @Tags Shops
// @Summary 更新shops表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shops.Shops true "更新shops表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sp/updateShops [put]
func (spApi *ShopsApi) UpdateShops(c *gin.Context) {
	var sp shops.Shops
	err := c.ShouldBindJSON(&sp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = spService.UpdateShops(sp)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindShops 用id查询shops表
// @Tags Shops
// @Summary 用id查询shops表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shops.Shops true "用id查询shops表"
// @Success 200 {object} response.Response{data=shops.Shops,msg=string} "查询成功"
// @Router /sp/findShops [get]
func (spApi *ShopsApi) FindShops(c *gin.Context) {
	ID := c.Query("ID")
	resp, err := spService.GetShops(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resp, c)
}

// GetShopsList 分页获取shops表列表
// @Tags Shops
// @Summary 分页获取shops表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopsReq.ShopsSearch true "分页获取shops表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sp/getShopsList [get]
func (spApi *ShopsApi) GetShopsList(c *gin.Context) {
	var pageInfo shopsReq.ShopsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := spService.GetShopsInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetShopsPublic 不需要鉴权的shops表接口
// @Tags Shops
// @Summary 不需要鉴权的shops表接口
// @accept application/json
// @Produce application/json
// @Param data query shopsReq.ShopsSearch true "分页获取shops表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sp/getShopsPublic [get]
func (spApi *ShopsApi) GetShopsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	spService.GetShopsPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的shops表接口信息",
	}, "获取成功", c)
}
