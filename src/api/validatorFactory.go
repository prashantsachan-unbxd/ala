package api

import(
    "fmt"
    )

var typeMap = map[string]ApiValidator{
    "httpCode": HttpCodeChecker{},
}
func Init(){
    fmt.Println("calling init of validatorFactory")
}
func GetValidator(valType string, jsonData map[string]interface{}) ApiValidator{
    dummy := typeMap[valType]
    return dummy.NewInstance(jsonData)
}


