package conf

import(
    "api"
    )

type basicConf struct{
    Api api.Api `json:"api"`
    ValidatorType string `json:"validator_type"`
    Tags []string `json:"tags"`
}
type ApiConf struct{
    Api api.Api 
    Validator api.ApiValidator
    Tags []string `json:"tags"`
}
type ConfLoader interface{
    ReadApiConf() ([]ApiConf,error)
}
type ConfWriter interface{
    WriteApiConf([] ApiConf) error
}
type ConfStore interface{
    ConfLoader
    ConfWriter
}
