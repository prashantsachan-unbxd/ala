package nw
import(
    "time"
    "sync"
    "api"
    "fmt"
    )
type IntervalExec struct{
    Interval time.Duration
}

func (e IntervalExec) Execute(apiData map[api.Api]api.RespCheck, done <-chan struct{}) <-chan Event   {
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
func fireAll(apiData map[api.Api]api.RespCheck, out chan<- Event){
    wg :=sync.WaitGroup{}
    for a,check := range apiData{
        a1 :=a
        check1 :=check
        wg.Add(1)
        go func(){
            timeStamp:= time.Now()
            status:=GetStatus(a1,check1)
            fmt.Println(a1.Method, a1.Url, " : ", status) 
            out<- Event{a1, timeStamp, status}
            wg.Done()
        }()
    }   
    wg.Wait()
}
