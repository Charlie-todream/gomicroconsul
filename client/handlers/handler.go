package handlers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/charlie/gomicroconsul/Services"
	"github.com/gin-gonic/gin"
)

// GetProdList 显示商品列表
func GetProdList(c *gin.Context) {
	prodService := c.Keys["ProdService"].(Services.ProdService)
	var prodReq Services.ProdRequest
	err := c.Bind(&prodReq)
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error()})
	} else {
		// 熔断器代码改造
		// 1.配置config
		configA := hystrix.CommandConfig{
			Timeout: 5000,
		}
		// 2.配置command
		hystrix.ConfigureCommand("getProds",configA)
		// 3. 执行Do 方法
		var prodRes * Services.ProdListResponse
		err := hystrix.Do("getProds", func() error {
			prodRes,err = prodService.GetProdList(context.Background(),&prodReq)
			return err
		},nil)
		if err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
		} else {
			c.JSON(200, gin.H{"data": prodRes.Data})
		}
	}

}