package main

import(
    "fmt"
    "time"
    "ex"    
    "result"
    "conf"
    "ui"
    mux "github.com/gorilla/mux"
    "net/http"
)

func main(){
    var confStore conf.ConfStore = &conf.FileConfStore{"./resource/apiConfig.json"}
    cnfMgr:= conf.ConfManager{[]conf.ConfLoader{confStore},confStore }  
    errs:= cnfMgr.Refresh(false)
    if errs!=nil && len(errs)>0{
        fmt.Println("error in reading configs")
        fmt.Println(errs)
        fmt.Println("exiting")
        return
    }   
    var exec ex.ApiExec
//    exec = & ex.SingleExec{cnfMgr}
    exec  = & ex.IntervalExec{Interval:10* time.Second, CnfMgr:cnfMgr}
 
    out:= exec.StartExec()
    
    sm:= result.TimedStateManager{6* time.Second, 6* time.Second} 
    dispatcher := result.SimpleDispatcher{Consumers:[]result.EventConsumer{&result.EventLogger{}, & result.StateCollector{&sm}, & result.KafkaForwarder{}}}
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


