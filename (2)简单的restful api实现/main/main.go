package main

import "github.com/gin-gonic/gin"

type User struct {
	ID  uint64
	Name  string
}

func main(){
	users :=[]User{{ID: 123, Name:"张三"},{ID:456,Name:"李四"}}
	r :=gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200,users)
	})
	r.POST("users", func(context *gin.Context) {
		//创建一个用户
	})
	r.DELETE("/users/123", func(context *gin.Context) {
		//删除ID为123的用户
	})
	r.PUT("/users/123", func(context *gin.Context) {
		//更新id为123的用户
	})
	Handle(r, []string{"GET", "POST"}, "/", func(c *gin.Context) {
		//同时注册GET、POST请求方法
	})

	r.Run(":8080")
}
//一次注册所有的HTTP Method的方法
//Gin提供了Any方法，可以一次性注册以上这些HTTP Method方法
func Handle(r *gin.Engine, httpMethods []string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	var routes gin.IRoutes
	for _, httpMethod := range httpMethods {
		routes = r.Handle(httpMethod, relativePath, handlers...)
	}
	return routes
}



