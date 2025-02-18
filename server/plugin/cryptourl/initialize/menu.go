package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{{ParentId: 0, Path: "cryptourlMenu", Name: "cryptourlMenu", Hidden: false, Component: "view/routerHolder.vue", Sort: 0, Meta: model.Meta{Title: "加密链接", Icon: "school"}}, {ParentId: 0, Path: "EL", Name: "EL", Hidden: false, Component: "plugin/cryptourl/view/encryptedlink.vue", Sort: 10, Meta: model.Meta{Title: "加密链接", Icon: "aim"}}}
	utils.RegisterMenus(entities...)
}
