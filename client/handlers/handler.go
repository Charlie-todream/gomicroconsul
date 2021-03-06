package handlers

import (
	"context"
	"github.com/charlie/gomicroconsul/Services"
	"github.com/gin-gonic/gin"
	"strconv"
)

func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdID: id, ProdName: pname}
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
func GetProdDetail(ginCtx *gin.Context) {

	var prodReq Services.ProdRequest
	PanicIfError(ginCtx.BindUri(&prodReq)) //这里要绑定的是uri，因为是get请求参数从uri中拿到的不是form表单
	prodService := ginCtx.Keys["ProdService"].(Services.ProdService) //类型断言为对应的请求类型
	resp, _ := prodService.GetProdDetail(context.Background(), &prodReq)
	ginCtx.JSON(200, gin.H{"data": resp.Data})

}


// 降级后的默认商品
func defaultProds()(*Services.ProdListResponse,error) {
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		models = append(models, newProd(20+i, "prodName"+strconv.Itoa(100+int(i))))
	}
	res := &Services.ProdListResponse{}
	res.Data = models
	return res,nil
}


// GetProdList 显示商品列表
func GetProdList(c *gin.Context) {
	prodService := c.Keys["ProdService"].(Services.ProdService)
	var prodReq Services.ProdRequest
	err := c.Bind(&prodReq)
	if err != nil {
		c.JSON(500, gin.H{"status": err.Error()})
	} else {
		prodRes, _ := prodService.GetProdList(context.Background(), &prodReq)
		c.JSON(200, gin.H{"data": prodRes.Data})
	}



	//if err != nil {
	//	c.JSON(500, gin.H{
	//		"status": err.Error()})
	//} else {
	//	// 熔断器代码改造
	//	// 1.配置config
	//	configA := hystrix.CommandConfig{
	//		Timeout: 5000,
	//	}
	//	// 2.配置command
	//	hystrix.ConfigureCommand("getProds",configA)
	//	// 3. 执行Do 方法
	//	var prodRes * Services.ProdListResponse
	//	err := hystrix.Do("getProds", func() error {
	//		prodRes,err = prodService.GetProdList(context.Background(),&prodReq)
	//		return err
	//	}, func(err error) error {
	//		prodRes,err = defaultProds()
	//		return err
	//	})
	//
	//}

}