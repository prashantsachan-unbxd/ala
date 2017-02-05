package conf

import(
    "encoding/json"
    "io/ioutil" 
    "fmt"
    )
type FileConfStore struct{
    FilePath string
}
func readFromFile(path string)([] basicConf,error){
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
func (d *FileConfStore)ReadApiConf() ([]ApiConf,error){
    basics,err := readFromFile(d.FilePath)
    if err !=nil{
        return nil, err
    }
    var configs []ApiConf
    for _,b:= range basics{
        c,err := fromBasic(b)
        if err!=nil{
            fmt.Println("unable to load config: "+fmt.Sprintf("%#v", b)+"\n"+err.Error())
        }else{
            configs = append(configs, c)
        }
    }
    return configs, nil
}
func writeToFile(basics []basicConf, filePath string) error{
    b, err := json.Marshal(basics)
    if err != nil { return err}

    ioutil.WriteFile(filePath, b, 0644)
    return nil
}
func (d *FileConfStore) WriteApiConf(configs []ApiConf) error{
    var basics []basicConf
    for _,c:= range configs{
        b := toBasic(c)
        basics = append(basics, b)
    }
    return writeToFile(basics, d.FilePath)
}
