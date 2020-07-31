package controllers

import (
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"

	"gintest/util"
	"gintest/App/services"

)

type Tag struct {
	Name string `json:"name" form:"name" binding:"required"` 
	State int `json:"state" form:"state" binding:"min=0,max=1"`
}

type TagId struct{
	Id int `uri:"id" binding:"required"`
}

var tagServer services.TagContract

func init(){
	tagServer = &services.TagService{}
}
//获取多个文章标签
func GetTags(c *gin.Context) {
	tag := Tag{}

	err := c.ShouldBind(&tag)
	if err == nil {
		ret := tagServer.GetTags(tag,c)
		if !ret.GetStatus() {
			util.ApiAutoReturn(c,40001,"查询失败",nil)
		} else{
			util.ApiAutoReturn(c,0,"查询成功",ret["data"])
		}
	} else {
		util.ApiAutoReturn(c,500,err.Error(),nil)
	}


}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {

	tag := Tag{}

	code,msg := util.SUCCESS,"添加成功"
	//接收请求参数
	err := c.ShouldBind(&tag)
	if err != nil {
		code,msg = util.CUSTOM_ERROR,err.Error()
	}else{
		maps := make(map[string]interface{})
		maps["name"] = tag.Name
		ret := tagServer.AddTag(maps,util.StructToMap(tag))
		if !ret.GetStatus() {
			if ret.GetMsg()=="" {
				code,msg = util.CUSTOM_ERROR,"添加tag失败"
			} else {
				code,msg = util.CUSTOM_ERROR,ret.GetMsg()
			}
		}
	}

	util.ApiAutoReturn(c,code,msg,nil)
}

//修改文章标签
func EditTag(c *gin.Context) {
	var data = make(map[string]interface{})
	id := util.StringToInt(c.Param("id"))
	name := c.PostForm("name")


	valid := validation.Validation{}
	state := -1;
	if arg := c.PostForm("state"); arg != "" {
		state,_= strconv.Atoi(arg)
        valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
    }
	valid.Required(id,"id").Message("ID不能为空")
	valid.MaxSize(name, 10, "name").Message("名称最长为10字符")

	code,msg:= util.SUCCESS,"编辑成功"
	if ! valid.HasErrors() {
		if(name != ""){
			data["name"] = name;
		}
		if(state != -1){
			data["state"] = state;
		}
		ret := tagServer.EditTag(id,data)
		if !ret.GetStatus(){
			code,msg = util.CUSTOM_ERROR,ret.GetMsg()
		}
	}else{
		code,msg = util.CUSTOM_ERROR,valid.Errors[0].Message
	}

	util.ApiAutoReturn(c,code,msg,nil)
	
}

//删除文章标签
func DeleteTag(c *gin.Context) {

	tag := TagId{}
	/* id := util.StringToInt(c.Param("id"))

	valid := validation.Validation{}
	valid.Required(id,"id").Message("ID不能为空") */
	code,msg:= util.SUCCESS,"删除成功"
	err := c.ShouldBindUri(&tag)
	if err != nil {
		code,msg = util.CUSTOM_ERROR,err.Error()
	}else{
		maps := make(map[string]interface{})
		maps["id"] = tag.Id
		ret := tagServer.DeleteTag(maps)
        if !ret.GetStatus() {
			if ret.GetMsg()=="" {
				code,msg = util.CUSTOM_ERROR,"删除失败"
			} else {
				code,msg = util.CUSTOM_ERROR,ret.GetMsg()
			}
		}
	}

	
	/* if ! valid.HasErrors() {
		maps := make(map[string]interface{})
		maps["id"] = id
		ret := tagServer.DeleteTag(maps)
        if !ret.GetStatus() {
			if ret.GetMsg()=="" {
				code,msg = util.CUSTOM_ERROR,"删除失败"
			} else {
				code,msg = util.CUSTOM_ERROR,ret.GetMsg()
			}
		}
	}else{
		code,msg = util.CUSTOM_ERROR,valid.Errors[0].Message
	} */

	util.ApiAutoReturn(c,code,msg,nil)
}