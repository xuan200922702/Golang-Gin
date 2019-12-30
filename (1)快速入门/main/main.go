package main

import "github.com/gin-gonic/gin"

func main()  {
	r :=gin.Default()    //实例化一个默认的gin示例
	r.GET("/", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"Blog":"www.zhangpengxuan.com",
			"wechat":"zhangpengxuan",
		})
		
	})
	r.Run(":8080")
}
