package subscribe

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/subscribe"
    subscribeReq "github.com/flipped-aurora/gin-vue-admin/server/model/subscribe/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SubscribeApi struct {}



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
		response.FailWithMessage("创建失败:" + err.Error(), c)
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
		response.FailWithMessage("删除失败:" + err.Error(), c)
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
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
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
		response.FailWithMessage("更新失败:" + err.Error(), c)
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
		response.FailWithMessage("查询失败:" + err.Error(), c)
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
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetSubscribePublic 不需要鉴权的订阅接口
// @Tags Subscribe
// @Summary 不需要鉴权的订阅接口
// @accept application/json
// @Produce application/json
// @Param data query subscribeReq.SubscribeSearch true "分页获取订阅列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sub/getSubscribePublic [get]
func (subApi *SubscribeApi) GetSubscribePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    subService.GetSubscribePublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的订阅接口信息",
    }, "获取成功", c)
}
// SelfCreateSub 仅当前用户创建当前用户关联的商品推送记录
// @Tags Subscribe
// @Summary 仅当前用户创建当前用户关联的商品推送记录
// @accept application/json
// @Produce application/json
// @Param data query subscribeReq.SubscribeSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /sub/selfCreateSub [POST]
func (subApi *SubscribeApi)SelfCreateSub(c *gin.Context) {
    // 请添加自己的业务逻辑
    err := subService.SelfCreateSub()
    if err != nil {
        global.GVA_LOG.Error("失败!", zap.Error(err))
   		response.FailWithMessage("失败", c)
   		return
   	}
   	response.OkWithData("返回数据",c)
}

// DeleteSubscribePublic 前端用户删除自己已订阅的商品
// @Tags Subscribe
// @Summary 前端用户删除自己已订阅的商品
// @accept application/json
// @Produce application/json
// @Param data query subscribeReq.SubscribeSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /sub/deleteSubscribePublic [POST]
func (subApi *SubscribeApi)DeleteSubscribePublic(c *gin.Context) {
    // 请添加自己的业务逻辑
    err := subService.DeleteSubscribePublic()
    if err != nil {
        global.GVA_LOG.Error("失败!", zap.Error(err))
   		response.FailWithMessage("失败", c)
   		return
   	}
   	response.OkWithData("返回数据",c)
}

