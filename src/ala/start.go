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
    var confStore conf.ConfStore = &conf.FileConfStore{"./resource/apiConfig.json"}
    cnfMgr:= conf.ConfManager{[]conf.ConfLoader{confStore},confStore }  
    errs:= cnfMgr.Refresh(false)
//    apiConfigs,err :=confStore.ReadApiConf()
    if errs!=nil && len(errs)>0{
        fmt.Println("error in reading configs")
        fmt.Println(errs)
        fmt.Println("exiting")
        return
    }   
    var exec ex.ApiExec
//    exec = & ex.SingleExec{apiConfigs}
    exec  = & ex.IntervalExec{Interval:5* time.Second, ApiData:cnfMgr.GetConfs()}
 
    out:= exec.StartExec()
    
    sm:= result.TimedStateManager{6* time.Second,make(map[api.Api]time.Time)} 
    dispatcher := result.SimpleDispatcher{Consumers:[]result.EventConsumer{&result.EventLogger{}, & result.StateCollector{&sm}}}
    dispatcher.StartDispatch(out)
    controllers:=[]ui.ReqController{
        &ui.StateController{&sm},
        &ui.ConfController{ cnfMgr},
    }
    r:= mux.NewRouter()
    for _,h:= range controllers{
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


