package products

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/products"
	productsReq "server/model/products/request"
)

type ProductsApi struct{}

// CreateProducts 创建products表
// @Tags Products
// @Summary 创建products表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body products.Products true "创建products表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /pd/createProducts [post]
func (pdApi *ProductsApi) CreateProducts(c *gin.Context) {
	var pd products.Products
	err := c.ShouldBindJSON(&pd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = pdService.CreateProducts(&pd)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteProducts 删除products表
// @Tags Products
// @Summary 删除products表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body products.Products true "删除products表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /pd/deleteProducts [delete]
func (pdApi *ProductsApi) DeleteProducts(c *gin.Context) {
	ID := c.Query("ID")
	err := pdService.DeleteProducts(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProductsByIds 批量删除products表
// @Tags Products
// @Summary 批量删除products表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /pd/deleteProductsByIds [delete]
func (pdApi *ProductsApi) DeleteProductsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := pdService.DeleteProductsByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProducts 更新products表
// @Tags Products
// @Summary 更新products表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body products.Products true "更新products表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /pd/updateProducts [put]
func (pdApi *ProductsApi) UpdateProducts(c *gin.Context) {
	var pd products.Products
	err := c.ShouldBindJSON(&pd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = pdService.UpdateProducts(pd)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindProducts 用id查询products表
// @Tags Products
// @Summary 用id查询products表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query products.Products true "用id查询products表"
// @Success 200 {object} response.Response{data=products.Products,msg=string} "查询成功"
// @Router /pd/findProducts [get]
func (pdApi *ProductsApi) FindProducts(c *gin.Context) {
	ID := c.Query("ID")
	repd, err := pdService.GetProducts(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(repd, c)
}

// GetProductsList 分页获取products表列表
// @Tags Products
// @Summary 分页获取products表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query productsReq.ProductsSearch true "分页获取products表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /pd/getProductsList [get]
func (pdApi *ProductsApi) GetProductsList(c *gin.Context) {
	var pageInfo productsReq.ProductsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pdService.GetProductsInfoList(pageInfo)
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

// GetProductsPublic 不需要鉴权的products表接口
// @Tags Products
// @Summary 不需要鉴权的products表接口
// @accept application/json
// @Produce application/json
// @Param data query productsReq.DigitalProductsSearch true "分页获取products表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pd/getProductsPublic [get]
func (pdApi *ProductsApi) GetProductsPublic(c *gin.Context) {
	var pageInfo productsReq.DigitalProductsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pdService.GetProductsPublicDigital(pageInfo)
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
