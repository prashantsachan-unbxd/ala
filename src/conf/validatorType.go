package conf

import(
    "reflect"
    "api"
    "fmt"
    )

var typeMap = map[string]reflect.Type{
    "httpCode": reflect.TypeOf(api.HttpCodeChecker{}),
}

func GetValidator(typeStr string) api.ApiValidator{
    t:= typeMap[typeStr]
    fmt.Println("reflect.Type value t:", t)
    v:= reflect.New(t).Elem()
    fmt.Println("element instance : ", v)
    i:= v.Interface()
    fmt.Println("interface{} value:", i)
    validator:= i.(api.ApiValidator)
    fmt.Println("ApiValidator value:", validator)
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
