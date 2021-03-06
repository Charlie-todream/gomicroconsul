package main

import (
	ProdServcie "github.com/charlie/gomicroconsul/ProdService"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
	"net/http"
)

func main() {
	// 新建_一个consul注册地址
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)


	ginRouter := gin.Default()

	ginRouter.Handle("GET", "/user", func(context *gin.Context) {
		context.String(200, "user api")
	})
	ginRouter.Handle("GET", "/news", func(context *gin.Context) {
		context.String(200, "news api")
	})

	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", func(context *gin.Context) {
			var pr ProdServcie.ProdsRequest
			err := context.Bind(&pr)
			if err != nil || pr.Size <=0 {
				log.Println(err)
				pr = ProdServcie.ProdsRequest{Size: 2}
			}
			context.JSON(
				http.StatusOK,
				gin.H{
					"data":ProdServcie.NewProdList(pr.Size),
				})
		})

	}

	server := web.NewService(
		web.Name("prodservice"), // 注册到consul服务中的service name
		web.Address(":8001"),
		web.Handler(ginRouter),
		web.Metadata(map[string]string{"protocol": "http"}), // 否则500错误
	web.Registry(consulReg), // 注册到哪个服务器上的consul中
	)
	server.Init()
	server.Run()
}
