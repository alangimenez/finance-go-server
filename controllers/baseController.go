package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPrice(c *gin.Context) float64 {
	stringPrice := c.Query("price")
	price, err := strconv.ParseFloat(stringPrice, 64)
	if err != nil {
		return 0.0
	}
	return price
}

func GetTicket(c *gin.Context) string {
	return c.Query("ticket")
}
