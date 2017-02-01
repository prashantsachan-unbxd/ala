package conf

import(
    )

func fromBasic(src basicConf)ApiConf{
   valid:= GetValidator(src.ValidatorType)
   return ApiConf{src.Api, valid, src.Tags}    
}
func toBasic(src ApiConf)basicConf{
    validType:= GetValidatorType(src.Validator)
    return basicConf{src.Api, validType, src.Tags}
}
