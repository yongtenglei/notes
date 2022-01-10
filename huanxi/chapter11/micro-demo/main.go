package main

import (
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/rey/micro-demo/biz"
	_ "github.com/rey/micro-demo/dao/mysql"
	_ "github.com/rey/micro-demo/dao/redis"
	"github.com/rey/micro-demo/internal"
	"github.com/rey/micro-demo/pkg/util"
	"github.com/rey/micro-demo/proto/account"
	"github.com/rey/micro-demo/setting"
	_ "github.com/rey/micro-demo/setting"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// Use Random Port
	//ip := flag.String("ip", "0.0.0.0", "specific ip")
	//port := flag.Int("port", 9409, "specific port")
	//flag.Parse()

	randomPort := util.RandomPort("localhost")

	addr := fmt.Sprintf("%s:%d", setting.AccountServiceConf.AccountWebServerConfig.Host, randomPort)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	// 注册服务
	account.RegisterAccountServiceServer(server, &biz.AccountServer{})

	// 注册健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	defaultConfig := api.DefaultConfig()
	defaultConfig.Address = fmt.Sprintf("%s:%d", internal.ConsulHost, internal.ConsulPort)

	accountServiceWebServerConsulClient, err := api.NewClient(defaultConfig)
	if err != nil {
		panic(err)
	}

	randomId := uuid.New().String()
	agentServiceRegistration := &api.AgentServiceRegistration{
		ID:      randomId,
		Name:    setting.AccountServiceConf.AccountWebServerConfig.Name,
		Tags:    setting.AccountServiceConf.AccountWebServerConfig.Tags,
		Address: setting.AccountServiceConf.AccountWebServerConfig.Host,
		Port:    randomPort,
	}

	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", agentServiceRegistration.Address, agentServiceRegistration.Port),
		Timeout:                        "3s",
		Interval:                       "3s",
		DeregisterCriticalServiceAfter: "5s",
	}

	agentServiceRegistration.Check = check

	err = accountServiceWebServerConsulClient.Agent().ServiceRegister(agentServiceRegistration)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Service %s start on %s:%d", randomId, setting.AccountServiceConf.AccountWebServerConfig.Host, randomPort)

	if err = server.Serve(l); err != nil {
		panic(err)
	}

}
