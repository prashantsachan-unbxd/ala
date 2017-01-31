package conf

import(
    "api"
    )

type basicConf struct{
    Api api.Api `json:"api"`
    ValidatorType string `json:"validator_type"`
}
type ApiConf struct{
    Api api.Api 
    Validator api.ApiValidator
}
type ConfLoader interface{
    Read() ([]ApiConf,error)
}
//    func Write([] ApiConf) error
//func writeBasicConf(configs []basicConf, filePath string)
//func WriteApiConf(configs []ApiConf, filePath string)
