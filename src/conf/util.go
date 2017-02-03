package conf

import(
    "api"
    )

func fromBasic(src basicConf)ApiConf{
   valid:= api.GetValidator(src.ValidatorType)
   return ApiConf{src.Api, valid, src.Tags}    
}
func toBasic(src ApiConf)basicConf{
    validType:= api.GetValidatorType(src.Validator)
    return basicConf{src.Api, validType, src.Tags}
}
