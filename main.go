package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "你测试成功了")
	})
	r.Run(":801")
}
