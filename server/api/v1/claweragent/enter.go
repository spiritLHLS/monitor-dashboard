package claweragent

import "server/service"

type ApiGroup struct{ AgentIpRecordApi }

var airService = service.ServiceGroupApp.ClaweragentServiceGroup.AgentIpRecordService
