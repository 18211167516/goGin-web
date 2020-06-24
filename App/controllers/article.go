package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"

	"gintest/App/models"
	"gintest/util"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id,_:= strconv.Atoi(c.Param("id"))

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	
	code,msg := util.SUCCESS,"查询成功"
	var data interface {}
	if ! valid.HasErrors() {
        if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
		} else{
			code,msg = util.CUSTOM_ERROR,"文章不存在"
		}

	}else{
		code,msg = util.CUSTOM_ERROR,valid.Errors[0].Message
	}

	c.JSON(200, gin.H{
        "error_code" : code,
        "msg" : util.GetMsg(code,msg),
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
        "error_code" : 0,
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

	code,msg := 0,"添加成功"

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
			code,msg = util.CUSTOM_ERROR,"标签ID不存在"
		}

	}else{
		code,msg = util.CUSTOM_ERROR,valid.Errors[0].Message
	}

	c.JSON(200, gin.H{
		"error_code" : code,
		"msg" : util.GetMsg(code,msg),
		"data" :make(map[string]string),
	})	
}

//修改文章
func EditArticle(c *gin.Context) {
	id,_:= strconv.Atoi(c.Param("id"))
	tagId,_ := strconv.Atoi(c.PostForm("tag_id"))
	title := c.PostForm("title")
    desc := c.PostForm("desc")
    content := c.PostForm("content")
	modifiedBy := c.PostForm("modified_by")
	
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
    valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
    valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
    valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
    valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	state := -1;
	if arg := c.PostForm("state"); arg != "" {
        state,_= strconv.Atoi(arg)
        valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
    }
	code,msg := 0,"编辑成功"

	if tagId>0 {
		if !models.ExistTagByID(tagId){
			valid.SetError("tagId","标签ID不存在")
		}
	}

	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data := make(map[string]interface{})
			
			data["modified_by"] = modifiedBy
			if title != "" {
				data["title"] = title
			}
			
			if desc != "" {
				data["desc"] = desc
			}
			
			if content != "" {
				data["content"] = content
			}
			
			if state !=-1 {
				data["state"] = state
			}

			if tagId > 0 {
				data["tag_id"] = tagId
			}
			
			models.EditArticle(id,data)
			
		} else{
			code,msg = util.CUSTOM_ERROR,"文章不存在"
		}
	}else{
		code,msg = util.CUSTOM_ERROR,valid.Errors[0].Message
	}

	c.JSON(200, gin.H{
		"error_code" : code,
		"msg" : util.GetMsg(code,msg),
		"data" :make(map[string]string),
	})	
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))

	vaild := validation.Validation{}
	vaild.Required(id,"id").Message("ID不能为空")

	code,msg := 0,"删除成功"

	if !vaild.HasErrors() {
		//
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id);
		}else {
			code,msg = util.CUSTOM_ERROR,"文章ID不存在"
		}
	}else{
		code,msg = util.CUSTOM_ERROR,vaild.Errors[0].Message
	}

	c.JSON(200,gin.H{
		"error_code" : code,
		"msg" : util.GetMsg(code,msg),
		"data" :make(map[string]string),
	})
}