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

func (e IntervalExec) Execute(apiData map[api.Api]api.RespCheck, done <-chan struct{}) {
    go func(){
        ticker := time.NewTicker(e.Interval )
        for{
        select{
            case t:= <-ticker.C : 
                fmt.Println("launching batch at: ", t)
                go fireAll(apiData)
            case <-done :
                ticker.Stop()
                fmt.Println("stopping interval executor")
        }
        }
    }()
}
func fireAll(apiData map[api.Api]api.RespCheck){
    fmt.Println("api's to fire: ", apiData)
    wg :=sync.WaitGroup{}
    for a,check := range apiData{
        a1 :=a
        check1 :=check
        wg.Add(1)
        go func(){
            status:=GetStatus(a1,check1)
            fmt.Println(a1.Method, a1.Url, " : ", status) 
            wg.Done()
        }()
    }   
    wg.Wait()
}
