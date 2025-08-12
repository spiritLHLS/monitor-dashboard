package digitalproducts

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/digitalproducts"
	digitalproductsReq "github.com/flipped-aurora/gin-vue-admin/server/model/digitalproducts/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DigitalProductsApi struct{}

// CreateDigitalProducts 创建数字商品表
// @Tags DigitalProducts
// @Summary 创建数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body digitalproducts.DigitalProducts true "创建数字商品表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dpd/createDigitalProducts [post]
func (dpdApi *DigitalProductsApi) CreateDigitalProducts(c *gin.Context) {
	var dpd digitalproducts.DigitalProducts
	err := c.ShouldBindJSON(&dpd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dpd.CreatedBy = utils.GetUserID(c)
	err = dpdService.CreateDigitalProducts(&dpd)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteDigitalProducts 删除数字商品表
// @Tags DigitalProducts
// @Summary 删除数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body digitalproducts.DigitalProducts true "删除数字商品表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dpd/deleteDigitalProducts [delete]
func (dpdApi *DigitalProductsApi) DeleteDigitalProducts(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := dpdService.DeleteDigitalProducts(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDigitalProductsByIds 批量删除数字商品表
// @Tags DigitalProducts
// @Summary 批量删除数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dpd/deleteDigitalProductsByIds [delete]
func (dpdApi *DigitalProductsApi) DeleteDigitalProductsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := dpdService.DeleteDigitalProductsByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDigitalProducts 更新数字商品表
// @Tags DigitalProducts
// @Summary 更新数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body digitalproducts.DigitalProducts true "更新数字商品表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dpd/updateDigitalProducts [put]
func (dpdApi *DigitalProductsApi) UpdateDigitalProducts(c *gin.Context) {
	var dpd digitalproducts.DigitalProducts
	err := c.ShouldBindJSON(&dpd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dpd.UpdatedBy = utils.GetUserID(c)
	err = dpdService.UpdateDigitalProducts(dpd)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDigitalProducts 用id查询数字商品表
// @Tags DigitalProducts
// @Summary 用id查询数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询数字商品表"
// @Success 200 {object} response.Response{data=digitalproducts.DigitalProducts,msg=string} "查询成功"
// @Router /dpd/findDigitalProducts [get]
func (dpdApi *DigitalProductsApi) FindDigitalProducts(c *gin.Context) {
	ID := c.Query("ID")
	redpd, err := dpdService.GetDigitalProducts(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(redpd, c)
}

// GetDigitalProductsList 分页获取数字商品表列表
// @Tags DigitalProducts
// @Summary 分页获取数字商品表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query digitalproductsReq.DigitalProductsSearch true "分页获取数字商品表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dpd/getDigitalProductsList [get]
func (dpdApi *DigitalProductsApi) GetDigitalProductsList(c *gin.Context) {
	var pageInfo digitalproductsReq.DigitalProductsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dpdService.GetDigitalProductsInfoList(pageInfo)
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

// GetDigitalProductsPublic 不需要鉴权的数字商品表接口
// @Tags DigitalProducts
// @Summary 不需要鉴权的数字商品表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dpd/getDigitalProductsPublic [get]
func (dpdApi *DigitalProductsApi) GetDigitalProductsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	dpdService.GetDigitalProductsPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的数字商品表接口信息",
	}, "获取成功", c)
}

// ConvertProductsToDigital 将products表数据转换为数字商品表
// @Tags DigitalProducts
// @Summary 将products表数据转换为数字商品表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param productIds body []uint true "需要转换的产品ID列表"
// @Success 200 {object} response.Response{msg=string} "转换成功"
// @Router /dpd/convertProductsToDigital [post]
func (dpdApi *DigitalProductsApi) ConvertProductsToDigital(c *gin.Context) {
	var productIds []uint
	err := c.ShouldBindJSON(&productIds)
	if err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}
	if len(productIds) == 0 {
		response.FailWithMessage("产品ID列表不能为空", c)
		return
	}
	userID := utils.GetUserID(c)
	successCount, failCount, err := dpdService.ConvertProductsToDigital(productIds, userID)
	if err != nil {
		global.GVA_LOG.Error("转换失败!", zap.Error(err))
		response.FailWithMessage("转换失败: "+err.Error(), c)
		return
	}
	message := fmt.Sprintf("转换完成，成功: %d 条，失败: %d 条", successCount, failCount)
	response.OkWithMessage(message, c)
}
