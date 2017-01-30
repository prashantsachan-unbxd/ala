package main

import(
    "fmt"
    "time"
    "ex"    
    "result"
    "conf"
    "api"
    "ui"
    "net/http"
)

func main(){
    apiConfigs:= conf.ReadApiConf("./resource/apiConfig.json")
    fmt.Println("apiConfigs: \n", apiConfigs)
    var exec ex.ApiExec
//    exec = & ex.SingleExec{apiConfigs}
    exec  = & ex.IntervalExec{Interval:5* time.Second, ApiData:apiConfigs}
 
    out:= exec.StartExec()
    
    sm:= result.TimedStateManager{6* time.Second,make(map[api.Api]time.Time)} 
    dispatcher := result.SimpleDispatcher{Consumers:[]result.EventConsumer{&result.EventLogger{}, & result.StateCollector{&sm}}}
    dispatcher.StartDispatch(out)
    handlerMap:=map[string]ui.ReqHandler{
        "/state": &ui.JsonStateHandler{&sm},
        "/state.html": &ui.HtmlStateHandler{&sm},
    }
    for path,h:= range handlerMap{
        fmt.Println("setting handler for", path)
        http.HandleFunc(path,h.HandleFunc())
    }
    fmt.Println("listening on port 8080")
    http.ListenAndServe(":8080", nil)
    
    exec.StopExec()
    dispatcher.StopDispatch()
}


