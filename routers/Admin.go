package routers

import (
	"gin1/controllers/admin"
	"gin1/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middlewares.Init)
	{
		adminRouters.GET("/", admin.UserController{}.GetList)
		adminRouters.GET("/add", admin.UserController{}.Add)
		adminRouters.GET("/edit", admin.UserController{}.Edit)
		adminRouters.GET("/del", admin.UserController{}.Del)
	}
	articleRouters := r.Group("/admin/article")
	{
		articleRouters.GET("/", admin.ArticleController{}.GetList)
		articleRouters.GET("/add", admin.ArticleController{}.Add)
		articleRouters.GET("/edit/:id", admin.ArticleController{}.Edit)
		articleRouters.GET("/del", admin.ArticleController{}.Del)
	}

	//用户路由
	userRouters := r.Group("/admin/user")
	{
		userRouters.GET("/", admin.UserController{}.GetList)
		userRouters.GET("/add", admin.UserController{}.Add)
		userRouters.GET("/edit/:id", admin.UserController{}.Edit)
		userRouters.GET("/del", admin.UserController{}.Del)
		userRouters.Any("/doadd", admin.UserController{}.DoAdd)

	}

	//nav路由
	navRouters := r.Group("/admin/nav")
	{
		navRouters.Any("/", admin.NavController{}.GetList)
		navRouters.Any("/add", admin.NavController{}.Add)
		navRouters.Any("/edit/:id", admin.NavController{}.Edit)
		navRouters.Any("/del", admin.NavController{}.Del)
		navRouters.Any("/getinfo", admin.NavController{}.GetInfo)
	}

	//学生路由
	studentRouters := r.Group("/admin/student")
	{
		studentRouters.GET("/", admin.StudentController{}.GetList)
		studentRouters.GET("/add", admin.StudentController{}.Add)
		studentRouters.GET("/edit", admin.StudentController{}.Edit)
		studentRouters.GET("/del", admin.StudentController{}.Del)
	}
}
