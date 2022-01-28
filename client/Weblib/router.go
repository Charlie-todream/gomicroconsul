package Weblib

import (
	"github.com/charlie/gomicroconsul/Services"
	"github.com/charlie/gomicroconsul/client/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouter(prodService Services.ProdService)  *gin.Engine{
	ginRouter := gin.Default()
	ginRouter.Use(InitMiddleware(prodService)) // 使用中间件去封装prodService 到ctx中
	ginRouter.Use(ErrorMiddleware())

	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST","/prods",handlers.GetProdList)
		v1Group.Handle("GET","/prods/:pid",handlers.GetProdDetail)
	}
	return ginRouter
}