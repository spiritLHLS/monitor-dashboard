package digitalproducts

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ DigitalProductsRouter }

var dpdApi = api.ApiGroupApp.DigitalproductsApiGroup.DigitalProductsApi
