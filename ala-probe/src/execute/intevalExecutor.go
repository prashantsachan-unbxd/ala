package execute
import(
    "time"
    log "github.com/Sirupsen/logrus"
    topo "topology"
    "result"
    )
type IntervalExec struct{
    Interval time.Duration
    ServiceStore topo.ServiceDao
    REDao RuleEngineDao
    done chan struct{}
}

func (e *IntervalExec) StartExec()<-chan result.Event   {
    out:= make(chan result.Event)
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
                    
                    services, err := e.ServiceStore.GetAllServices()
                    log.WithFields(log.Fields{"module":"executor","timestamp":t,
                        "numServices":len(services)}).Info("launching executionbatch")
                    if err == nil{
                        go fetchMetrics(e.REDao, services, out)
                    }else{
                        log.WithFields(log.Fields{"module":"executor","error": err}).Fatal(
                            "failed to launch batch, error in getting services")
                    }

                }
            case <-e.done :
                // wait till next cycle to close the channel
                // allow current batch to finish
                terminated = true
                log.WithFields(log.Fields{"module":"executor"}).Info(
                    "stopping interval executor, won't schedule any more batches")
        }
        }
    }()
    return out
}
func (e *IntervalExec)StopExec(){
    close(e.done)
}