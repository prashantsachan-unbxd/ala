package conf

import(
    "api"
    structs "github.com/fatih/structs"
    "encoding/json"
    )

func fromBasic(src basicConf)(ApiConf, error){
   valid, err:= api.GetValidator(src.ValidatorType, src.Validator)
    if err !=nil{
        return ApiConf{},err
    }
    return ApiConf{src.Api, src.ValidatorType,valid, src.Tags}, nil 
}
func toBasic(src ApiConf)basicConf{
    validType:= src.Validator.Type()
    validData := structs.Map(src.Validator)
        return basicConf{src.Api,   validType, validData, src.Tags}
}
func FromJson(data string) (ApiConf, error){
    var basic basicConf
    err  := json.Unmarshal([]byte(data), &basic)
    if err !=nil{
        return ApiConf{}, err
    }else{
        cnf, err := fromBasic(basic)
        if err !=nil{
            return ApiConf{},err
        }
        return cnf, nil
    }
}
