package main

import(
    "fmt"
    "time"
    "ex"    
    "result"
    "conf"
)

func main(){
    apiConfigs:= conf.ReadApiConf("./resource/apiConfig.json")
    fmt.Println("apiConfigs: \n", apiConfigs)
   
    var exec ex.ApiExec
//    exec = & ex.SingleExec{apiConfigs}
    exec  = & ex.IntervalExec{Interval:5* time.Second, ApiData:apiConfigs}
 
    out:= exec.StartExec()
    
    fmt.Println("sleeping for 30 sec")
    
    dispatcher := result.SimpleDispatcher{Consumers:[]result.EventConsumer{&result.EventLogger{}}}
    dispatcher.StartDispatch(out)
    time.Sleep(30* time.Second)
    fmt.Println("jaag utha shaitan")
    exec.StopExec()
    dispatcher.StopDispatch()
}



