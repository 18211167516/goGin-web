package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gintest/App/models"
	"gintest/util"

	"github.com/astaxie/beego/validation"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id,_:= strconv.Atoi(c.Param("id"))

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	
	code,msg := 200,"查询成功"
	var data interface {}
	if ! valid.HasErrors() {
        if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
		} else{
			code,msg = 400,"文章不存在"
		}

	}else{
		code,msg = 500,valid.Errors[0].Message
	}

	c.JSON(code, gin.H{
        "code" : code,
        "msg" : msg,
        "data" : data,
    })
}

//获取多个文章
func GetArticles(c *gin.Context) {
    maps := make(map[string]interface{})
    data := make(map[string]interface{})
	valid := validation.Validation{}
	var state int = -1
    if arg := c.Query("state"); arg != "" {
        state,_= strconv.Atoi(arg)
        maps["state"] = state

        valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
    }

    var tagId int = -1
    if arg := c.Query("tag_id"); arg != "" {
        tagId,_= strconv.Atoi(arg)
        maps["tag_id"] = tagId

        valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
    } 

    data["lists"] = models.GetArticles(util.GetPage(c), 10, maps)
    data["total"] = models.GetArticleTotal(maps)

    c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "查询成功",
        "data" : data,
    })
}

//新增文章
func AddArticle(c *gin.Context) {
	data := make(map[string]interface{})
	
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	createdBy := c.PostForm("createdBy")
	state,_:= strconv.Atoi(c.DefaultPostForm("state","0"))
	tagId,_:= strconv.Atoi(c.DefaultPostForm("tag_id","0"))
	
	valid := validation.Validation{}
    valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
    valid.Required(title, "title").Message("标题不能为空")
    valid.Required(desc, "desc").Message("简述不能为空")
    valid.Required(content, "content").Message("内容不能为空")
    valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code,msg := 200,"添加成功"

	if ! valid.HasErrors() {
        if models.ExistTagByID(tagId) {
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			data["tag_id"] = tagId

			models.AddArticle(data)
			
		} else{
			code,msg = 400,"标签ID不存在"
		}

	}else{
		code,msg = 500,valid.Errors[0].Message
	}

	c.JSON(code, gin.H{
		"code" : code,
		"msg" : msg,
		"data" :make(map[string]string),
	})	
}

//修改文章
func EditArticle(c *gin.Context) {
}

//删除文章
func DeleteArticle(c *gin.Context) {
}