package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
    jwt "github.com/dgrijalva/jwt-go"

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
            _, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = util.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = util.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
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