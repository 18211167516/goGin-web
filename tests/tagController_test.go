package tests
import (
    "testing"
    "fmt"
    "bytes"

    "gintest/util"

    "github.com/stretchr/testify/assert"
)


func testRouteAuto(t *testing.T){
    _,_,w := PerformRequest("GET", "/auth?username=test&password=test123456","application/json; charset=utf-8", nil)

    fmt.Printf("auth 路由返回%s",w.Body.String())
    fmt.Println()
}


/* GET  /api/v1/tags */
func testGetTags(t *testing.T){

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
        _,_,w := PerformRequest("GET", "/api/v1/tags?"+v.param,"application/json; charset=utf-8",nil)
        assert.Contains(t, w.Body.String(),fmt.Sprintf("\"error_code\":%d",v.code))
        fmt.Println()
    }
}

/* POST  /api/v1/tags */
func testAddTag(t *testing.T){

    testcase := []TestCase{
        {
            code:0,
            param:`{"name":"5678999999"}`,
            errMsg:`添加成功`,
            method:"POST",
            desc:"验证添加成功",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
        {
            code:40001,
            param:`{"name":"5678999999"}`,
            errMsg:`记录已存在`,
            method:"POST",
            desc:"验证记录已存在",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
        {
            code:40001,
            param:`{"age":"12323"}`,
            errMsg:`Key: 'Addtag.Name' Error:Field validation for 'Name' failed on the 'required' tag`,
            method:"POST",
            desc:"验证字段name校验失败",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
        {
            code:40001,
            param:`{"name":"56789","state":3}`,
            errMsg:`Key: 'Addtag.State' Error:Field validation for 'State' failed on the 'max' tag`,
            method:"POST",
            desc:"验证字段state校验失败",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
    }

    for k,v:=range testcase {
        _,_,w := PerformRequest(v.method,v.url,v.content_type, bytes.NewBufferString(v.param))
        //assert.Contains(t, w.Body.String(),fmt.Sprintf("\"error_code\":%d",v.code))
        
        fmt.Printf("第%d个测试用例：%s",k+1,v.desc)
        fmt.Println()
        if v.showBody {
            fmt.Printf("接口返回%s",w.Body.String())
            fmt.Println()
        }
        s := struct{
            Error_code int `json:"error_code"`
            Msg  string `json:"msg"`
            Data  interface{} `json:"data"`
        }{}
        err := util.JsonToStruct([]byte(w.Body.String()),&s)
        assert.NoError(t,err)
        //fmt.Printf("msg=%s,errMsg=%s",s.Msg,v.errMsg)
        assert.Equal(t, v.errMsg,s.Msg,"错误信息不一致")
        assert.Equal(t,v.code,s.Error_code,"错误码不一致")
    }

}

/*PUT /api/v1/tags/:id*/
func testEditTag(t *testing.T) {
    testcase := []TestCase{
        {
            code:0,
            param:`{"name":"5678999999"}`,
            errMsg:`添加成功`,
            method:"POST",
            desc:"验证添加成功",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
        {
            code:40001,
            param:`{"name":"5678999999"}`,
            errMsg:`记录已存在`,
            method:"POST",
            desc:"验证记录已存在",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
        {
            code:40001,
            param:`{"age":"12323"}`,
            errMsg:`Key: 'Addtag.Name' Error:Field validation for 'Name' failed on the 'required' tag`,
            method:"POST",
            desc:"验证字段name校验失败",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
        {
            code:40001,
            param:`{"name":"56789","state":3}`,
            errMsg:`Key: 'Addtag.State' Error:Field validation for 'State' failed on the 'max' tag`,
            method:"POST",
            desc:"验证字段state校验失败",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
    }

    for k,v:=range testcase {
        _,_,w := PerformRequest(v.method,v.url,v.content_type, bytes.NewBufferString(v.param))
        //assert.Contains(t, w.Body.String(),fmt.Sprintf("\"error_code\":%d",v.code))
        
        fmt.Printf("第%d个测试用例：%s",k+1,v.desc)
        fmt.Println()
        if v.showBody {
            fmt.Printf("接口返回%s",w.Body.String())
            fmt.Println()
        }
        s := struct{
            Error_code int `json:"error_code"`
            Msg  string `json:"msg"`
            Data  interface{} `json:"data"`
        }{}
        err := util.JsonToStruct([]byte(w.Body.String()),&s)
        assert.NoError(t,err)
        //fmt.Printf("msg=%s,errMsg=%s",s.Msg,v.errMsg)
        assert.Equal(t, v.errMsg,s.Msg,"错误信息不一致")
        assert.Equal(t,v.code,s.Error_code,"错误码不一致")
    }
}

/*DELETE /api/v1/tags/:id*/
func testDeleteTag(t *testing.T) {
    testcase := []TestCase{
        {
            code:0,
            param:`{"name":"5678999999"}`,
            errMsg:`添加成功`,
            method:"POST",
            desc:"验证添加成功",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
        {
            code:40001,
            param:`{"name":"5678999999"}`,
            errMsg:`记录已存在`,
            method:"POST",
            desc:"验证记录已存在",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
        {
            code:40001,
            param:`{"age":"12323"}`,
            errMsg:`Key: 'Addtag.Name' Error:Field validation for 'Name' failed on the 'required' tag`,
            method:"POST",
            desc:"验证字段name校验失败",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
        {
            code:40001,
            param:`{"name":"56789","state":3}`,
            errMsg:`Key: 'Addtag.State' Error:Field validation for 'State' failed on the 'max' tag`,
            method:"POST",
            desc:"验证字段state校验失败",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"application/json",
        },
    }

    for k,v:=range testcase {
        _,_,w := PerformRequest(v.method,v.url,v.content_type, bytes.NewBufferString(v.param))
        //assert.Contains(t, w.Body.String(),fmt.Sprintf("\"error_code\":%d",v.code))
        
        fmt.Printf("第%d个测试用例：%s",k+1,v.desc)
        fmt.Println()
        if v.showBody {
            fmt.Printf("接口返回%s",w.Body.String())
            fmt.Println()
        }
        s := struct{
            Error_code int `json:"error_code"`
            Msg  string `json:"msg"`
            Data  interface{} `json:"data"`
        }{}
        err := util.JsonToStruct([]byte(w.Body.String()),&s)
        assert.NoError(t,err)
        //fmt.Printf("msg=%s,errMsg=%s",s.Msg,v.errMsg)
        assert.Equal(t, v.errMsg,s.Msg,"错误信息不一致")
        assert.Equal(t,v.code,s.Error_code,"错误码不一致")
    }
}

func TestTagAll(t *testing.T)  {
    //t.Run("TestRouteGetTags", testGetTags)
    t.Run("TestAddTag", testAddTag)
    //t.Run("TestEditTag",testEditTag)
}
 