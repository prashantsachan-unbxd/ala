package main

import(
    "fmt"
    "net/http"
    "nw"    
    "api"
)

func main(){
    url  := "http://www.google.com"
    apiStat := GetStatus("GET",url) 
    fmt.Println("api status for url: ", url, " is: ", apiStat) 
}

func GetStatus(method, url string) api.ApiStatus{
    client := nw.GetSimpleClient()
    exec := nw.SimpleApiExec{client}
    req, err := http.NewRequest(method, url, nil)
    if err != nil{
       fmt.Println("can't crate api with method", method, "& url:", url)
       return api.YELLOW
    }
    res, err := exec.Fire(req)
    if err !=nil {
        fmt.Println("error executing api:", method, url)
        return api.RED
    }
    respC := api.HttpCodeChecker{}
    apiStat :=  respC.GetStatus(*res, err)
    return apiStat
}
