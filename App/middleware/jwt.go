package middleware

import (
    "time"
    "net/http"

    "github.com/gin-gonic/gin"

    "gintest/util"
)

func JWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        var code int
        var data interface{}

        code = util.SUCCESS
        token := c.Query("token")
        if token == "" {
            code = util.ERROR_AUTH_CHECK_TOKEN_EMPTY
        } else {
            claims, err := util.ParseToken(token)
            if err != nil {
                code = util.ERROR_AUTH_CHECK_TOKEN_FAIL
            } else if time.Now().Unix() > claims.ExpiresAt {
                code = util.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
            }
        }

        if code != util.SUCCESS {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code" : code,
                "msg" : util.GetMsg(code),
                "data" : data,
            })

            c.Abort()
            return
        }

        c.Next()
    }
}