package main

import(
    "fmt"
    "time"
    // mux "github.com/gorilla/mux"
    // "net/http"
    topo "topology"
    "execute"
    "encoding/json"

)

func main(){
    var serviceDao topo.ServiceDao = &topo.FileServiceDao{"./resource/serviceConf.json"}
    serviceDao.Init()
    var REDao execute.RuleEngineDao = execute.RuleEngineDao{"http://ec2-54-173-96-124.compute-1.amazonaws.com", 8081, "",""}
    var exec execute.Executor
//    exec = & ex.SingleExec{cnfMgr}
    exec  = & execute.IntervalExec{Interval:7* time.Second, ServiceStore:serviceDao, REDao: REDao}
    fmt.Println("starting executor")
    out:= exec.StartExec()
    for{
        v:= <- out
        eStr,jErr:= json.Marshal(v)
        if jErr !=nil{
            fmt.Println("error in marshalling evnt", jErr)
        }else{
            fmt.Println("event Received:", string(eStr))
        }
    }
    
    
    // exec.StopExec()
    
}


