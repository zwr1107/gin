package main

import (
	"gin1/routers"
	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
	Id      string
}

type User struct {
	Username string
	Password string
	Address  string
}

// 中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("request", "中间件")
		c.Next()
	}
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	//定义模板
	r.LoadHTMLGlob("templates/**/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	//配置全局中间件
	//r.Use(MiddleWare())

	routers.ApiRoutersInit(r)
	routers.WebRoutersInit(r)
	routers.AdminRoutersInit(r)

	// 启动HTTP 服务，默认在0.0.0.0:8080 启动服务
	r.Run()
}
