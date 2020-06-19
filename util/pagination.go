package util

import (
    "github.com/gin-gonic/gin"
	"gintest/config"
	"strconv"
)

func GetPage(c *gin.Context) int {
    result := 0
	page,_:= strconv.Atoi(c.DefaultQuery("page","0"))
    if page > 0 {
        result = (page - 1) * config.PageSize
    }
    return result
}