package controllers

import (
	"net/http"
	"strconv"

	"github.com/cuongtop4598/WebBlockChain/models"
	"github.com/gin-gonic/gin"
)

// ProductController retrieves a list of available products
func ProductController(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, models.Products)
}

// buyProduct increments the amount of product item for buying
func BuyProduct(c *gin.Context) {
	// confir, product ID sent is valid
	// remember to import the `strconv` package
	if productid, err := strconv.Atoi(c.Param("productID")); err == nil {
		// find product, and increment cart
		for i := 0; i < len(models.Products); i++ {
			if models.Products[i].ID == productid {
				models.Products[i].Amount++
			}
		}
		//return a pointer to the updated products list
		c.JSON(http.StatusOK, &models.Products)
	} else {
		// product ID is invalid
		c.AbortWithStatus(http.StatusNotFound)
	}
}
