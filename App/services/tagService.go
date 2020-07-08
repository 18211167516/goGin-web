package services

import (
	"github.com/gin-gonic/gin"

	"gintest/App/models"
	"gintest/util"
)

type TagContract interface {
	TagGetsContract //GET tag
	TagAddContract  //Add tag
	TagExistContract //exist tag
	TagDeleteContract//delete tag
	TagEditContract//edit tag
}

type TagGetsContract interface {
	GetTags(maps map[string]interface{},c *gin.Context) util.M
}

type TagAddContract interface {
	AddTag(maps map[string]interface{},data map[string]interface{}) util.M
}

type TagExistContract interface {
	ExistTag(maps map[string]interface{}) bool
}

type TagDeleteContract interface {
	DeleteTag(maps map[string]interface{}) util.M
}

type TagEditContract interface {
	EditTag(id int,data interface{}) util.M
}

type TagService struct {
}

func (T TagService) GetTags(maps map[string]interface{},c *gin.Context) util.M {
	data := make(map[string]interface{})
	data["lists"] = models.GetTags(util.GetPage(c), 10, maps)
	data["total"] = models.GetTagTotal(maps)
	return util.DataReturn(true,"查询成功",data)
}

func (T TagService) AddTag(maps map[string]interface{},data map[string]interface{}) util.M {
	//先通过maps查询数据是否存在
	if(T.ExistTag(maps)){
		return util.DataReturn(false,"记录已存在",nil)
	}else{
		return util.DataReturn(models.AddTag(data),"",nil)
	}
}

func (T TagService) ExistTag(maps map[string]interface{}) bool{
	
	return models.ExistTagByMaps(maps)
}

func (T TagService) DeleteTag(maps map[string]interface{}) util.M {
	if(T.ExistTag(maps)){
		return util.DataReturn(models.DeleteTag(maps),"",nil)
	}else{
		return util.DataReturn(false,"记录不存在",nil)
	}
}

func (T TagService) EditTag(id int,data interface{}) util.M{
	if(models.ExistTagByID(id)){
		return util.DataReturn(models.EditTag(id,data),"",nil)
	}else{
		return util.DataReturn(false,"记录不存在",nil)
	}
}