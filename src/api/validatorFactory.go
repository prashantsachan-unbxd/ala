package api

import(
    "errors"
    )
const VALIDATOR_TYPE_HTTPCODE ="httpCode"
var typeMap = map[string]ApiValidator{
    VALIDATOR_TYPE_HTTPCODE: HttpCodeChecker{},
}
func GetValidator(valType string, jsonData map[string]interface{}) (ApiValidator, error){
    dummy, ok:= typeMap[valType]
    if !ok{
        return nil, errors.New("invalid validator type: "+valType)
    }
    return dummy.NewInstance(jsonData), nil
}


