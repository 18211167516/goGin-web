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

    testcase := []TestCase{
        {
            code:0,
            param:`name=5678999999`,
            errMsg:`查询成功`,
            method:"GET",
            desc:"验证查询成功",
            haveErr:true,
            showBody:true,
            bindStruct: &struct {
                Name string `json:"name" form:"name" binding:"required"` 
	            State int `json:"state" form:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags",
            content_type:"",
            ext1:1,
        },
    }

    for k,v:=range testcase {
        _,_,w := PerformRequest(v.method, "/api/v1/tags?"+v.param,v.content_type,bytes.NewBufferString(v.param))
        //assert.Contains(t, w.Body.String(),fmt.Sprintf("\"error_code\":%d",v.code))
        fmt.Println()
        fmt.Printf("第%d个测试用例：%s",k+1,v.desc)
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
        //assert.Equal(t, v.ext1,s.Data.(map[string]interface{})["total"],"条数不一致")
        assert.Equal(t, v.errMsg,s.Msg,"错误信息不一致")
        assert.Equal(t,v.code,s.Error_code,"错误码不一致")
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
            errMsg:`Key: 'Tag.Name' Error:Field validation for 'Name' failed on the 'required' tag`,
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
            errMsg:`Key: 'Tag.State' Error:Field validation for 'State' failed on the 'max' tag`,
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
            param:`name=123&state=1`,
            errMsg:`编辑成功`,
            method:"PUT",
            desc:"验证tag编辑成功",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags/1",
            content_type:"application/x-www-form-urlencoded",
        },
        {
            code:40001,
            param:`name=23245`,
            errMsg:`记录不存在`,
            method:"PUT",
            desc:"验证记录不存在",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags/10",
            content_type:"application/x-www-form-urlencoded",
        },
        {
            code:40001,
            param:`name=23245`,
            errMsg:`ID不能为空`,
            method:"PUT",
            desc:"验证ID参数为空",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags/0",
            content_type:"application/x-www-form-urlencoded",
        },
        {
            code:40001,
            param:`name=123456789011231313123`,
            errMsg:"名称最长为10字符",
            method:"PUT",
            desc:"验证字段name校验失败",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags/1",
            content_type:"application/x-www-form-urlencoded",
        },
        {
            code:40001,
            param:`name=56789&state=3`,
            errMsg:"状态只允许0或1",
            method:"PUT",
            desc:"验证字段state校验失败",
            haveErr:true,
            showBody:false,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags/1",
            content_type:"application/x-www-form-urlencoded",
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
            code:40001,
            param:``,
            errMsg:`记录不存在`,
            method:"DELETE",
            desc:"验证记录不存在",
            haveErr:true,
            showBody:true,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags/899",
            content_type:"",
        },
        {
            code:0,
            param:`name=111`,
            errMsg:`删除成功`,
            method:"DELETE",
            desc:"验证删除成功",
            haveErr:true,
            showBody:true,
            bindStruct: &struct {
                Name string `json:"name" from:"name" binding:"required"` 
	            State int `json:"state" from:"state" binding:"min=0,max=1"`
            }{},
            url:"/api/v1/tags/1",
            content_type:"",
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
    t.Run("TestRouteGetTags", testGetTags)
    //t.Run("TestAddTag", testAddTag)
    //t.Run("TestEditTag",testEditTag)
    //t.Run("TestDeleteTag",testDeleteTag)
}
 