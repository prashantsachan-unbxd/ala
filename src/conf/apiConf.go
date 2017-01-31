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

