package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
type Gin struct {
    C *gin.Context
}

func NewGin(c *gin.Context) Gin {
	return Gin{c}
}

func ApiAutoReturn(c *gin.Context,code int,msg string,data interface{}) {
	g := NewGin(c)

	fmt.Printf("%T",g)
	g.ApiReturn(code,msg,data)
	//g.Return(code,msg,data)
}

func (g *Gin) ApiReturn(code int,msg string,data interface{}) {
	Headtype := g.C.GetHeader("Response-type")
	obj,httCode := gin.H{
		"error_code" : code,
		"msg" : GetMsg(code,msg),
		"data" : data,
	},200

	switch Headtype{
		case "json":
			g.C.JSON(httCode,obj)
		case "xml":
			g.C.XML(httCode,obj)
		default:
			g.C.JSONP(httCode,obj)
	}
}


func DataReturn(status bool,msg string,data interface{}) map[string]interface{}{
	result := make(map[string]interface{})
	result["status"] = status
	result["msg"]    = msg
	result["data"]   = data
	return result
}