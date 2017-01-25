package ex
import(
    "time"
    "sync"
//    "api"
    "fmt"
    "conf"
    )
type IntervalExec struct{
    Interval time.Duration
}

func (e *IntervalExec) Execute(apiData []conf.ApiConf, done <-chan struct{}) <-chan Event   {
    out:= make(chan Event)
    go func(){
        terminated := false
        ticker := time.NewTicker(e.Interval )
        for{
        select{
            case t:= <-ticker.C : 
                if terminated{
                    close(out)
                    ticker.Stop()
                }else{
                    fmt.Println("launching batch at: ", t)
                    go fireAll(apiData, out)
                }
            case <-done :
                // wait till next cycle to close the channel
                // allow current batch to finish
                terminated = true
                fmt.Println("stopping interval executor, won't schedule any more batches")
        }
        }
    }()
    return out
}
func fireAll(apiData []conf.ApiConf, out chan<- Event){
    wg :=sync.WaitGroup{}
    for _,c := range apiData{
        conf :=c
        wg.Add(1)
        go func(){
            timeStamp:= time.Now()
            status:=GetStatus(conf.Api,conf.Validator)
          //  fmt.Println(conf.Api.Method, conf.Api.Url, " : ", status) 
            out<- Event{conf.Api, timeStamp, status}
            wg.Done()
        }()
    }   
    wg.Wait()
}
