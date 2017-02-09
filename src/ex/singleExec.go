package ex
import(
    "conf"
    )
type SingleExec struct{
    CnfMgr conf.ConfManager
}

func (e *SingleExec) StartExec() <-chan Event  {
    out:= make(chan Event)      
    go fireAll(e.CnfMgr.GetConfs(), out)
    return out
}

func (e *SingleExec) StopExec(){
    //nothing to do
}

