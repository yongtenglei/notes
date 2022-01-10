package setting

import (
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

var NacosConf *NacosConfig
var AccountServiceConf *AccountServerConf

func initNacos() {
	viper.SetConfigName("nacos_config") // name of config file (without extension)
	viper.SetConfigType("yaml")         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("..")           // optionally look for config in the working directory
	viper.AddConfigPath(".")            // optionally look for config in the working directory

	//viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("%s:%s", "load config failed", err))
	}

	if err := viper.Unmarshal(&NacosConf); err != nil {
		panic(err)
	}

	fmt.Println(NacosConf)

}

func initFromNacos() {
	// 服务端设置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: NacosConf.Host,
			Port:   uint64(NacosConf.Port),
		},
	}

	// 客户端设置
	clientConfig := constant.ClientConfig{
		NamespaceId:         NacosConf.Namespace,
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
		DataId: NacosConf.DataId,
		Group:  NacosConf.Group,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(content)

	json.Unmarshal([]byte(content), &AccountServiceConf)
	fmt.Println("===============")
	fmt.Println(AccountServiceConf)
	fmt.Println("===============")

}

func init() {
	initNacos()

	initFromNacos()

	fmt.Printf("%v\n", AccountServiceConf)
	fmt.Printf("%v\n", AccountServiceConf.AccountWebServerConfig)
	fmt.Printf("%v\n", AccountServiceConf.AccountWebClientConfig)
}
