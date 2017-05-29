package execute
import(
    "time"
    log "github.com/Sirupsen/logrus"
    topo "topology"
    "result"
    "execute/probe"
    )
//IntervalExec runs at a specific interval & computes metric Computation for all the services available
type IntervalExec struct{
    //Interval after which next batch is to be scheduled
    Interval time.Duration
    //ServiceStore (SerivceDao) to read configuration for all the services
    ServiceStore topo.ServiceDao
    //REDao : RuleEngineDao to get ProbeConfiguration & compute metric
    REDao RuleEngineDao
    //PCDao : to retrieve probeconfigs for a serviceClass
    PCDao probe.ProbeConfigDao
    //channel to stop execution
    done chan struct{}
}

func (e *IntervalExec) StartExec()<-chan result.Event   {
    out:= make(chan result.Event)
    e.done = make(chan struct{})
    go func(){
        terminated := false
        //Start a new Ticker to schedule batches at regular intervals
        ticker := time.NewTicker(e.Interval )
        for{
        select{
            case t:= <-ticker.C : 
                if terminated{
                    close(out)
                    ticker.Stop()
                }else{
                    // time to run metric computation for all services
                    services, err := e.ServiceStore.GetAllServices()
                    log.WithFields(log.Fields{"module":"executor","timestamp":t,
                        "numServices":len(services)}).Info("launching executionbatch")
                    if err != nil{
                        log.WithFields(log.Fields{"module":"executor","error": err}).Fatal(
                            "failed to launch batch, error in getting services")
                    }
                    go fetchMetrics(e.REDao, e.PCDao, services, out)

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