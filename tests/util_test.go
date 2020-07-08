package tests

import (
    "testing"
    "fmt"

    "gintest/util"
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

func TestJsonToStruct(t *testing.T) {
    var s = struct{
        Name string `json:"name"`
        Age int `json:"age"`
    }{}

    str  := `{"name":123","age":20}`
    err := util.JsonToStruct([]byte(str),&s)
    if err !=nil {
        fmt.Println("Umarshal failed:", err)
        return
    }

    fmt.Printf("name=%s,age=%d",s.Name,s.Age)
}

func TestRetData(t *testing.T) {
    data := make(map[string]interface{})
    data["name"]= "123"
    ret := util.DataReturn(true,"查询成功",data)
    fmt.Printf("状态返回：%t;错误信息返回:%s;数据=%v",ret.GetStatus(),ret.GetMsg(),ret["data"])
}
