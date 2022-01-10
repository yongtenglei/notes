package internal

import (
	"fmt"
	"testing"

	"github.com/rey/micro-demo/setting"
)

func TestConsulRegister(t *testing.T) {

	err := ConsulRegister(
		setting.AccountServiceConf.AccountWebClientConfig.Name,
		setting.AccountServiceConf.AccountWebClientConfig.ID,
		setting.AccountServiceConf.AccountWebClientConfig.Host,
		int(setting.AccountServiceConf.AccountWebClientConfig.Port),
		setting.AccountServiceConf.AccountWebClientConfig.Tags,
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ok")

}
