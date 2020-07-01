package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_"gintest/docs"
	"gintest/App/controllers"
	_"gintest/App/middleware"
	_"gintest/config"
)

func InitRouter() *gin.Engine{
	//同时向文件和控制台写入日志
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	//r.LoadHTMLGlob(config.AppSetting.Template)

	/* r.GET("/index", func(c *gin.Context) {
		// 子目录的模版文件，需要加上目录名，例如：posts/index.tmpl
		c.HTML(200, "index/index", gin.H{
			"title": "Posts",
		})
	}) */
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/auth",controllers.GetAuth)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(middleware.JWT())
    {
        //获取标签列表
        apiv1.GET("/tags", controllers.GetTags)
        //新建标签
        apiv1.POST("/tags", controllers.AddTag)
        //更新指定标签
        apiv1.PUT("/tags/:id", controllers.EditTag)
        //删除指定标签
		apiv1.DELETE("/tags/:id", controllers.DeleteTag)
		
		//获取文章列表
		apiv1.GET("/articles", controllers.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", controllers.GetArticle)
        //新建文章
        apiv1.POST("/articles", controllers.AddArticle)
        //更新指定文章
        apiv1.PUT("/articles/:id", controllers.EditArticle)
        //删除指定文章
		apiv1.DELETE("/articles/:id", controllers.DeleteArticle)
	}
	

	return r;
}
