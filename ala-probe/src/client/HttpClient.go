package client

import (
    "errors"
    "strings"
    "net/http"
    resp "response"
    topo "topology"
    "net/http/httputil"
    "fmt"
    "strconv"
    )
const HTTP_CONFIG_PATH = "path"
const HTTP_CONFIG_METHOD = "method"
const HTTP_CONFIG_DATA = "data"

var Empty = &HttpClient{}
type HttpClient struct{
    client http.Client
    req http.Request
}
func (this *HttpClient) isEmpty()bool{
    return this == Empty
}
func (this *HttpClient) New(config map[string]interface{}, service topo.Service) (ProbeClient, error){
    client := http.Client{}
    url := service.Host + ":"+strconv.Itoa(service.Port) +config[HTTP_CONFIG_PATH].(string)
    method := "GET"
    m,ok:= config[HTTP_CONFIG_METHOD]
    if(ok){
        method = m.(string)
    }
    data:=""
    d,ok := config[HTTP_CONFIG_DATA]
    if(ok){
        data = d.(string)
    }
    req, err := http.NewRequest(method, url, strings.NewReader(data))
    if(err !=nil){
        return Empty, err
    }else{
        return &HttpClient{client, *req}, nil
    }
}
    // Executes the probe request as per config & return a result
func (this *HttpClient) Execute()(resp.ProbeResponse, error){
    if (this == Empty){
        return nil, errors.New("empty http client found")
    }
    rDump, _ := httputil.DumpRequest(&this.req, true)
    fmt.Println(string(rDump))
    res, err := this.client.Do(&this.req)
    if err !=nil{
        return nil,err
    }else{
        return & resp.HttpResponse{*res}, nil
    }
}