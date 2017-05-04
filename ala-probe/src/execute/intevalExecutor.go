package execute
import(
    "time"
    "fmt"
    topo "topology"
    )
type IntervalExec struct{
    Interval time.Duration
    ServiceStore topo.ServiceDao
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
                    services, err := e.ServiceStore.GetAllServices()
                    if err == nil{
                        go fireAll(services, out)
                    }else{
                        fmt.Println("failed to launch batch, error in getting services", err)
                    }

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