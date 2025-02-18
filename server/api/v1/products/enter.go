package products

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ProductsApi }

var pdService = service.ServiceGroupApp.ProductsServiceGroup.ProductsService
