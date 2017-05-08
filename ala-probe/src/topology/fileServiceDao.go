package topology

import(
    "fmt"
    "encoding/json"
    "io/ioutil"
    "errors"
    )

var services []Service
type FileServiceDao struct{
    FilePath string
}
func (d *FileServiceDao) Init()(){
    configs, err := loadFromFile(d.FilePath)
    if err!=nil{
        fmt.Println("unable to read service configs from file", d.FilePath,"\n",err)
    }else{
        services = configs
        fmt.Println("loaded service configs from file", d.FilePath)
    }
}

func (d *FileServiceDao) GetAllServices()([]Service,error){
    return services, nil
}
func(d *FileServiceDao) AddService(s Service)error{
    services = append(services,s)
    err:= writeToFile(services, d.FilePath)
    if err !=nil{
        services = services[:len(services)-1]
    }
    return err

}
func(d *FileServiceDao) DeleteService(id string) error{
    idx := -1
    for i,s := range services{
        if s.Id ==  id{
            idx = i
            break
        }
    }
    if idx <0{
        return errors.New("No service config with id: "+ id)
    }else{
        removed := append(services[:idx], services[idx+1:]...)
        err := writeToFile(removed, d.FilePath)
        if err ==nil{
            services = removed
        }
        return err
    }
    
}


func loadFromFile(path string)([] Service,error){
    file, err := ioutil.ReadFile(path)
    if err != nil {
        return nil,err
    }

    var configs [] Service
    err = json.Unmarshal(file, &configs)
    if err != nil {
        return nil, err
    }
    return configs,nil
}

func writeToFile(configs []Service, filePath string) error{
    b, err := json.Marshal(configs)
    if err != nil { return err}
    ioutil.WriteFile(filePath, b, 0644)
    return nil
}   
