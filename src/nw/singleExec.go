package nw
import(
    "api"
    "fmt"
    )
type SingleExec struct{
}

func (exec SingleExec) Execute(apiData map[api.Api]api.RespCheck, c<- chan struct{}){
    for a,check := range apiData{
        status := GetStatus(a, check)
        fmt.Println(a.Method, a.Url, " : ", status )
    }
}
