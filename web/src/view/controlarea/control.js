import service from '@/utils/request'

// controlSpiders
// @Tags Spiders
// @Summary 修改是否启用爬虫
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param  enable_spiders bool
// @Success 200 {string} string "{"success":true,"data":"","msg":"修改成功"}"
// @Router /spiders/controlspiders [post]
export const controlSpiders = (data) => {
    return service({
        url: '/spiders/controlspiders',
        method: 'post',
        data
    })
}

// getSpidersStatus
// @Tags Spiders
// @Summary 查询是否启用了爬虫
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param  无需传入参数
// @Success 200 {string} string "{"success":true,"data": true/false,"msg":"查询成功"}"
// @Router /spiders/getspidersstatus [post]
export const getSpidersStatus = () => {
    return service({
        url: '/spiders/getspidersstatus',
        method: 'post',
    })
}

// controlTelegramPush
// @Tags Spiders
// @Summary 修改是否启用Telegram推送
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param  enable_tgpush bool
// @Success 200 {string} string "{"success":true,"data":"","msg":"修改成功"}"
// @Router /telegrampush/modifytgpushcheckstatus [post]
export const controlTelegramPush = (data) => {
    return service({
        url: '/telegrampush/modifytgpushcheckstatus',
        method: 'post',
        data
    })
}

// getTelegramPushStatus
// @Tags Spiders
// @Summary 获取Telegram推送状态
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param  无需传入参数
// @Success 200 {string} string "{"success":true,"data": true/false,"msg":"查询成功"}"
// @Router /telegrampush/gettgpushcheckstatus [post]
export const getTelegramPushStatus = () => {
    return service({
        url: '/telegrampush/gettgpushcheckstatus',
        method: 'post',
    })
}

// controlFAProductsCheck
// @Tags AllPdSpiders
// @Summary 修改是否启用商家所有商品爬虫
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param  enable_allpdspiders bool
// @Success 200 {string} string "{"success":true,"data":"","msg":"修改成功"}"
// @Router /AllPdSpiders/modifyfaproductscheckstatus [post]
export const controlFAProductsCheck = (data) => {
    return service({
        url: '/AllPdSpiders/modifyfaproductscheckstatus',
        method: 'post',
        data
    })
}

// getFAProductsCheckStatus
// @Tags AllPdSpiders
// @Summary 获取商家所有商品爬虫状态
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param  无需传入参数
// @Success 200 {string} string "{"success":true,"data": true/false,"msg":"查询成功"}"
// @Router /AllPdSpiders/getfaproductscheckstatus [post]
export const getFAProductsCheckStatus = () => {
    return service({
        url: '/AllPdSpiders/getfaproductscheckstatus',
        method: 'post',
    })
}

// GetPublicRegisterStatus 获取是否公开注册
// @Tags EcsUsers
// @Summary 获取是否公开注册
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/getPublicRegisterStatus [GET]
export const getPublicRegisterStatus = () => {
    return service({
        url: '/eusr/getPublicRegisterStatus',
        method: 'GET'
    })
}

// ControlPublicRegister 启用关闭公开注册
// @Tags EcsUsers
// @Summary 启用关闭公开注册
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/controlPublicRegister [POST]
export const controlPublicRegister = (data) => {
    return service({
        url: '/eusr/controlPublicRegister',
        method: 'POST',
        data
    })
}

// GetPublicPushStatus 获取整体推送是否启用
// @Tags PusherConfig
// @Summary 获取整体推送是否启用
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /pc/getPublicPushStatus [GET]
export const getPublicPushStatus = () => {
    return service({
        url: '/pc/getPublicPushStatus',
        method: 'GET'
    })
}

// ControlPublicPushStatus 修改整体推送是否启用
// @Tags PusherConfig
// @Summary 修改整体推送是否启用
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /pc/controlPublicPushStatus [POST]
export const controlPublicPushStatus = (data) => {
    return service({
        url: '/pc/controlPublicPushStatus',
        method: 'POST',
        data
    })
}

// GetTelegramBotPushStatus 获取是否启用TG的BOT进行推送
// @Tags PusherConfig
// @Summary 获取是否启用TG的BOT进行推送
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /pc/getTelegramBotPushStatus [GET]
export const getTelegramBotPushStatus = () => {
    return service({
        url: '/pc/getTelegramBotPushStatus',
        method: 'GET'
    })
}

// ControlTelegramBotPushStatus 修改是否启用TG的BOT进行推送
// @Tags PusherConfig
// @Summary 修改是否启用TG的BOT进行推送
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /pc/controlTelegramBotPushStatus [POST]
export const controlTelegramBotPushStatus = (data) => {
    return service({
        url: '/pc/controlTelegramBotPushStatus',
        method: 'POST',
        data
    })
}