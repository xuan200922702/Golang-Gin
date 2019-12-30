package main

import "github.com/gin-gonic/gin"
//Gin的路由采用的是httprouter，所以它的路由参数的定义和httprouter也是一样的。
//Gin路径中的匹配都是字符串，它是不区分数字、字母和汉字的，都匹配
//通配符重复了，路由必须要唯一。Gin内部使用的路由是httprouter
func main(){
	r :=gin.Default()
	//r.GET("/users/:id", func(c *gin.Context) {
	//	id :=c.Param("id")
	//	c.String(200,"The user id is %s",id)
	//})
	//还有一种不常用的就是*号类型的参数，表示匹配所有
	///users/*id和/users这两个路由是不冲突的，可以被Gin注册
	//自动重定向的原理，得益于gin.RedirectTrailingSlash 等于true的配置。
	// 如果我们把它改为false就不会自动重定向了。
	r.RedirectTrailingSlash=false
	r.GET("/users/*id", func(c *gin.Context) {
		id :=c.Param("id")
		c.String(200,"The user id is %s",id)
	})
	//r.GET("/users", func(c *gin.Context) {
	//	c.String(200,"这是真正的/users")
	//})
	r.Run(":8080")
	//我们可以很灵活的实现我们的API，并且从路径中获取相应的参数进行操作。
	// 对于*号参数，不建议使用，因为匹配的太多，会导致我们自己搞不清楚哪些路由被注册了
}
