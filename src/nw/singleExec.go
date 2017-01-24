package nw
import(
    "api"
    "fmt"
    "time"
    )
type SingleExec struct{
}

func (exec SingleExec) Execute(apiData map[api.Api]api.RespCheck, c<- chan struct{}) <-chan Event  {
    out:= make(chan Event)      
    for a,check := range apiData{
        timeStamp:= time.Now()
        status := GetStatus(a, check)
        fmt.Println(a.Method, a.Url, " : ", status )
        out<- Event{a, timeStamp, status}
    }
    close(out)
    return out
}
