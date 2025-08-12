package tgchannel

import api "server/api/v1"

type RouterGroup struct{ TgchannelRouter }

var tgcApi = api.ApiGroupApp.TgchannelApiGroup.TgchannelApi
