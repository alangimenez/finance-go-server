package main

import (
	"github.com/alangimenez/finance-go-server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/tir", controllers.GetTirs)
	r.GET("/tir/price", controllers.CalculateTirWithGivenPrice)
	r.GET("/quotes/:assetType", controllers.RetrieveQuotesByAssetType)

	r.Run() // listen and serve on 0.0.0.0:8080
}
