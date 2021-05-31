package consul

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"portal/global/config"
)

var client *consulapi.Client

func InitConsul(){
	conf := config.GetAppConfig().Consul
	config := consulapi.DefaultConfig()
	address := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	config.Address = address
	var err error
	client, err = consulapi.NewClient(config) //创建客户端
	if err != nil {
		panic(err)
	}
}

func GetClient() *consulapi.Client {
	return client
}
