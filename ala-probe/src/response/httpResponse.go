package response

import(
    log "github.com/Sirupsen/logrus"
    "net/http"
    "io/ioutil"
    "encoding/json"
    )
const HTTP_FIELD_STATUS = "status"
const HTTP_FIELD_HEADERS = "headers"
const HTTP_FIELD_BODY = "body"
const HTTP_FIELD_VERSION = "version"
//HttpResponse is model class for HTTP response
type HttpResponse struct{
    m map[string]interface{}
}
func NewHttpResponse(response http.Response) *HttpResponse{
    defer response.Body.Close()
    respBody,err := ioutil.ReadAll(response.Body)
    if err !=nil{
        log.WithFields(log.Fields{"module": "httpResponse","error":err}).Info("error reading response body")
        respBody = nil    
    }
    m:= make(map[string]interface{})
    m[HTTP_FIELD_STATUS] = response.StatusCode
    m[HTTP_FIELD_HEADERS] = response.Header
    m[HTTP_FIELD_VERSION] = response.Proto
    //Try to convert this to map assuming it to be json response
    var respData  map[string]interface{}
    jsonErr:=json.Unmarshal(respBody, &respData)
    if jsonErr !=nil{
        log.WithFields(log.Fields{"module": "httpResponse","error":jsonErr, "value":string(respBody)}).Debug(
            "unable to parse resp body as JSON, passing it as string")
        m[HTTP_FIELD_BODY] = string(respBody)
    }else{
        m[HTTP_FIELD_BODY] = respData
    }
    return &HttpResponse{m}
}
func (this *HttpResponse) GetType()string{
    return "HTTP"
}

//AsMap returns HTTP response as a map
// response should contain keys: 'status','headers','body' & 'version'
func (this *HttpResponse) AsMap()map[string]interface{}{
    return this.m
}