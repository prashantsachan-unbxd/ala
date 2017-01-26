package ex
import(
    "time"
    "sync"
    "fmt"
    "conf"
    )
type IntervalExec struct{
    Interval time.Duration
    ApiData []conf.ApiConf
    done chan struct{}
}

func (e *IntervalExec) StartExec()<-chan Event   {
    out:= make(chan Event)
    e.done = make(chan struct{})
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
                    go fireAll(e.ApiData, out)
                }
            case <-e.done :
                // wait till next cycle to close the channel
                // allow current batch to finish
                terminated = true
                fmt.Println("stopping interval executor, won't schedule any more batches")
        }
        }
    }()
    return out
}
func (e *IntervalExec)StopExec(){
    close(e.done)
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
