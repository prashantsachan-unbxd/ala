package api

import(
    )
const VALIDATOR_TYPE_HTTPCODE ="httpCode"
var typeMap = map[string]ApiValidator{
    VALIDATOR_TYPE_HTTPCODE: HttpCodeChecker{},
}
func GetValidator(valType string, jsonData map[string]interface{}) ApiValidator{
    dummy := typeMap[valType]
    return dummy.NewInstance(jsonData)
}


