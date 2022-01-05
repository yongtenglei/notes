package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ServerConf struct {
	APPName string `mapstructure:"appname"`
	Version string `mapstructure:"version"`
}

var serverConf ServerConf

func Getenv(s string) int {
	// 寻找系统中可用的环境变量
	viper.AutomaticEnv()
	return viper.GetInt(s)

}

func main() {
	mode := Getenv("VIPER_MODE")
	// 假设 0 为 线下配置
	if mode == 0 {
		viper.SetConfigFile("./conf.yaml")
	}

	// option 1
	viper.SetConfigFile("./product.yaml")

	// option 2
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")

	// 读取配置
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// 解析成结构体
	if err := viper.Unmarshal(&serverConf); err != nil {
		fmt.Println("Read Config failed\n", err)
	}

	// 热加载配置
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config was Changed: ", e.Name)
	})

	if err := viper.Unmarshal(&serverConf); err != nil {
		fmt.Println("Read Config failed\n", err)
	}

	fmt.Println(serverConf.APPName)
	fmt.Println(serverConf.Version)

}
