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