package conf

import(
    "fmt"
    "api"
    "io/ioutil"
    "encoding/json"
    )

type basicConf struct{
    Api api.Api `json:"api"`
    ValidatorType string `json:"validator_type"`
}
type ApiConf struct{
    Api api.Api 
    Validator api.ApiValidator
}

func fromBasic(src basicConf)ApiConf{
   valid:= GetValidator(src.ValidatorType)
   // fmt.Println("got validator: ", valid)
   return ApiConf{src.Api, valid}    
}
func toBasic(src ApiConf)basicConf{
    validType:= GetValidatorType(src.Validator)
    return basicConf{src.Api, validType}
}
func readBasicConf(path string)[] basicConf{
    file, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println("Config File Missing. ", err)
    }

    var config [] basicConf
    err = json.Unmarshal(file, &config)
    if err != nil {
        fmt.Println("Config Parse Error: ", err)
    }
    return config  
}
func ReadApiConf(path string) []ApiConf{
    basics := readBasicConf(path)
    //fmt.Println("basics:", basics)
    var configs []ApiConf
    for _,b:= range basics{
        c := fromBasic(b)
        configs = append(configs, c)
    }
    return configs
}

//func writeBasicConf(configs []basicConf, filePath string)
//func WriteApiConf(configs []ApiConf, filePath string)
