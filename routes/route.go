package routes

import(
	"github.com/gin-gonic/gin"
	"gintest/App/controllers"
)

func InitRouter() *gin.Engine{
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiv1 := r.Group("/api/v1")
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
