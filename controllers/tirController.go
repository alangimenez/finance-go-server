package controllers

import (
	"net/http"

	"github.com/alangimenez/finance-go-server/services"
	"github.com/gin-gonic/gin"
)

func GetTirs(c *gin.Context) {
	listOfTirResponse := services.GetTirs()

	c.JSON(http.StatusOK, listOfTirResponse)
}

func CalculateTirWithGivenPrice(c *gin.Context) {
	price := GetPrice(c)
	if price == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "price must be a float64 type",
		})
		return
	}
	ticket := GetTicket(c)
	if ticket == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ticket must not be empty",
		})
		return
	}

	tir, err := services.CalculateTirWithGivenPrice(price, ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tir)
}

func RetrieveQuotesByAssetType(c *gin.Context) {
	assetType := c.Param("assetType")

	quotes := services.GetQuotesByAssetType(assetType)

	c.JSON(http.StatusOK, quotes)
}
