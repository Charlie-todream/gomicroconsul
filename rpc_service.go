package main

import (
	"github.com/charlie/gomicroconsul/ServiceImpl"
	"github.com/charlie/gomicroconsul/Services"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	consuleReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	service := micro.NewService(
		micro.Name("prodservice"),
		micro.Address(":8011"),
		micro.Registry(consuleReg),
	)

	// 调用pb文件中的handler的方法注册service的服务和我们实现ProdServiceHandler接口的结构体指针即可
	Services.RegisterProdServiceHandler(service.Server(), new(ServiceImpl.ProdService))
	service.Init()
	service.Run()

}
