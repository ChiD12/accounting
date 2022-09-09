package src

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"example.com/accounting/src/routes"
)

func NewServer() *gin.Engine {

	router := gin.New()

	// var envVars = utils.GetEnvVars()

	// if envVars.DebugMode {
	// 	gin.SetMode(gin.DebugMode)
	// } else {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	// middlewares
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.Default())

	// static files serving
	router.Static("/images", "./images")

	routes.MakeRouters(router)
	return router

}
