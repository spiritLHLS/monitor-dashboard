package digitalproducts

import "server/service"

type ApiGroup struct{ DigitalProductsApi }

var dpdService = service.ServiceGroupApp.DigitalproductsServiceGroup.DigitalProductsService
