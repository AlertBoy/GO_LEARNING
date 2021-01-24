package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	userGroup := r.Group("/user")
	userGroup.Use(gin.BasicAuth(map[string]string{
		"cly":  "12345",
		"cly2": "123456",
	}))
	// 获取文件
	userGroup.GET("/files/:fileName", func(c *gin.Context) {
		_, ok := c.Params.Get("fileName")
		if !ok {
			c.Err()
		}
		c.File("D:\\Log_backup_20201007.zip")
	})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "你测试成功了")
	})
	r.Run(":801")
}
