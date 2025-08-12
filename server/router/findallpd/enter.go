package findallpd

import api "server/api/v1"

type RouterGroup struct{ FindallpdRouter }

var fapdApi = api.ApiGroupApp.FindallpdApiGroup.FindallpdApi
