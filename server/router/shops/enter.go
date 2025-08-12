package shops

import api "server/api/v1"

type RouterGroup struct{ ShopsRouter }

var spApi = api.ApiGroupApp.ShopsApiGroup.ShopsApi
