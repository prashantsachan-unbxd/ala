package main

import(
    "fmt"
    "time"
    "ex"    
    "result"
    "conf"
    "api"
    "ui"
    mux "github.com/gorilla/mux"
    "net/http"
)

func main(){
    var confStore conf.ConfStore
    confStore = &conf.FileConfStore{"./resource/apiConfig.json"}
    apiConfigs,err :=confStore.ReadApiConf()
    if err!=nil{
        fmt.Println(err)
        fmt.Println("exiting")
        return
    }   
    fmt.Println("apiConfigs: \n", apiConfigs)
    var exec ex.ApiExec
//    exec = & ex.SingleExec{apiConfigs}
    exec  = & ex.IntervalExec{Interval:5* time.Second, ApiData:apiConfigs}
 
    out:= exec.StartExec()
    
    sm:= result.TimedStateManager{6* time.Second,make(map[api.Api]time.Time)} 
    dispatcher := result.SimpleDispatcher{Consumers:[]result.EventConsumer{&result.EventLogger{}, & result.StateCollector{&sm}}}
    dispatcher.StartDispatch(out)
    handlers:=[]ui.ReqController{
//        &ui.JsonStateHandler{&sm},
//        &ui.HtmlStateHandler{&sm},
        &ui.StateController{&sm},
    }
    r:= mux.NewRouter()
    for _,h:= range handlers{
        h.Register(r)
    }
    fmt.Println("listening on port 8080")
    srv := &http.Server{
        Handler:      r,
        Addr:         ":8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    fmt.Println(srv.ListenAndServe()) 
    exec.StopExec()
    dispatcher.StopDispatch()
}


