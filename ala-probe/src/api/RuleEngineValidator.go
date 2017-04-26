package api
import(
    "fmt"
    "net/http"
    "strings"
    "encoding/json"
    "time"
    "io/ioutil"
)
var rePath = "/rule-engine/rule-results/"
var body_field_response  = "response"
var response_field_version = "version"
var response_field_status = "status"
var response_field_headers = "headers"
var response_field_body = "body"
var reMethod = "POST"

type RuleEngineValidator struct{
    reHost string
    reBody map[string]interface{}
}

func (c *RuleEngineValidator) GetStatus (resp http.Response, err error) ApiStatus{

    if err !=nil{
        fmt.Println(err)
        return STATUS_RED
    }else{
        // url
        url := c.reHost+rePath
        defer resp.Body.Close()
        respBody,err := ioutil.ReadAll(resp.Body)
        if err !=nil{
            fmt.Println("unable to  read api response body to string:", resp.Body)
            return STATUS_RED
        }
        // make rule resolution request body
        respMap:= make(map[string]interface{})
        respMap[response_field_version] = resp.Proto
        respMap[response_field_status] = resp.StatusCode
        respMap[response_field_headers] = resp.Header
        respMap[response_field_body] = string(respBody)

        c.reBody [body_field_response] = respMap
        body, err := json.Marshal(c.reBody)
        if err !=nil{
            fmt.Println("unable to convert to Json: ", c.reBody, "due to error:", err)
            return STATUS_RED
        }else{
            return c.ruleEngineSuccess(url, string(body))
        }
   
    }
}
func (c *RuleEngineValidator) NewInstance(jsonData map[string]interface{})ApiValidator{
    // get rule Engine endpoint
    // data object 
    // 
    return & RuleEngineValidator{jsonData["reHost"].(string), jsonData["reBody"].(map[string]interface{})}
}
func (c *RuleEngineValidator) Type() string{
    return VALIDATOR_TYPE_RE_HTTPCODE
}
func (c *RuleEngineValidator) String() string{
    return c.Type()
}

func (c *RuleEngineValidator) ruleEngineSuccess(url string , reqBody string) ApiStatus{
    client := http.Client{Timeout: 5000* time.Millisecond}
    req,err:= http.NewRequest(reMethod, url, strings.NewReader(reqBody))
    if err !=nil{
        fmt.Println("unable to create rule resolution httpReq for:", )
        return STATUS_RED
    }
    req.Header.Set("Content-Type", "application/json")
    res, err := client.Do(req)
    if err !=nil {
        fmt.Println("error executing api:", req, "\n", err)
        return STATUS_RED
    }
    if res.StatusCode >=300 || res.StatusCode<200{
        fmt.Println("rule engine response status: ", res.StatusCode, url)
        return STATUS_RED
    }
    defer res.Body.Close()
    var m map[string]interface{}
    resBody, err:= ioutil.ReadAll(res.Body)
    if err !=nil{
        fmt.Println("error converting RE response to string: ", res.Body)
        return STATUS_RED
    }
    err = json.Unmarshal(resBody, & m)
    if err !=nil{
        fmt.Println("error converting RE response : ", resBody , "to map, with reqBody:", reqBody)
        return STATUS_RED
    }

    // if any one of the rule results is not 'true', its red
    for k, v := range m {
        vmap := v.(map[string]interface{})
        if ! vmap["value"].(bool){
            fmt.Println("non-true value:", vmap["value"],"for ruleID:" ,k)
            return STATUS_RED
        }
            
    }
    return STATUS_GREEN
}