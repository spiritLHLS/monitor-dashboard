package router

import (
	"server/router/claweragent"
	"server/router/digitalproducts"
	"server/router/ecsusers"
	"server/router/example"
	"server/router/findallpd"
	"server/router/invite_codes"
	"server/router/partitions"
	"server/router/privmsg"
	"server/router/products"
	"server/router/shops"
	"server/router/subscribe"
	"server/router/system"
	"server/router/tgchannel"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System          system.RouterGroup
	Example         example.RouterGroup
	Products        products.RouterGroup
	Shops           shops.RouterGroup
	Partitions      partitions.RouterGroup
	Tgchannel       tgchannel.RouterGroup
	Findallpd       findallpd.RouterGroup
	Ecsusers        ecsusers.RouterGroup
	Subscribe       subscribe.RouterGroup
	Privmsg         privmsg.RouterGroup
	Invite_codes    invite_codes.RouterGroup
	Digitalproducts digitalproducts.RouterGroup
	Claweragent     claweragent.RouterGroup
}
