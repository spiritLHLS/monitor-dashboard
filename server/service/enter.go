package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/ecsusers"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/findallpd"
	"github.com/flipped-aurora/gin-vue-admin/server/service/invite_codes"
	"github.com/flipped-aurora/gin-vue-admin/server/service/partitions"
	"github.com/flipped-aurora/gin-vue-admin/server/service/privmsg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/products"
	"github.com/flipped-aurora/gin-vue-admin/server/service/shops"
	"github.com/flipped-aurora/gin-vue-admin/server/service/subscribe"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/tgchannel"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	ExampleServiceGroup      example.ServiceGroup
	ProductsServiceGroup     products.ServiceGroup
	ShopsServiceGroup        shops.ServiceGroup
	PartitionsServiceGroup   partitions.ServiceGroup
	TgchannelServiceGroup    tgchannel.ServiceGroup
	FindallpdServiceGroup    findallpd.ServiceGroup
	EcsusersServiceGroup     ecsusers.ServiceGroup
	SubscribeServiceGroup    subscribe.ServiceGroup
	PrivmsgServiceGroup      privmsg.ServiceGroup
	Invite_codesServiceGroup invite_codes.ServiceGroup
}
