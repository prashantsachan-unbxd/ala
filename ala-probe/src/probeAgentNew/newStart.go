package main

import(
    "fmt"
    "time"
    // mux "github.com/gorilla/mux"
    // "net/http"
    topo "topology"
    "execute"

)

func main(){
    var serviceDao topo.ServiceDao = &topo.FileServiceDao{"./resource/serviceConf.json"}
    serviceDao.Init()
    var REDao execute.RuleEngineDao = execute.RuleEngineDao{"http://ec2-54-173-96-124.compute-1.amazonaws.com", 8081, "",""}
    var exec execute.Executor
//    exec = & ex.SingleExec{cnfMgr}
    exec  = & execute.IntervalExec{Interval:5* time.Second, ServiceStore:serviceDao, REDao: REDao}
    fmt.Println("starting executor")
    out:= exec.StartExec()
    v:= <-out
    fmt.Println("received event from executor", v)
    
    
    // exec.StopExec()
    
}


