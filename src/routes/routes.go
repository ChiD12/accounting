package routes

import (
	"github.com/gin-gonic/gin"
)

// NewRouter <function>
// is used to create a GIN engine instance where all controller and routes will be placed
func MakeRouters(router *gin.Engine) *gin.Engine {

	// endpoints
	v1 := router.Group("v1")
	{
		news := v1.Group("auth")
		{
			controllers := Controller{}

			news.GET("/", controllers.GetUsers)
			news.GET("/sources", controllers.Login)
			news.GET("/types", controllers.Login)
		}
	}

	return router
}
