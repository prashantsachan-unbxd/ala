package conf

import(
    "encoding/json"
    "io/ioutil" 
    )
type FileConfDao struct{
    FilePath string
}
func readBasicConf(path string)([] basicConf,error){
    file, err := ioutil.ReadFile(path)
    if err != nil {
        return nil,err
    }

    var config [] basicConf
    err = json.Unmarshal(file, &config)
    if err != nil {
        return nil, err
    }
    return config,nil
}
func (d *FileConfDao)Read() ([]ApiConf,error){
    basics,err := readBasicConf(d.FilePath)
    if err !=nil{
        return nil, err
    }
    var configs []ApiConf
    for _,b:= range basics{
        c := fromBasic(b)
        configs = append(configs, c)
    }
    return configs, nil
}
//func writeBasicConf(configs []basicConf, filePath string)
//func WriteApiConf(configs []ApiConf, filePath string)
