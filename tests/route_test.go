package tests
import (
    "testing"
    "fmt"
    "net/http"
    "net/http/httptest"
    //"bytes"
    "io"

	"gintest/routes"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func performRequest(mothod,url string,body io.Reader) (r *http.Request ,w *httptest.ResponseRecorder){
    router := routes.InitRouter()
    r = httptest.NewRequest(mothod, url, body)
    w = httptest.NewRecorder()
    router.ServeHTTP(w,r)
    return
}

func Server(w *httptest.ResponseRecorder, r *http.Request){
    router := routes.InitRouter()
	router.ServeHTTP(w,r)
}
func TestRouteAuto(t *testing.T){
    gin.SetMode(gin.TestMode)

    _,w := performRequest("GET", "/auth?username=test&password=test123456", nil)

    fmt.Printf("auth 路由返回%s",w.Body.String())
    fmt.Println()
}

func TestRouteGetTags(t *testing.T){
    gin.SetMode(gin.TestMode)

    testcase := []struct{
        code int //状态码
        param string //参数
        //bindStruct interface{} //带绑定的结构体
        errMsg string //错误信息
    }{
        {
            code:0,
            param:"name=123",
            errMsg:"",
        },
        {
            code:40001,
            param:"",
            errMsg:"",
        },
    }

    for _,v:=range testcase {
        _,w := performRequest("GET", "/api/v1/tags?"+v.param,nil)
        //fmt.Printf("auth 路由返回%s",w.Body.String())
        assert.Contains(t, w.Body.String(),fmt.Sprintf("\"error_code\":%d",v.code))
        fmt.Println()
        /* if c.haveErr {
            //err := ctx.ShouldBindJSON(c.bindStruct)
            assert.Error(t, err)
            assert.Equal(t, c.errMsg, err.Error())
        } else {
            assert.NoError(t, ctx.ShouldBindJSON(c.bindStruct))
        } */
    }

    
}
 