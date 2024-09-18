package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ecsusers"
	"github.com/flipped-aurora/gin-vue-admin/server/model/findallpd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/partitions"
	"github.com/flipped-aurora/gin-vue-admin/server/model/products"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shops"
	"github.com/flipped-aurora/gin-vue-admin/server/model/subscribe"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tgchannel"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(products.Products{}, shops.Shops{}, partitions.Partitionspage{}, tgchannel.Tgchannel{}, findallpd.Findallpd{}, ecsusers.EcsUsers{}, subscribe.Subscribe{})
	if err != nil {
		return err
	}
	return nil
}
