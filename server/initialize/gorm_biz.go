package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/products"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shops"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(products.Products{}, shops.Shops{})
	if err != nil {
		return err
	}
	return nil
}
