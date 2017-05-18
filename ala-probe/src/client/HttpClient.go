package client

import (
    "errors"
    "strings"
    "net/http"
    log "github.com/Sirupsen/logrus"
    resp "response"
    topo "topology"
    "net/http/httputil"
    "strconv"
    )
const HTTP_CONFIG_PATH = "path"
const HTTP_CONFIG_METHOD = "method"
const HTTP_CONFIG_DATA = "data"

var Empty = &HttpClient{}

//HttpClient is an HTTP protocol implementation to probe a service
//it connects to the service over HTTP makes a requsts & wraps the response as HttpResponse
type HttpClient struct{
    client http.Client
    req http.Request
}
func (this *HttpClient) isEmpty()bool{
    return this == Empty
}
//New constructs a new HttpClient object with details required to make an HTTP call
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
//Execute method makes an HTTP call to the service according to config supplied at initialization
//it wraps & returns the response of the reqest into HttpResponse
func (this *HttpClient) Execute()(resp.ProbeResponse, error){
    if (this == Empty){
        return nil, errors.New("empty http client found")
    }
    rDump, _ := httputil.DumpRequest(&this.req, true)
    log.WithFields(log.Fields{"module": "httpClient","request":rDump}).Info("sending Http Request")
    res, err := this.client.Do(&this.req)
    if err !=nil{
        return nil,err
    }else{
        return resp.NewHttpResponse(*res), nil
    }
}