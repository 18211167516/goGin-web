package tests

import (
    "testing"
    "fmt"
    "net/http"
    "net/http/httptest"

    "gintest/util"

    "github.com/gin-gonic/gin"
)

func TestStringToInt(t *testing.T) {
    maps := make(map[string]string)
    maps["a"] = "1"
    maps["b"] = ""
    maps["c"] = "str"
    for k,v :=range(maps) {
        t.Logf("key = %s,val=%d",k,util.StringToInt(v))
    } 
}

func TestApiReturn(t *testing.T) {

    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request, _ = http.NewRequest("GET", "/aaa", nil)
    c.Request.Header.Set("Response-type", "xml")

    util.ApiAutoReturn(c,40001,"自定义错误","123")
    fmt.Printf("%s",w.Body.String())
    fmt.Println()
}

