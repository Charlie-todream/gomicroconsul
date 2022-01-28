package Weblib

import (
	"fmt"
	"github.com/charlie/gomicroconsul/Services"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(prodService Services.ProdService) gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["ProdService"] = prodService
	}
}

func ErrorMiddleware() gin.HandlerFunc{
	return func(context *gin.Context) {
		defer func() {
			if r := recover();r !=nil {
				context.JSON(500,gin.H{"status":fmt.Sprintf("%s",r)})
				context.Abort()
			}
		}()
		context.Next()
	}
}


