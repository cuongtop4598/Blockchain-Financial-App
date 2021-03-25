package main

import (
	"net/http"

	"github.com/cuongtop4598/WebBlockChain/controllers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve frontend satic files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.Use(static.Serve("/src", static.LocalFile("./src", true)))
	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"title":   "Home",
				"message": "pong",
			})
		})

		// Our API will consit of two routes
		// /products - which will retrieve a list of products a user can see
		// /products/buy/:productID - which will capture amount of products and sent to a particular product
		api.GET("/products", controllers.ProductController)
		api.POST("/products/buy/:productID", controllers.BuyProduct)

		// Start and run the server
		router.Run(":3004")
	}
}
