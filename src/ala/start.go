package main

import(
    "fmt"
    "time"
    "ex"    
    "api"
    "result"
)

func main(){
    
    apiMap:= make( map[api.Api]api.ApiValidator)
    
    apiMap[api.Api{"GET", "http://localhost:8000", nil}] = api.HttpCodeChecker{}
    apiMap[api.Api{"GET", "http://www.google.co.in", nil}] = api.HttpCodeChecker{}
    fmt.Println("apiMap: ", apiMap)
    var exec ex.ApiExec
    //exec = ex.SingleExec{}
    exec  = ex.IntervalExec{5* time.Second}
    done := make(chan struct{})
    out:= exec.Execute(apiMap, done)
    
    fmt.Println("sleeping for 30 sec")
    
    dispatcher := result.SimpleDispatcher{[]result.EventConsumer{result.EventLogger{}}, new (chan struct{})}
    dispatcher.StartDispatch(out)
    time.Sleep(30* time.Second)
    fmt.Println("jaag utha shaitan")
    done <- struct{}{}
    //for e:= range out{
    //    dispatcher.Dispatch(e)
    //}
    close(done)
    dispatcher.StopDispatch()
}



