package tgchannel

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ TgchannelRouter }

var tgcApi = api.ApiGroupApp.TgchannelApiGroup.TgchannelApi
