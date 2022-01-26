package Weblib

import (
	"github.com/charlie/gomicroconsul/Services"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(prodService Services.ProdService) gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["ProdService"] = prodService
	}
}
