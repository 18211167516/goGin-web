package util

import (

)

func ApiReturn(code int,msg string,data interface{}) {

}

func DataReturn(status bool,msg string,data interface{}) map[string]interface{}{
	result := make(map[string]interface{})
	result["status"] = status
	result["msg"]    = msg
	result["data"]   = data
	return result
}