package nw

import (
    "net/http"
    "fmt"
    "api"
)

type ApiExec interface{
    Execute(map[api.Api]api.RespCheck)
}

func getSimpleClient()  http.Client{
    DefaultClient := http.Client{}
    return  DefaultClient
}
func GetStatus(a api.Api, respCheck api.RespCheck) api.ApiStatus{
    client := getSimpleClient()
    req,err:= http.NewRequest(a.Method, a.Url, a.Data)
    if err !=nil{
        fmt.Println("unable to create httpReq for:", a.Method, a.Url, a.Data)
        return api.STATUS_YELLOW
    }
    res, err := client.Do(req)
    if err !=nil {
        fmt.Println("error executing api:", req, "\n", err)
        return api.STATUS_RED
    }
    apiStat :=  respCheck.GetStatus(*res, err)
    return apiStat
}
