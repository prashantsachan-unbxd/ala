package nw
import(
    "net/http"
    "api"
    "fmt"
    )
type SingleExec struct{
}

func (exec SingleExec) Execute(apiData map[api.Api]api.RespChecker){
    for a,_ := range apiData{
        req,err := http.NewRequest(a.Method, a.Url, a.Data)
        var status api.ApiStatus
        if err !=nil{
            status= api.STATUS_RED
        }else{
            status = GetStatus(req)
        }
        fmt.Println(a.Method, a.Url, " : ", status )
    }
}
