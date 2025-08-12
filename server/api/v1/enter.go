package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/digitalproducts"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ecsusers"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/findallpd"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/invite_codes"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/partitions"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/privmsg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/products"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/shops"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/subscribe"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/tgchannel"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup          system.ApiGroup
	ExampleApiGroup         example.ApiGroup
	ProductsApiGroup        products.ApiGroup
	ShopsApiGroup           shops.ApiGroup
	PartitionsApiGroup      partitions.ApiGroup
	TgchannelApiGroup       tgchannel.ApiGroup
	FindallpdApiGroup       findallpd.ApiGroup
	EcsusersApiGroup        ecsusers.ApiGroup
	SubscribeApiGroup       subscribe.ApiGroup
	PrivmsgApiGroup         privmsg.ApiGroup
	Invite_codesApiGroup    invite_codes.ApiGroup
	DigitalproductsApiGroup digitalproducts.ApiGroup
}
