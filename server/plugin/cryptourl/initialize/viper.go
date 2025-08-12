package initialize

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"server/global"
	"server/plugin/cryptourl/plugin"
)

func Viper() {
	err := global.GVA_VP.UnmarshalKey("cryptourl", &plugin.Config)
	if err != nil {
		err = errors.Wrap(err, "初始化配置文件失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
