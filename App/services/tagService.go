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
}

type TagGetsContract interface {
	GetTags(maps map[string]interface{},c *gin.Context) map[string]interface{}
}

type TagAddContract interface {
	AddTag(maps map[string]interface{},data map[string]interface{}) map[string]interface{}
}

type TagExistContract interface {
	ExistTag(maps map[string]interface{}) bool
}

type TagService struct {
}

func (T TagService) GetTags(maps map[string]interface{},c *gin.Context) map[string]interface{} {
	data := make(map[string]interface{})
	data["lists"] = models.GetTags(util.GetPage(c), 10, maps)
	data["total"] = models.GetTagTotal(maps)
	return data
}

func (T TagService) AddTag(maps map[string]interface{},data map[string]interface{}) map[string]interface{} {
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