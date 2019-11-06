package fridayroutes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MountRoute func will mount all rest routes
func MountRoute(router *gin.RouterGroup) {
	router.GET("/health", healthCheck)
	router.GET("/pricing", getPricing)
	router.GET("/rate/up", rateAdUp)
	router.GET("/rate/down", rateAdDown)
	router.GET("/rate/rating", getAdRating)
	router.GET("/similar", getSimilar)

	router.OPTIONS("/*any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Goosfraba"})
	})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Still Alive"})
}

func getPricing(c *gin.Context) {
	region := c.Query("region")
	category := c.Query("cg")
	query := c.Query("q")
	c.JSON(handleGerPricing(region, category, query))
}

func rateAdUp(c *gin.Context) {
	listID := c.Query("list_id")
	c.JSON(handleRateAdUp(listID))
}

func rateAdDown(c *gin.Context) {
	listID := c.Query("list_id")
	c.JSON(handleRateAdDown(listID))
}
func getAdRating(c *gin.Context) {
	listID := c.Query("list_id")
	c.JSON(handlegetAdRating(listID))
}

func getSimilar(c *gin.Context) {
	price := c.Query("price")
	category := c.Query("cg")
	query := c.Query("q")
	c.JSON(handleGetSimilarAds(category, price, query))
}
