package conf

import(
    "api"
    structs "github.com/fatih/structs"
    )

func fromBasic(src basicConf)ApiConf{
   valid:= api.GetValidator(src.ValidatorType, src.Validator)
   return ApiConf{src.Api, src.ValidatorType,valid, src.Tags}    
}
func toBasic(src ApiConf)basicConf{
    validType:= src.Validator.Type()
    validData := structs.Map(src.Validator)
        return basicConf{src.Api,   validType, validData, src.Tags}
}
