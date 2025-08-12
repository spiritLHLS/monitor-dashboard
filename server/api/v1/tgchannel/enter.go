package tgchannel

import "server/service"

type ApiGroup struct{ TgchannelApi }

var tgcService = service.ServiceGroupApp.TgchannelServiceGroup.TgchannelService
