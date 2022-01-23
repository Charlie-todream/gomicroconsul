package main

import (
	"context"
	"github.com/charlie/gomicroconsul/Models"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	myhttp "github.com/micro/go-plugins/client/http/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
)

func callApiTwo(s selector.Selector)  {
	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)

	//req := myClient.NewRequest("prodservice", "/v1/prods", map[string]interface{}{"size":4})
	req := myClient.NewRequest("prodservice", "/v1/prods",Models.ProdRequest{Size: 6})
	//var rsp map[string]interface{}
	var rsp Models.ProdListResponse // 这里使用生成的response对象，客户端只需要传入这个就可以了，无需关心服务端返回什么格式，因为服务端已经用rpc框架定义好了这一切，我们使用即可
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(rsp.GetData())
}

func main()  {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
		)
	//name ,_:= consulReg.GetService("prodservice")
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin), // 设置查询策略，这里是轮询
		)

	callApiTwo(mySelector)

}