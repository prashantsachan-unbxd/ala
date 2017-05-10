package main

import(
    log "github.com/Sirupsen/logrus"
    "time"
    "os"
    topo "topology"
    "execute"
    "result"

)
func init(){
    log.SetFormatter(&log.JSONFormatter{})
    log.SetOutput(os.Stdout)
    log.SetLevel(log.InfoLevel)
}
func main(){
    var serviceDao topo.ServiceDao = &topo.FileServiceDao{"./resource/serviceConf.json"}
    serviceDao.Init()
    var REDao execute.RuleEngineDao = execute.RuleEngineDao{"http://ec2-54-173-96-124.compute-1.amazonaws.com", 8081, "",""}
    var exec execute.Executor
    exec  = & execute.IntervalExec{Interval:7* time.Second, ServiceStore:serviceDao, REDao: REDao}
    log.WithFields(log.Fields{"module": "main",}).Info("starting executor")    
    
    out:= exec.StartExec()
    dispatcher := result.SimpleDispatcher{Consumers:[]result.EventConsumer{&result.EventLogger{}, & result.KafkaForwarder{[]string{"localhost:9092"}}}}
    // dispatcher := result.SimpleDispatcher{Consumers:[]result.EventConsumer{&result.EventLogger{}}}
    log.WithFields(log.Fields{"module": "main",}).Info("starting dispatcher")
    dispatcher.StartDispatch(out)
    //wait forever
    select{}
    
    // exec.StopExec()
    
}


