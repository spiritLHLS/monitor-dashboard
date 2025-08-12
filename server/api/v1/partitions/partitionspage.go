package partitions

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/partitions"
	partitionsReq "server/model/partitions/request"
)

type PartitionspageApi struct{}

// CreatePartitionspage 创建partitionspage表
// @Tags Partitionspage
// @Summary 创建partitionspage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body partitions.Partitionspage true "创建partitionspage表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /psp/createPartitionspage [post]
func (pspApi *PartitionspageApi) CreatePartitionspage(c *gin.Context) {
	var psp partitions.Partitionspage
	err := c.ShouldBindJSON(&psp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = pspService.CreatePartitionspage(&psp)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePartitionspage 删除partitionspage表
// @Tags Partitionspage
// @Summary 删除partitionspage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body partitions.Partitionspage true "删除partitionspage表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /psp/deletePartitionspage [delete]
func (pspApi *PartitionspageApi) DeletePartitionspage(c *gin.Context) {
	ID := c.Query("ID")
	err := pspService.DeletePartitionspage(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePartitionspageByIds 批量删除partitionspage表
// @Tags Partitionspage
// @Summary 批量删除partitionspage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /psp/deletePartitionspageByIds [delete]
func (pspApi *PartitionspageApi) DeletePartitionspageByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := pspService.DeletePartitionspageByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePartitionspage 更新partitionspage表
// @Tags Partitionspage
// @Summary 更新partitionspage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body partitions.Partitionspage true "更新partitionspage表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /psp/updatePartitionspage [put]
func (pspApi *PartitionspageApi) UpdatePartitionspage(c *gin.Context) {
	var psp partitions.Partitionspage
	err := c.ShouldBindJSON(&psp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = pspService.UpdatePartitionspage(psp)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPartitionspage 用id查询partitionspage表
// @Tags Partitionspage
// @Summary 用id查询partitionspage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query partitions.Partitionspage true "用id查询partitionspage表"
// @Success 200 {object} response.Response{data=partitions.Partitionspage,msg=string} "查询成功"
// @Router /psp/findPartitionspage [get]
func (pspApi *PartitionspageApi) FindPartitionspage(c *gin.Context) {
	ID := c.Query("ID")
	repsp, err := pspService.GetPartitionspage(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(repsp, c)
}

// GetPartitionspageList 分页获取partitionspage表列表
// @Tags Partitionspage
// @Summary 分页获取partitionspage表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query partitionsReq.PartitionspageSearch true "分页获取partitionspage表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /psp/getPartitionspageList [get]
func (pspApi *PartitionspageApi) GetPartitionspageList(c *gin.Context) {
	var pageInfo partitionsReq.PartitionspageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pspService.GetPartitionspageInfoList(pageInfo)
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

// GetPartitionspagePublic 不需要鉴权的partitionspage表接口
// @Tags Partitionspage
// @Summary 不需要鉴权的partitionspage表接口
// @accept application/json
// @Produce application/json
// @Param data query partitionsReq.PartitionspageSearch true "分页获取partitionspage表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /psp/getPartitionspagePublic [get]
func (pspApi *PartitionspageApi) GetPartitionspagePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	pspService.GetPartitionspagePublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的partitionspage表接口信息",
	}, "获取成功", c)
}
