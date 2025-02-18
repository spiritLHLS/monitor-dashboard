package shops

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ShopsRouter }

var spApi = api.ApiGroupApp.ShopsApiGroup.ShopsApi
