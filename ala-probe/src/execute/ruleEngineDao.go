package execute

import(
    "encoding/json"
    "errors"
    "net/http"
    "time"
    "bytes"
    "io/ioutil"
    "fmt"
    "strconv"
    "net/http/httputil"
    "execute/response"
)

type RuleEngineDao struct{
    Host string
    Port int
    User string
    Pass string
}
const RULE_RESOLUTION_PATH = "/rule-engine/rule-results/"
const RULE_RESULTION_METHOD = "POST"

const SEGMENT_FIELD_DOMAIN = "domain"
const SEGMENT_FIELD_SUBDOMAIN = "subdomain"
const SEGMENT_FIELD_SERVICECLASS = "serviceClass"
const SEGMENT_FIELD_RESPONSE = "response"


const SEGMENT_VALUE_DOMAIN = "platform.monitoring"
const SEGMENT_VALUE_SUBDOMAIN = "probeConfig"

const RESPONSE_FIELD_VALUE = "value"


// sends HTTP requests to ruleEngine returns back the response in ruleId-> response format
func (e *RuleEngineDao) resolveRule( segment map[string]interface{}) (map[string] interface{}, error){
    client := http.Client{Timeout: 5000* time.Millisecond}
    url:= e.Host +":"+ strconv.Itoa(e.Port)+RULE_RESOLUTION_PATH
    // fmt.Println("firign rule Resolve requst to url: ", url)
    data,jsonErr :=  json.Marshal(segment);
    if jsonErr!=nil{
        return nil, jsonErr
    }
    req,reqErr:= http.NewRequest(RULE_RESULTION_METHOD, url, bytes.NewReader(data))
    if reqErr !=nil{
        return nil, reqErr
    }
    req.Header.Add("Content-Type", "application/json")
    if(e.User!= "" && e.Pass !=""){
        req.SetBasicAuth(e.User, e.Pass)
    }
    rDump, _ := httputil.DumpRequest(req, true)
    fmt.Println(string(rDump))
    res, httpErr := client.Do(req)
    if httpErr !=nil {
        return nil, httpErr
    }
    if (res.StatusCode<200 || res.StatusCode >=300){
        return nil, errors.New("RuleEngine response Code: "+strconv.Itoa(res.StatusCode))
    }
    var respData  map[string]interface{}
    defer res.Body.Close()
    body, readErr:= ioutil.ReadAll(res.Body)
    if readErr !=nil{
        return nil, readErr
    }
    umErr:= json.Unmarshal(body, &respData)
    if umErr!= nil{
        return nil, umErr
    }
    return respData, nil
}

// Resolves rules for given segment & Returns RuleId-> value mapping
func (e *RuleEngineDao) resolveToVal(segment map[string]interface{}) (map[string]interface{},error){
    rules, reError:= e.resolveRule(segment)
    if reError!=nil{
        return nil, reError
    }
    result := make(map[string]interface{})
    for k,v := range rules{
        val := v.(map[string]interface{})
        result[k] = val[RESPONSE_FIELD_VALUE]
    }
    return result,nil
}

// returns List of ProbeConfigs for given serviceClass 
func (e *RuleEngineDao) GetProbeConfigs(serviceClass string) ([]ProbeConfig, error){
    // fmt.Println("fetching probeConfig for serice class:",serviceClass)
    segment := map[string]interface{}{
        SEGMENT_FIELD_DOMAIN: SEGMENT_VALUE_DOMAIN,
        SEGMENT_FIELD_SUBDOMAIN: SEGMENT_VALUE_SUBDOMAIN,
        SEGMENT_FIELD_SERVICECLASS: serviceClass,
    }
    ruleVals,reErr := e.resolveToVal(segment)
    if reErr !=nil{
        return nil, reErr
    }
    var probeConfs []ProbeConfig
    for _,v := range ruleVals{
        var c ProbeConfig
        jsonErr:= json.Unmarshal([]byte(v.(string)),&c)
        if jsonErr !=nil{
            fmt.Println("unable to create ProbeConfig from: "+v.(string))
            fmt.Println(jsonErr)
        }else{
            probeConfs = append(probeConfs, c);
        }
    }
    return probeConfs, nil
}

// return the first rule matching with the metricName
func (e *RuleEngineDao) GetMetricVal(resp response.ProbeResponse, segment map[string]interface{}, defaultVal interface{})(interface{}, error){
    segment[SEGMENT_FIELD_RESPONSE] = resp.AsMap()
    values,err := e.resolveToVal(segment)
    //ideally there should be exactly one rule with a specific metricName
    if err != nil{
        return defaultVal, err
    }else{
        // return any value (expecting a single value)
        for _,v:= range values{
            return v, nil
        }
        sStr,_:= json.Marshal(segment)
        return nil, errors.New("Empty Rule result with metricName: "+ string(sStr))
    }

}
