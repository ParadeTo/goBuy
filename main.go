package main

import (
	"github.com/gin-gonic/gin"
	"goBuy/config"
	"fmt"
)

func main() {
	config.Load()

	r := gin.New()

	// data output middleware
	r.Use(func(c *gin.Context) {
		c.Next()

		data, _ := c.Get("data")
		status := c.Writer.Status()

		if status == 200 {
			c.JSON(200, gin.H{
				"code": 200,
				"msg": "ok",
				"data": data,
			})
		} else {
			c.JSON(status, gin.H{
				"code": -1,
				"msg": c.Err(),
			})
		}
	})

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		c.Set("data", gin.H{"name": "youxingzhi"})
	})

	r.Use(func(context *gin.Context) {
		fmt.Println("This will not execute if path (/test) match successfully")
	})

	r.Run(":" + config.Conf.Port)
}
