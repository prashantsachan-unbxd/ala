package nw

import (
    "net/http"
    "fmt"
    "api"
)

type ApiExec interface{
    Execute(map[api.Api]api.RespChecker)
}

func getSimpleClient()  http.Client{
    DefaultClient := http.Client{}
    return  DefaultClient
}
func GetStatus(req *http.Request) api.ApiStatus{
    client := getSimpleClient()
    res, err := client.Do(req)
    if err !=nil {
        fmt.Println("error executing api:", req)
        return api.STATUS_RED
    }
    respC := api.HttpCodeChecker{}
    apiStat :=  respC.GetStatus(*res, err)
    return apiStat
}
