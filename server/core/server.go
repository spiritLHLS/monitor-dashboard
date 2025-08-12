package core

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	"server/initialize"
	"server/service/system"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		initialize.RedisList()
	}
	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	time.Sleep(10 * time.Second) // 需要等待10秒确保所有前面的组件完全起来了，否则会出现接口异常导致panic
	s := initServer(address, Router)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
