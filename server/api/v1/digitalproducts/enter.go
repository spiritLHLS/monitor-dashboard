package digitalproducts

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ DigitalProductsApi }

var dpdService = service.ServiceGroupApp.DigitalproductsServiceGroup.DigitalProductsService
