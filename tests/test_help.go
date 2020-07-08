package tests

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	"gintest/routes"
)

type TestCase struct{
    code int //状态码
    param string //参数
    method string //请求类型
    desc string //描述
    haveErr bool //是否有错误
    showBody bool //是否展示返回
    bindStruct interface{} //带绑定的结构体
    errMsg string //错误信息
    url string //链接
    content_type string //
}

func PerformRequest(mothod,url,content_type string,body io.Reader) (c *gin.Context,r *http.Request ,w *httptest.ResponseRecorder){
    router := routes.InitRouter()
    w = httptest.NewRecorder()
    c, _ = gin.CreateTestContext(w)
    r = httptest.NewRequest(mothod, url, body)
    c.Request = r
    c.Request.Header.Set("Content-Type", content_type)
    router.ServeHTTP(w,r)
    return
}

func Server(w *httptest.ResponseRecorder, r *http.Request){
    router := routes.InitRouter()
	router.ServeHTTP(w,r)
}