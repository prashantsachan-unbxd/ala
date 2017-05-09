package response

import(
    "net/http"
    "io/ioutil"
    )
const HTTP_FIELD_STATUS = "status"
const HTTP_FIELD_HEADERS = "headers"
const HTTP_FIELD_BODY = "body"
const HTTP_FIELD_VERSION = "version"

type HttpResponse struct{
    Resp http.Response
}
    
type httpRespModel struct{
    status int 
    headers http.Header 
    body string 
    version string 
}
func (this *httpRespModel) asMap()map[string]interface{}{
    m:= make(map[string]interface{})
    m[HTTP_FIELD_STATUS] = this.status
    m[HTTP_FIELD_HEADERS] = this.headers
    m[HTTP_FIELD_VERSION] = this.version
    m[HTTP_FIELD_BODY] = this.body
    return m;
}
func (this *HttpResponse) GetType()string{
    return "HTTP"
}

func (this *HttpResponse) AsMap()map[string]interface{}{
    defer this.Resp.Body.Close()
    respBody,err := ioutil.ReadAll(this.Resp.Body)
    if err !=nil{
        respBody = nil    
    }
    modelResp:= httpRespModel{ this.Resp.StatusCode, this.Resp.Header,string(respBody), this.Resp.Proto}
    return modelResp.asMap()
}