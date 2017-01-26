package ex
import(
    "net/http"
    "api"
    "conf"
    "time"
    "fmt"
    "sync"
    "strings"
    )
func getSimpleClient()  http.Client{
    DefaultClient := http.Client{Timeout: 900* time.Millisecond}
    return  DefaultClient
}
func GetStatus(a api.Api, respCheck api.ApiValidator) api.ApiStatus{
    client := getSimpleClient()
    req,err:= http.NewRequest(a.Method, a.Url, strings.NewReader(a.Data))
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
func fireAll(apiData []conf.ApiConf, out chan<- Event){
    wg :=sync.WaitGroup{}
    for _,c := range apiData{
        conf :=c
        wg.Add(1)
        go func(){
            timeStamp:= time.Now()
            status:=GetStatus(conf.Api,conf.Validator)
          //  fmt.Println(conf.Api.Method, conf.Api.Url, " : ", status) 
            out<- Event{conf.Api, timeStamp, status}
            wg.Done()
        }()
    }   
    wg.Wait()
}
