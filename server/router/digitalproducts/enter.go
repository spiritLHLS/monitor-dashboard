package digitalproducts

import api "server/api/v1"

type RouterGroup struct{ DigitalProductsRouter }

var dpdApi = api.ApiGroupApp.DigitalproductsApiGroup.DigitalProductsApi
