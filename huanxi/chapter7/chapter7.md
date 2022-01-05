[TOC]

# Chapter 7

## ZAP

### quick start

#### suger logger

```go
package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()

	// 第一个参数为msg, 之后为附加
	sugar.Infow("from infow",
		// Structured context as loosely typed key-value pairs.
		"stutus", "ok",
	)

	// 直接写入msg
	sugar.Infof("status %s", "ok")
}
// output:
// {"level":"info","ts":1639041275.566963,"caller":"suger/main.go:14","msg":"from infow","stutus":"ok"}
// {"level":"info","ts":1639041275.5670037,"caller":"suger/main.go:20","msg":"status ok"}


```

#### logger

```go
package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// 第一个参数为msg, 之后为附加
	logger.Info("from info",
		zap.String("status", "ok"),
		zap.Int("code", 200),
	)

}
// output:
// {"level":"info","ts":1639041198.5196414,"caller":"logger/main.go:12","msg":"from info","status":"ok","code":200}

```

### To file

```go
package main

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = append(config.OutputPaths, "./mylog.log")
	return config.Build()

}

func main() {
	logger, _ = NewLogger()
	defer logger.Sync() // flushes buffer, if any

	// 第一个参数为msg, 之后为附加
	logger.Info("from info",
		zap.String("status", "ok"),
		zap.Int("code", 200),
	)

}

// output:
// {"level":"info","ts":1639041666.5249786,"caller":"to_file/main.go:21","msg":"from info","status":"ok","code":200}


```

```log

{"level":"info","ts":1639041666.5249786,"caller":"to_file/main.go:21","msg":"from info","status":"ok","code":200}

```

### 全局 logger zap.L() 与 全局 suger logger zap.S()

```go
	logger, _ = NewLogger()
	defer logger.Sync() // flushes buffer, if any

	zap.L().Info("global logger")
	zap.S().Info("global suger logger")

```

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNly1gx7pom89dgj30qg0akq68.jpg">

</div>

直接调用全局 logger 并无输出, 它只是一个空壳.

```go
func main() {
	logger, _ = NewLogger()
	defer logger.Sync() // flushes buffer, if any

	zap.ReplaceGlobals(logger)
	zap.L().Info("global logger")
	zap.S().Info("global suger logger")

}

```

<div align=center><img src="https://tva3.sinaimg.cn/large/006cK6rNly1gx7psqkbycj30pk08u42j.jpg">

</div>

使用 zap.ReplaceGlobals(logger), 进行填充.

## viper

### Quick Start 读取配置到结构体, 热加载

```go
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

func main() {
	// option 1
	viper.SetConfigFile("./conf.yaml")

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

```

```yaml
// conf.yaml
appname: "app"
version: "0.0.1"
```

<div align=center><img src="https://tva3.sinaimg.cn/large/006cK6rNly1gx7qc2td65j30jx09bmzw.jpg">

</div>

### 切换生产环境配置

```go
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
```

<div align=center><img src="https://tva3.sinaimg.cn/large/006cK6rNly1gx7qpyx890j30hm093aco.jpg">

</div>
