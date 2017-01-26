package ex
import(
    "conf"
    )
type SingleExec struct{
    ApiData [] conf.ApiConf
}

func (e *SingleExec) StartExec() <-chan Event  {
    out:= make(chan Event)      
    go fireAll(e.ApiData, out)
    return out
}

func (e *SingleExec) StopExec(){
    //nothing to do
}

