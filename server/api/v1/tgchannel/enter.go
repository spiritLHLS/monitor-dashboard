package tgchannel

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ TgchannelApi }

var tgcService = service.ServiceGroupApp.TgchannelServiceGroup.TgchannelService
