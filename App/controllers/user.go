package controllers

import (
	"github.com/gin-gonic/gin"
)

type user struct{
	
}


//获取多个文章标签
func GetTags(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "1233",
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}