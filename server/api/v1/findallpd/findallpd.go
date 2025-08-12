package findallpd

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/findallpd"
	findallpdReq "server/model/findallpd/request"
)

type FindallpdApi struct{}

// CreateFindallpd 创建findallpd表
// @Tags Findallpd
// @Summary 创建findallpd表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body findallpd.Findallpd true "创建findallpd表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /fapd/createFindallpd [post]
func (fapdApi *FindallpdApi) CreateFindallpd(c *gin.Context) {
	var fapd findallpd.Findallpd
	err := c.ShouldBindJSON(&fapd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = fapdService.CreateFindallpd(&fapd)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteFindallpd 删除findallpd表
// @Tags Findallpd
// @Summary 删除findallpd表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body findallpd.Findallpd true "删除findallpd表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /fapd/deleteFindallpd [delete]
func (fapdApi *FindallpdApi) DeleteFindallpd(c *gin.Context) {
	ID := c.Query("ID")
	err := fapdService.DeleteFindallpd(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteFindallpdByIds 批量删除findallpd表
// @Tags Findallpd
// @Summary 批量删除findallpd表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /fapd/deleteFindallpdByIds [delete]
func (fapdApi *FindallpdApi) DeleteFindallpdByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := fapdService.DeleteFindallpdByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateFindallpd 更新findallpd表
// @Tags Findallpd
// @Summary 更新findallpd表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body findallpd.Findallpd true "更新findallpd表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /fapd/updateFindallpd [put]
func (fapdApi *FindallpdApi) UpdateFindallpd(c *gin.Context) {
	var fapd findallpd.Findallpd
	err := c.ShouldBindJSON(&fapd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = fapdService.UpdateFindallpd(fapd)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindFindallpd 用id查询findallpd表
// @Tags Findallpd
// @Summary 用id查询findallpd表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query findallpd.Findallpd true "用id查询findallpd表"
// @Success 200 {object} response.Response{data=findallpd.Findallpd,msg=string} "查询成功"
// @Router /fapd/findFindallpd [get]
func (fapdApi *FindallpdApi) FindFindallpd(c *gin.Context) {
	ID := c.Query("ID")
	refapd, err := fapdService.GetFindallpd(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(refapd, c)
}

// GetFindallpdList 分页获取findallpd表列表
// @Tags Findallpd
// @Summary 分页获取findallpd表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query findallpdReq.FindallpdSearch true "分页获取findallpd表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /fapd/getFindallpdList [get]
func (fapdApi *FindallpdApi) GetFindallpdList(c *gin.Context) {
	var pageInfo findallpdReq.FindallpdSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := fapdService.GetFindallpdInfoList(pageInfo)
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

// GetFindallpdPublic 不需要鉴权的findallpd表接口
// @Tags Findallpd
// @Summary 不需要鉴权的findallpd表接口
// @accept application/json
// @Produce application/json
// @Param data query findallpdReq.FindallpdSearch true "分页获取findallpd表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /fapd/getFindallpdPublic [get]
func (fapdApi *FindallpdApi) GetFindallpdPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	fapdService.GetFindallpdPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的findallpd表接口信息",
	}, "获取成功", c)
}
