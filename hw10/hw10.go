package hw10

import (
	"github.com/gin-gonic/gin"
	"log"
)

type DataStr struct {
	Data string `json:"data"`
}

func AppConf() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", Hello)
	router.GET("/data", Data)
	router.GET("/bad", BadRequest)
	return router
}

func AppStart() {
	r := AppConf()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func Hello(c *gin.Context) {
	c.String(200, "Hello!")
}

func Data(c *gin.Context) {
	d := DataStr{Data: "very useful data"}
	c.JSON(200, d)
}

func BadRequest(c *gin.Context) {
	c.JSON(400, gin.H{
		"error": "bad request",
	})
}
