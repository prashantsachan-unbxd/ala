package response

import(
    "encoding/json"
    "net/http"
    "io/ioutil"
    "fmt"
    )
const HTTP_FIELD_STATUS = "status"
const HTTP_FIELD_HEADERS = "headers"
const HTTP_FIELD_BODY = "body"
const HTTP_FIELD_VERSION = "version"

type HttpResponse struct{
    Resp http.Response
}
    
func (this *HttpResponse) getType()string{
    return "HTTP"
}

func (this *HttpResponse) getJson()string{
    m:= make(map[string]interface{})
    m[HTTP_FIELD_STATUS]= this.Resp.StatusCode
    m[HTTP_FIELD_HEADERS]= this.Resp.Header
    defer this.Resp.Body.Close()
    respBody,err := ioutil.ReadAll(this.Resp.Body)
    if err !=nil{
        respBody = nil    
    }
    m[HTTP_FIELD_BODY] = respBody
    m[HTTP_FIELD_VERSION] = this.Resp.Proto
    val, err := json.Marshal(m)
    if err !=nil{
        fmt.Println("HttpResponse: unable to convert to Json: ", m, "due to error:", err)
        return ""
    }else{
        return string(val)
    }
}