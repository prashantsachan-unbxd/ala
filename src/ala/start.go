package main

import(
    "fmt"
    "time"
    "nw"    
    "api"
)

func main(){
    
    apiMap:= make( map[api.Api]api.RespCheck)
    
    apiMap[api.Api{"GET", "http://localhost:8000", nil}] = api.HttpCodeChecker{}
    apiMap[api.Api{"GET", "http://www.google.co.in", nil}] = api.HttpCodeChecker{}
    fmt.Println("apiMap: ", apiMap)
    var exec nw.ApiExec
    //exec = nw.SingleExec{}
    exec  = nw.IntervalExec{5* time.Second}
    done := make(chan struct{})
    out:= exec.Execute(apiMap, done)
    
    fmt.Println("sleeping for 30 sec")
    time.Sleep(30* time.Second)
    fmt.Println("jaag utha shaitan")
    done <- struct{}{}
    for e:= range out{
        fmt.Println("event:", e)
    }
    close(done)
}



