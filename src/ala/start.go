package main

import(
    "nw"    
    "api"
)

func main(){
    
    apiMap:= make( map[api.Api]api.RespChecker)
    
    apiMap[api.Api{"GET", "http://localhost:8000", nil}] = api.HttpCodeChecker{}
    apiMap[api.Api{"GET", "http://www.google.com", nil}] = api.HttpCodeChecker{}
    var exec nw.ApiExec
    exec = nw.SingleExec{}
    exec.Execute(apiMap)
}



