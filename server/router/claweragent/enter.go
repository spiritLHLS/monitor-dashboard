package claweragent

import api "server/api/v1"

type RouterGroup struct{ AgentIpRecordRouter }

var airApi = api.ApiGroupApp.ClaweragentApiGroup.AgentIpRecordApi
