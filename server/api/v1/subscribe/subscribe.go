package subscribe

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	productsReq "github.com/flipped-aurora/gin-vue-admin/server/model/products/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/subscribe"
	subscribeReq "github.com/flipped-aurora/gin-vue-admin/server/model/subscribe/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SubscribeApi struct{}

// CreateSubscribe 创建订阅
// @Tags Subscribe
// @Summary 创建订阅
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body subscribe.Subscribe true "创建订阅"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sub/createSubscribe [post]
func (subApi *SubscribeApi) CreateSubscribe(c *gin.Context) {
	var sub subscribe.Subscribe
	err := c.ShouldBindJSON(&sub)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = subService.CreateSubscribe(&sub)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSubscribe 删除订阅
// @Tags Subscribe
// @Summary 删除订阅
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body subscribe.Subscribe true "删除订阅"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sub/deleteSubscribe [delete]
func (subApi *SubscribeApi) DeleteSubscribe(c *gin.Context) {
	ID := c.Query("ID")
	err := subService.DeleteSubscribe(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSubscribeByIds 批量删除订阅
// @Tags Subscribe
// @Summary 批量删除订阅
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sub/deleteSubscribeByIds [delete]
func (subApi *SubscribeApi) DeleteSubscribeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := subService.DeleteSubscribeByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSubscribe 更新订阅
// @Tags Subscribe
// @Summary 更新订阅
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body subscribe.Subscribe true "更新订阅"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sub/updateSubscribe [put]
func (subApi *SubscribeApi) UpdateSubscribe(c *gin.Context) {
	var sub subscribe.Subscribe
	err := c.ShouldBindJSON(&sub)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = subService.UpdateSubscribe(sub)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSubscribe 用id查询订阅
// @Tags Subscribe
// @Summary 用id查询订阅
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query subscribe.Subscribe true "用id查询订阅"
// @Success 200 {object} response.Response{data=subscribe.Subscribe,msg=string} "查询成功"
// @Router /sub/findSubscribe [get]
func (subApi *SubscribeApi) FindSubscribe(c *gin.Context) {
	ID := c.Query("ID")
	resub, err := subService.GetSubscribe(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resub, c)
}

// GetSubscribeList 分页获取订阅列表
// @Tags Subscribe
// @Summary 分页获取订阅列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query subscribeReq.SubscribeSearch true "分页获取订阅列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sub/getSubscribeList [get]
func (subApi *SubscribeApi) GetSubscribeList(c *gin.Context) {
	var pageInfo subscribeReq.SubscribeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := subService.GetSubscribeInfoList(pageInfo)
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

// SelfGetSub 前端用户获取自己已订阅的商品
// @Tags Subscribe
// @Summary 不需要鉴权的订阅接口
// @accept application/json
// @Produce application/json
// @Param data query subscribeReq.SubscribeSearch true "分页获取订阅列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sub/selfGetSub [get]
func (subApi *SubscribeApi) SelfGetSub(c *gin.Context) {
	var searchInfo subscribeReq.SubscribeSearch
	err := c.ShouldBindQuery(&searchInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uuid := utils.GetUserUuid(c)
	list, total, err := subService.SelfGetSub(uuid, searchInfo)
	if err != nil {
		global.GVA_LOG.Error("获取订阅失败!", zap.Error(err))
		response.FailWithMessage("获取订阅失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     searchInfo.Page,
		PageSize: searchInfo.PageSize,
	}, "获取成功", c)
}

// SelfCreateSub 前端用户创建关联的商品推送记录
// @Tags Subscribe
// @Summary 仅当前用户创建当前用户关联的商品推送记录
// @accept application/json
// @Produce application/json
// @Param data body subscribeReq.CreateSubscribeRequest true "创建订阅请求"
// @Success 200 {object} response.Response{msg=string} "成功"
// @Router /sub/selfCreateSub [POST]
func (subApi *SubscribeApi) SelfCreateSub(c *gin.Context) {
	var createReq subscribeReq.SubscribeSearch
	err := c.ShouldBindJSON(&createReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uuid := utils.GetUserUuid(c)
	if createReq.ProductId == nil {
		response.FailWithMessage("ProductId不能为空", c)
		return
	}
	if uuid.String() == "" {
		response.FailWithMessage("用户的UUID不能为空", c)
		return
	}
	if createReq.NotifyChannel == "" {
		response.FailWithMessage("商品对应的推送渠道不能为空", c)
		return
	}
	err = subService.SelfCreateSub(uuid, *createReq.ProductId, createReq.NotifyChannel)
	if err != nil {
		global.GVA_LOG.Error("创建订阅失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("创建订阅失败: %s", err.Error()), c)
		return
	}
	response.OkWithMessage("创建订阅成功", c)
}

// SelfDeleteSub 前端用户删除自己已订阅的商品
// @Tags Subscribe
// @Summary 前端用户删除自己已订阅的商品
// @accept application/json
// @Produce application/json
// @Param data body subscribeReq.DeleteSubscribeRequest true "删除订阅请求"
// @Success 200 {object} response.Response{msg=string} "成功"
// @Router /sub/selfDeleteSub [POST]
func (subApi *SubscribeApi) SelfDeleteSub(c *gin.Context) {
	var deleteReq subscribeReq.SubscribeSearch
	err := c.ShouldBindJSON(&deleteReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uuid := utils.GetUserUuid(c)
	if deleteReq.ProductId == nil {
		response.FailWithMessage("ProductId不能为空", c)
		return
	}
	err = subService.SelfDeleteSub(uuid, *deleteReq.ProductId)
	if err != nil {
		global.GVA_LOG.Error("删除订阅失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("删除订阅失败: %s", err.Error()), c)
		return
	}
	response.OkWithMessage("删除订阅成功", c)
}

// SelfUpdateSub 前端用户更新自己已订阅的商品信息
// @Tags Subscribe
// @Summary 前端用户更新自己已订阅的商品信息
// @accept application/json
// @Produce application/json
// @Param data body subscribeReq.UpdateSubscribeRequest true "更新订阅请求"
// @Success 200 {object} response.Response{msg=string} "成功"
// @Router /sub/selfUpdateSub [POST]
func (subApi *SubscribeApi) SelfUpdateSub(c *gin.Context) {
	var updateReq subscribeReq.SubscribeSearch
	err := c.ShouldBindJSON(&updateReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uuid := utils.GetUserUuid(c)
	if uuid.String() == "" {
		response.FailWithMessage("用户的uuid不能为空", c)
		return
	}
	if updateReq.ProductId == nil {
		response.FailWithMessage("ProductId不能为空", c)
		return
	}
	err = subService.SelfUpdateSub(uuid, *updateReq.ProductId, updateReq.NotifyChannel)
	if err != nil {
		global.GVA_LOG.Error("更新订阅失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("更新订阅失败: %s", err.Error()), c)
		return
	}
	response.OkWithMessage("更新订阅成功", c)
}

// SelfGetAllPd 前端用户获取所有的商品
// @Tags Subscribe
// @Summary 前端用户获取所有的商品
// @accept application/json
// @Produce application/json
// @Param data query productsReq.ProductsSearch true "分页获取商品列表"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /sub/selfGetAllPd [GET]
func (subApi *SubscribeApi) SelfGetAllPd(c *gin.Context) {
	var searchInfo productsReq.ProductsSearch
	err := c.ShouldBindQuery(&searchInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := subService.SelfGetAllPd(searchInfo)
	if err != nil {
		global.GVA_LOG.Error("获取商品列表失败!", zap.Error(err))
		response.FailWithMessage("获取商品列表失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     searchInfo.Page,
		PageSize: searchInfo.PageSize,
	}, "获取成功", c)
}

// SelfBatchUpdateStatus 前端用户修改订阅状态
// @Tags Subscribe
// @Summary 前端用户修改订阅状态
// @accept application/json
// @Produce application/json
// @Param data query subscribeReq.SubscribeSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /sub/selfBatchUpdateStatus [POST]
func (subApi *SubscribeApi) SelfBatchUpdateStatus(c *gin.Context) {
	var ids subscribe.UpdateStatusIDs
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uuid := utils.GetUserUuid(c)
	if uuid.String() == "" {
		response.FailWithMessage("用户的uuid不能为空", c)
		return
	}
	err = subService.SelfBatchUpdateStatus(uuid.String(), ids.IDs)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("错误：%s", err.Error()), c)
		return
	}
	response.OkWithData("返回数据", c)
}
