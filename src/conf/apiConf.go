package conf

import(
    "github.com/satori/go.uuid"
    "api"
    )

type basicConf struct{
    Id uuid.UUID `json:"id"`
    Api api.Api `json:"api"`
    ValidatorType string `json:"validator_type"`
    Validator map[string]interface{} `json:"validator"`
    Tags []string `json:"tags"`
}
type ApiConf struct{
    Id uuid.UUID `json:"id"`
    Api api.Api `json:"api"`
    ValidatorType string `json:"validator_type"`
    Validator api.ApiValidator `json:"validator"`
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
