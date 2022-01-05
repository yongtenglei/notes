package internal

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

var AccountServiceWebClientConsulClient *api.Client

func init() {
	// config
	defaultConfig := api.DefaultConfig()
	h := ConsulHost
	p := ConsulPort
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)

	var err error
	// client
	AccountServiceWebClientConsulClient, err = api.NewClient(defaultConfig)
	if err != nil {
		log.Fatal(err)
	}

}

const (
	ConsulHost = "172.19.0.4"
	ConsulPort = 8500
)

func ConsulRegister(name, id, host string, port int, tags []string) error {

	// 注册结构体
	//agentServiceRegistration := &api.AgentServiceRegistration{
	//ID:      id,
	//Name:    name,
	//Tags:    tags,
	//Address: host,
	//Port:    port,
	//}
	agentServiceRegistration := new(api.AgentServiceRegistration)
	agentServiceRegistration.ID = id
	agentServiceRegistration.Name = name
	agentServiceRegistration.Port = port
	agentServiceRegistration.Tags = tags
	agentServiceRegistration.Address = host

	// 健康检查结构体
	//serverAddr := fmt.Sprintf("http://%s:%d/health", host, port)
	//serverAddr := "http://localhost:9410/health"
	//check := &api.AgentServiceCheck{
	//HTTP:                           serverAddr,
	//Timeout:                        "3s",
	//Interval:                       "1s",
	//DeregisterCriticalServiceAfter: "5s",
	//}
	check := new(api.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", agentServiceRegistration.Address, agentServiceRegistration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s"

	agentServiceRegistration.Check = check

	err := AccountServiceWebClientConsulClient.Agent().ServiceRegister(agentServiceRegistration)
	fmt.Println("---------------------------4--------------------")

	return err
}

func ConsulServicesList() error {
	// config
	defaultConfig := api.DefaultConfig()
	h := ConsulHost
	p := ConsulPort
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)

	// client
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		return err
	}

	servicesList, err := client.Agent().Services()
	if err != nil {
		return err
	}

	for k, v := range servicesList {
		fmt.Printf("%s, %v\n", k, v)
		fmt.Println("===================")
	}
	return nil
}

func ConsulServicesListWithFilter() error {
	// config
	defaultConfig := api.DefaultConfig()
	h := ConsulHost
	p := ConsulPort
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)

	// client
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		return err
	}

	servicesList, err := client.Agent().ServicesWithFilter("Service==svc1")
	if err != nil {
		return err
	}

	for k, v := range servicesList {
		fmt.Printf("%s, %v\n", k, v)
		fmt.Println("===================")
	}
	return nil
}

func ConsulDeRegister(id string) error {
	err := AccountServiceWebClientConsulClient.Agent().ServiceDeregister(id)
	return err
}
