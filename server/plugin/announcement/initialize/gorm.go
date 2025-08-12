package initialize

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"server/global"
	"server/plugin/announcement/model"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate(
		new(model.Info),
	)
	if err != nil {
		err = errors.Wrap(err, "注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
