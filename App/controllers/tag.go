package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gintest/App/models"
	"gintest/util"

	"github.com/astaxie/beego/validation"
)

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
        "error_code" : 0,
        "msg" : "查询成功",
        "data" : data,
    })
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	name := c.PostForm("name")
	state,_:= strconv.Atoi(c.DefaultPostForm("state","0"))
	
	valid := validation.Validation{}
    valid.Required(name, "name").Message("名称不能为空")
    valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	
	code,msg := util.SUCCESS,"添加成功"

	if ! valid.HasErrors() {
        if ! models.ExistTagByName(name) {
			models.AddTag(name, state)
		} else{
			code,msg = util.CUSTOM_ERROR,"标签名已存在"
		}

	}else{
		code,msg = util.CUSTOM_ERROR,valid.Errors[0].Message
	}

	c.JSON(200, gin.H{
		"error_code" : code,
		"msg" : util.GetMsg(code,msg),
		"data" : make(map[string]string),
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
	var data = make(map[string]interface{})
	id,_:= strconv.Atoi(c.Param("id"))
	name := c.PostForm("name")
	valid := validation.Validation{}
	state := -1;
	if arg := c.PostForm("state"); arg != "" {
		state,_= strconv.Atoi(arg)
        valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
    }
	valid.Required(id,"id").Message("ID不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code,msg:= util.SUCCESS,"编辑标签成功"
	if ! valid.HasErrors() {
        if models.ExistTagByID(id) {
			if(name != ""){
				data["name"] = name;
			}
			if(state != -1){
				data["state"] = state;
			}
			models.EditTag(id, data)
		} else{
			code,msg = util.CUSTOM_ERROR,"ID不存在";
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

//删除文章标签
func DeleteTag(c *gin.Context) {
	id,_:= strconv.Atoi(c.Param("id"))

	valid := validation.Validation{}
	valid.Required(id,"id").Message("ID不能为空")

	code,msg:= util.SUCCESS,"删除标签成功"
	if ! valid.HasErrors() {
        if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else{
			code,msg = util.CUSTOM_ERROR,"ID不存在";
		}
	}else{
		code,msg = util.CUSTOM_ERROR,valid.Errors[0].Message
	}

	c.JSON(200, gin.H{
		"code" : code,
		"msg" : util.GetMsg(code,msg),
		"data" :make(map[string]string),
	})
}