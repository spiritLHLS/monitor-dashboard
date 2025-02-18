package products

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ProductsRouter }

var pdApi = api.ApiGroupApp.ProductsApiGroup.ProductsApi
