package main

import (
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"io/ioutil"
	"log"
	"net/http"
)

// etcd 通过轮询获取服务
// 前提启动之前多个后端服务
// 基本方式调用后端服务

func callAPI(addr, path, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func main() {
	// consul 连接句柄
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"))

	// 获取服务
	getService, err := consulReg.GetService("prodservice")
	if err != nil {
		log.Fatalf("get service failed, err:%v\n", err)
		return
	}

	next := selector.RoundRobin(getService)

	node, err := next()
	if err != nil {
		log.Fatalln(err)
		return
	}

	res, err := callAPI(node.Address, "/v1/prods", "POST")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(res)

}