package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	fmt.Print("asdfsdf")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}