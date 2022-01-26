package main

import (
	"context"
	"fmt"
	"github.com/charlie/gomicroconsul/Services"
	"github.com/charlie/gomicroconsul/client/Weblib"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
)
// 使用go-micro官方提供的wrapper来对路由进行装饰
type logWrapper struct {
	client.Client
}

// 重写call方法
func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func newLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}
// 在gin 中调用上节课构建的rpc服务
func main()  {

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
		)

	myService := micro.NewService(
		micro.Name("prodservice.client"),
		micro.WrapClient(newLogWrapper),  // 使用装饰器

		)
	prodService := Services.NewProdService("ProdService",myService.Client())
	service := web.NewService(
		web.Name("ProdService.client"),
		web.Address(":9000"),
		web.Handler(Weblib.InitRouter(prodService)),
		web.Registry(consulReg),
		)
	service.Init()
	service.Run()
}

