package initialize

import (
	"server/global"
	"server/model/digitalproducts"
	"server/model/ecsusers"
	"server/model/findallpd"
	"server/model/invite_codes"
	"server/model/partitions"
	"server/model/privmsg"
	"server/model/products"
	"server/model/shops"
	"server/model/subscribe"
	"server/model/tgchannel"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(products.Products{}, shops.Shops{}, partitions.Partitionspage{}, tgchannel.Tgchannel{}, findallpd.Findallpd{}, ecsusers.EcsUsers{}, subscribe.Subscribe{}, privmsg.PusherConfig{}, invite_codes.InviteCodes{}, digitalproducts.DigitalProducts{})
	if err != nil {
		return err
	}
	return nil
}
