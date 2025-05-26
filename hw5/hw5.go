package hw5

import "github.com/gin-gonic/gin"

// Ping handler returns a simple "pong" response.
func Ping(c *gin.Context) {
	c.String(200, "pong")
}

// AppConf sets up the Gin router for hw5.
func AppConf() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", Ping)
	return router
}
