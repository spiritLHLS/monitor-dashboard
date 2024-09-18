package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cryptourl/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cryptourl/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var EncryptedLink = new(EL)

type EL struct{}

// CreateEncryptedLink 创建加密链接
// @Tags EncryptedLink
// @Summary 创建加密链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EncryptedLink true "创建加密链接"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /EL/createEncryptedLink [post]
func (a *EL) CreateEncryptedLink(c *gin.Context) {
	var info model.EncryptedLink
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceEncryptedLink.CreateEncryptedLink(&info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEncryptedLink 删除加密链接
// @Tags EncryptedLink
// @Summary 删除加密链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EncryptedLink true "删除加密链接"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /EL/deleteEncryptedLink [delete]
func (a *EL) DeleteEncryptedLink(c *gin.Context) {
	ID := c.Query("ID")
	err := serviceEncryptedLink.DeleteEncryptedLink(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEncryptedLinkByIds 批量删除加密链接
// @Tags EncryptedLink
// @Summary 批量删除加密链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /EL/deleteEncryptedLinkByIds [delete]
func (a *EL) DeleteEncryptedLinkByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := serviceEncryptedLink.DeleteEncryptedLinkByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEncryptedLink 更新加密链接
// @Tags EncryptedLink
// @Summary 更新加密链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EncryptedLink true "更新加密链接"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /EL/updateEncryptedLink [put]
func (a *EL) UpdateEncryptedLink(c *gin.Context) {
	var info model.EncryptedLink
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceEncryptedLink.UpdateEncryptedLink(info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEncryptedLink 用id查询加密链接
// @Tags EncryptedLink
// @Summary 用id查询加密链接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EncryptedLink true "用id查询加密链接"
// @Success 200 {object} response.Response{data=model.EncryptedLink,msg=string} "查询成功"
// @Router /EL/findEncryptedLink [get]
func (a *EL) FindEncryptedLink(c *gin.Context) {
	ID := c.Query("ID")
	reEL, err := serviceEncryptedLink.GetEncryptedLink(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reEL, c)
}

// GetEncryptedLinkList 分页获取加密链接列表
// @Tags EncryptedLink
// @Summary 分页获取加密链接列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.EncryptedLinkSearch true "分页获取加密链接列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /EL/getEncryptedLinkList [get]
func (a *EL) GetEncryptedLinkList(c *gin.Context) {
	var pageInfo request.EncryptedLinkSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceEncryptedLink.GetEncryptedLinkInfoList(pageInfo)
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

// GetEncryptedLinkPublic 不需要鉴权的加密链接接口
// @Tags EncryptedLink
// @Summary 不需要鉴权的加密链接接口
// @accept application/json
// @Produce application/json
// @Param data query request.EncryptedLinkSearch true "分页获取加密链接列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EL/getEncryptedLinkPublic [get]
func (a *EL) GetEncryptedLinkPublic(c *gin.Context) {
	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceEncryptedLink.GetEncryptedLinkPublic()
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的加密链接接口信息"}, "获取成功", c)
}

// BuildEncryptedUrl 生成加密链接
// @Tags EncryptedLink
// @Summary 生成加密链接
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EL/buildEncryptedUrl [POST]
func (a *EL) BuildEncryptedUrl(c *gin.Context) {
	// 请添加自己的业务逻辑
	var ids request.BuildEncryptedUrlIds
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceEncryptedLink.BuildEncryptedUrl(ids)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}

// HandleRedirect 处理短代码对应重定向
// @Tags EncryptedLink
// @Summary 处理短代码对应重定向
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EL/handleRedirect [GET]
func (a *EL) HandleRedirect(c *gin.Context) {
	shortCode := c.Query("shortCode")
	var link model.EncryptedLink
	if err := global.GVA_DB.Where("short_code = ?", shortCode).First(&link).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 1, "message": "Link not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"redirectUrl": link.RedirectUrl}})
}
