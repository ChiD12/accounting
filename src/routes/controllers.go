package routes

import (
	"github.com/gin-gonic/gin"
)

// NewsController <controller>
// is used for describing controller actions for news.
type Controller struct{}

// Get <function>
// is used to handle get action of news controller which will return <count> number of news.
// url: /v1/news?count=80 , by default <count> = 50
func (nc Controller) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"method":  "v1/news",
		"message": "Hello from Get function!",
	})
}

// GetSources <function>
// is used to handle get action of news controller which will return all news sources.
// url: /v1/news/sources
func (nc Controller) SignUp(c *gin.Context) {
	c.JSON(200, gin.H{
		"method":  "v1/news/sources",
		"message": "Hello from GetSources function!",
	})
}

// GetTypes <function>
// is used to handle get action of news controller which will return all news types.
// url: /v1/news/types
func (nc Controller) GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"method":  "v1/news/types",
		"message": "Hello from GetTypes function!",
	})
}
