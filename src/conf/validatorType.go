package conf

import(
    "reflect"
    "api"
//    "fmt"
    )

var typeMap = map[string]reflect.Type{
    "httpCode": reflect.TypeOf(api.HttpCodeChecker{}),
}

func GetValidator(typeStr string) api.ApiValidator{
    t:= typeMap[typeStr]
    v:= reflect.New(t).Elem()
    i:= v.Interface()
    validator:= i.(api.ApiValidator)
    return validator
}
func GetValidatorType(v api.ApiValidator) string{
    t := reflect.TypeOf(v)
    for k,v:= range typeMap{
        if v == t{
            return k
        }
    }
    return ""
}
