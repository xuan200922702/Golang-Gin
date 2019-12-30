package main

import "github.com/gin-gonic/gin"

func main(){
	r:=gin.Default()
	//传递数组
	r.GET("/", func(c *gin.Context) {
		c.JSON(200,c.QueryArray("media"))
	})
	//传递map
	r.GET("/map", func(c *gin.Context) {
		c.JSON(200,c.QueryMap("ids"))
	})
	r.Run(":8080")
}