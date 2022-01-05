package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rey/micro-demo/log"
	"github.com/rey/micro-demo/web/middlewares"
	"github.com/rey/micro-demo/web/routers"
	"go.uber.org/zap"
)

var (
	ip   *string
	port *int
)

func init() {
	ip = flag.String("ip", "127.0.0.1", "specific ip")
	port = flag.Int("port", 9410, "specific port")
	flag.Parse()
}

func main() {

	addr := fmt.Sprintf("%s:%d", *ip, *port)

	r := gin.Default()

	r.POST("/login", routers.LoginByPasswordHandler)

	accountGroup := r.Group("/v1/account")
	accountGroup.Use(middlewares.JWTAuth())
	{
		accountGroup.GET("/list", routers.AccountListHandler)

		accountGroup.GET("/captcha", routers.CaptchaHandler)
	}

	web_srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := web_srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	<-sig
	log.Logger.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := web_srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Server Shutdown", zap.Error(err))
	}
	log.Logger.Info("Server exiting")

}
