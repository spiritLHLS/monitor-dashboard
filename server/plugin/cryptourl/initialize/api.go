package initialize

import (
	"context"
	model "server/model/system"
	"server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{{Path: "/EL/handleRedirect", Description: "跳转加密地址", ApiGroup: "加密链接", Method: "GET"}}
	utils.RegisterApis(entities...)
}
