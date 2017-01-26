package ex
import(
    "conf"
    )
type SingleExec struct{
}

func (exec *SingleExec) Execute(apiData []conf.ApiConf, c<- chan struct{}) <-chan Event  {
    out:= make(chan Event)      
    go fireAll(apiData, out)
    return out
}

