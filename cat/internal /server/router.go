package server

import (
	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = true

	// Serve static files to frontend if server is started in production environment
	router.Use(static.Serve("/", static.LocalFile("./assets/build", true)))

	// Create API route group
	api := router.Group("/api")
	{
		api.POST("/signup", signUp)
		api.POST("/signin", signIn)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

// func setRouter() *gin.Engine {
// 	// Creates default gin router with Logger and Recovery middleware already attached
// 	router := gin.Default()

// 	// Create API route group
// 	api := router.Group("/api")
// 	{
// 		// Add /hello GET route to router and define route handler function
// 		api.GET("/hello", func(ctx *gin.Context) {
// 			ctx.JSON(200, gin.H{"msg": "world"})
// 		})
// 	}

// 	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

// 	return router
// }
