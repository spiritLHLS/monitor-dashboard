package service

import (
	"server/service/digitalproducts"
	"server/service/ecsusers"
	"server/service/example"
	"server/service/findallpd"
	"server/service/invite_codes"
	"server/service/partitions"
	"server/service/privmsg"
	"server/service/products"
	"server/service/shops"
	"server/service/subscribe"
	"server/service/system"
	"server/service/tgchannel"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup          system.ServiceGroup
	ExampleServiceGroup         example.ServiceGroup
	ProductsServiceGroup        products.ServiceGroup
	ShopsServiceGroup           shops.ServiceGroup
	PartitionsServiceGroup      partitions.ServiceGroup
	TgchannelServiceGroup       tgchannel.ServiceGroup
	FindallpdServiceGroup       findallpd.ServiceGroup
	EcsusersServiceGroup        ecsusers.ServiceGroup
	SubscribeServiceGroup       subscribe.ServiceGroup
	PrivmsgServiceGroup         privmsg.ServiceGroup
	Invite_codesServiceGroup    invite_codes.ServiceGroup
	DigitalproductsServiceGroup digitalproducts.ServiceGroup
}
