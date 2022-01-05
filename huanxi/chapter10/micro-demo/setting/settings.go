package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

var AccountServiceConf *AccountServerConf

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AddConfigPath("..")     // optionally look for config in the working directory

	//viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("%s:%s", "load config failed", err))
	}

	if err := viper.Unmarshal(&AccountServiceConf); err != nil {
		panic(err)
	}

	fmt.Printf("%v", AccountServiceConf)
	fmt.Printf("%v", AccountServiceConf.AccountWebServerConfig)
	fmt.Printf("%v", AccountServiceConf.AccountWebClientConfig)
	//fmt.Println(AccountServiceConf.AccountWebServerConfig.Name)
	//fmt.Println(AccountServiceConf.AccountWebServerConfig.Host)
	//fmt.Println(AccountServiceConf.AccountWebServerConfig.Port)
	//fmt.Println(AccountServiceConf.AccountWebServerConfig.Tags)
}
