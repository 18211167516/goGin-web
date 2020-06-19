package controllers

import (
	"github.com/gin-gonic/gin"
	"gintest/App/models"
	"gintest/util"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type user struct{
	
}


//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	state:= c.Query("state");

    maps := make(map[string]interface{})
    data := make(map[string]interface{})

    if name != "" {
        maps["name"] = name
    }

    if state != "" {
        maps["state"] = state
    }

    data["lists"] = models.GetTags(util.GetPage(c), 10, maps)
    data["total"] = models.GetTagTotal(maps)

    c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "查询成功",
        "data" : data,
    })
}

//新增文章标签
func AddTag(c *gin.Context) {
	name := c.PostForm("name")
	state,_:= strconv.Atoi(c.DefaultPostForm("state","0"))
	code := 200

	valid := validation.Validation{}
    valid.Required(name, "name").Message("名称不能为空")
    valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	if ! valid.HasErrors() {
        if ! models.ExistTagByName(name) {
			models.AddTag(name, state)
			c.JSON(code, gin.H{
				"code" : code,
				"msg" : "成功",
				"data" : make(map[string]string),
			})
		} else{
			c.JSON(400, gin.H{
				"code" : 400,
				"msg" : "姓名已存在",
				"data" : make(map[string]string),
			})
		}

	}else{
		c.JSON(500, gin.H{
			"code" : code,
			"msg" : valid.Errors[0].Message,
			"data" :make(map[string]string),
		})	
	}
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}