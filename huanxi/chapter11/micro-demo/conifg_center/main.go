package main

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	// 服务端设置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "192.168.1.8",
			Port:   8848,
		},
	}

	// 客户端设置
	clientConfig := constant.ClientConfig{
		NamespaceId:         "ae91b401-1490-4b03-a2bc-5e7e088c4e95",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		RotateTime:          "24h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// config client 设置
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}

	// 获得配置内容
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "config",
		Group:  "development",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(content)

}
