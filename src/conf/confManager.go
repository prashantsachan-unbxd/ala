package conf
import(
    "github.com/satori/go.uuid"
    "errors"
    )
var configs map[uuid.UUID]ApiConf
type ConfManager struct{
    Loaders []ConfLoader
    Writer ConfWriter
}
func (m ConfManager)Refresh(abortOnError bool)[]error{
    var errs []error
    tConfigs:= make(map[uuid.UUID]ApiConf)
    for _,l:= range m.Loaders{
        newConfs,err := l.ReadApiConf()
        if err!=nil{
            errs = append(errs, err)
        }else{
            for _,c := range newConfs{
                tConfigs[c.Id] = c
            }
        }
    }
    if errs== nil || len(errs)==0 && !abortOnError{
        configs = tConfigs
    }
    return errs
}
func toList(m map[uuid.UUID] ApiConf)[]ApiConf{
    var confs []ApiConf
    for _,c := range m{
        confs = append(confs, c)
    }
    return confs
}
func (m ConfManager)GetConfs()[]ApiConf{
    return toList(configs)
}
func (m ConfManager)FilterConfs(tags []string)[]ApiConf{
    filtered:=make( map[uuid.UUID]ApiConf)
    for id,c:= range configs{
        found:= false
        for _,t:= range c.Tags{
            if found{
                break
            } 
            for _,t1:= range tags{
                if(t == t1){
                    filtered[id] = c
                    found = true
                    break;
                }
            }
        }
    }
    return toList(filtered)
}
func (m ConfManager)AddConf(newConf ApiConf) (uuid.UUID,error){
    if newConf.Id == (uuid.UUID{}) {
        newConf.Id = uuid.NewV4()
    }
    _,ok:= configs[newConf.Id]
    if ok{
        return uuid.UUID{},errors.New("there already exists another apiConf with id:"+ newConf.Id.String())
    }else{
        configs[newConf.Id] = newConf
        err:= m.Writer.WriteApiConf(toList(configs))
        if err !=nil{
            delete(configs, newConf.Id)
            return uuid.UUID{},err 
        }else{
            return newConf.Id,nil
        }
    }

}
func (m ConfManager)DelteConf(id uuid.UUID )error{
    _,ok:= configs[id]
    if ok{
        delete(configs, id)
        return m.Writer.WriteApiConf(toList(configs))
    }else{
        return errors.New("No apiConf with id: "+id.String())
    }
}
