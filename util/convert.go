package util

import (
    "reflect"
    "strconv"
)

func StringToInt(str string) int{
    variable,_  := strconv.Atoi(str)
    return variable
}

func StructToMap(obj interface{}) map[string]interface{}{

	t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)
 
    var data = make(map[string]interface{})
    for i := 0; i < t.NumField(); i++ {
        data[t.Field(i).Name] = v.Field(i).Interface()
    }
    return data
}