package products

import "server/service"

type ApiGroup struct{ ProductsApi }

var pdService = service.ServiceGroupApp.ProductsServiceGroup.ProductsService
