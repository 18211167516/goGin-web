package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"

	"gintest/util"
	"gintest/App/models"
)

type auth struct {
    Username string `valid:"Required; MaxSize(50)"`
    Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
    username := c.Query("username")
    password := c.Query("password")

    valid := validation.Validation{}
    a := auth{Username: username, Password: password}
    ok, _ := valid.Valid(&a)

    data := make(map[string]interface{})
    code,msg := util.INVALID_PARAMS,"成功"
    if ok {
        isExist := models.CheckAuth(username, password)
        if isExist {
            token, err := util.GenerateToken(username, password)
            if err != nil {
                code = util.ERROR_AUTH_TOKEN
            } else {
                data["token"] = token
                
                code = util.SUCCESS
            }

        } else {
            code = util.ERROR_AUTH
        }
    } else {
        code,msg = util.CUSTOM_ERROR,valid.Errors[0].Message
    }

    c.JSON(200, gin.H{
        "error_code" : code,
        "msg" : util.GetMsg(code,msg),
        "data" : data,
    })
}