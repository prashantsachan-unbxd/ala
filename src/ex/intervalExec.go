package ex
import(
    "time"
    "fmt"
    "conf"
    )
type IntervalExec struct{
    Interval time.Duration
    CnfMgr conf.ConfManager
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
                    go fireAll(e.CnfMgr.GetConfs(), out)
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

