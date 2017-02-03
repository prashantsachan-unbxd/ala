package api

import(
    "reflect"
    )

var typeMap = map[string]reflect.Type{
    "httpCode": reflect.TypeOf(HttpCodeChecker{}),
}

func GetValidator(typeStr string) ApiValidator{
    t:= typeMap[typeStr]
    v:= reflect.New(t).Elem()
    i:= v.Interface()
    validator:= i.(ApiValidator)
    return validator
}
func GetValidatorType(v ApiValidator) string{
    t := reflect.TypeOf(v)
    for k,v:= range typeMap{
        if v == t{
            return k
        }
    }
    return ""
}
