package products

import api "server/api/v1"

type RouterGroup struct{ ProductsRouter }

var pdApi = api.ApiGroupApp.ProductsApiGroup.ProductsApi
