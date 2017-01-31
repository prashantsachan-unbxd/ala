package conf

import(
    "encoding/json"
    "io/ioutil" 
    "fmt"
    )
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
