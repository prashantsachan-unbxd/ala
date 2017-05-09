package main

import(
    "fmt"
    "time"
    topo "topology"
    "execute"
    "execute/result"

)

func main(){
    var serviceDao topo.ServiceDao = &topo.FileServiceDao{"./resource/serviceConf.json"}
    serviceDao.Init()
    var REDao execute.RuleEngineDao = execute.RuleEngineDao{"http://ec2-54-173-96-124.compute-1.amazonaws.com", 8081, "",""}
    var exec execute.Executor
    exec  = & execute.IntervalExec{Interval:7* time.Second, ServiceStore:serviceDao, REDao: REDao}
    fmt.Println("starting executor")
    
    
    out:= exec.StartExec()
    dispatcher := result.SimpleDispatcher{Consumers:[]result.EventConsumer{&result.EventLogger{}, & result.KafkaForwarder{[]string{"localhost:9092"}}}}
    // dispatcher := result.SimpleDispatcher{Consumers:[]result.EventConsumer{&result.EventLogger{}}}
    dispatcher.StartDispatch(out)
    //wait forever
    select{}
    
    // exec.StopExec()
    
}


