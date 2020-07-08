package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
type Gin struct {
    C *gin.Context
}

type M map[string]interface{}

type APIH struct{
	Error_code int `json:"error_code"`
	Msg  string `json:"msg"`
	Data  interface{} `json:"data"`
}

type RetData struct{
	Status bool `json:"status"`
	Msg  string `json:"msg"`
	Data  interface{} `json:"data"`
}

func NewGin(c *gin.Context) Gin {
	return Gin{c}
}

func NewRetData(status bool,msg string,data interface{}) RetData{
	return RetData{status,msg,data}
}

func NewAPIH() APIH {
	return APIH{}
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


func DataReturn(status bool,msg string,data interface{}) M{
	result :=M{
		"status" : status,
		"msg" : msg,
		"data" : data,
	}
	return result
}

func (m M) GetStatus() bool{
	return m["status"].(bool)
}

func (m M) GetMsg() string{
	return m["msg"].(string)
}