package v1

import (
	"server/api/v1/claweragent"
	"server/api/v1/digitalproducts"
	"server/api/v1/ecsusers"
	"server/api/v1/example"
	"server/api/v1/findallpd"
	"server/api/v1/invite_codes"
	"server/api/v1/partitions"
	"server/api/v1/privmsg"
	"server/api/v1/products"
	"server/api/v1/shops"
	"server/api/v1/subscribe"
	"server/api/v1/system"
	"server/api/v1/tgchannel"
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
	ClaweragentApiGroup     claweragent.ApiGroup
}
