package api

import(
    "errors"
    )
const VALIDATOR_TYPE_HTTPCODE ="httpCode"
const VALIDATOR_TYPE_RE_HTTPCODE = "REhttpCode"
var typeMap = map[string]ApiValidator{
    VALIDATOR_TYPE_HTTPCODE: HttpCodeChecker{},
    VALIDATOR_TYPE_RE_HTTPCODE: &RuleEngineValidator{},
}
func GetValidator(valType string, jsonData map[string]interface{}) (ApiValidator, error){
    dummy, ok:= typeMap[valType]
    if !ok{
        return nil, errors.New("invalid validator type: "+valType)
    }
    return dummy.NewInstance(jsonData), nil
}


