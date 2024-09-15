package findallpd

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ FindallpdRouter }

var fapdApi = api.ApiGroupApp.FindallpdApiGroup.FindallpdApi
