package execute

import (
        "fmt"
        topo "topology"
        "encoding/json"
        "execute/client"
        "execute/response"
        "time"
    )

const FIELD_METRIC_NAME = "metricName"
const FIELD_DEFAULT_VALUE = "defaultMetricValue"
const FIELD_METRICS = "metrics"
func fetchMetrics(reDao RuleEngineDao, services []topo.Service, out chan Event){
    // var results map[string]interface{}
    for _,s := range services{
        confs:= fetchProbeConfig(s,reDao);
        fmt.Println(" ProbeConfigs for service: "+s.Id)
        for _,c := range confs{
            str, _:= json.Marshal(c)
            fmt.Println(string(str))
            client, cErr:= client.GetClient(c.ProbeType, c.ProbeData, s)
            if(cErr!=nil){
                fmt.Println("service:", s.Id, "error instantiating client Type,data:", c.ProbeType, c.ProbeData)
                fmt.Println(cErr)
                // forward the default valued event 
            }
            pResp, pErr := client.Execute()
            if(pErr !=nil){
                fmt.Println("service:", s.Id, "clientType:",c.ProbeType, "error in probing")
                fmt.Println(pErr)
                //forward the default valued event
            }
            timestamp:= time.Now();
            metrics := getMetricValues(reDao, pResp, c.Metrics)
            for k,v := range metrics{
                fmt.Println("serviceId:", s.Id, k,"=",v)
                out <- Event{s,timestamp, k,v.(float64)}
                
            }
        }
    }
}

func getMetricValues(reDao RuleEngineDao, resp response.ProbeResponse, metrics[]map[string]interface{}) (map[string]interface{}){
    vals := make(map[string]interface{})
    if resp == nil {
        for _,m1 := range metrics{
            vals[m1[FIELD_METRIC_NAME].(string)] = m1[FIELD_DEFAULT_VALUE]
        }
    }else{
        for _,m := range metrics{
            var defaultVal interface{}
            segment := make(map[string]interface{})
            for k,v := range m{
                if k == FIELD_DEFAULT_VALUE{
                    defaultVal = v
                }else{
                    segment[k] = v
                }
            }
            mName:=  m[FIELD_METRIC_NAME].(string)
            val, reErr:= reDao.GetMetricVal(resp, segment, defaultVal)
            if reErr !=nil {
                fmt.Println("error retrieving metric Named:",mName , "from ProbeResponse" )
                fmt.Println("Error:",reErr)
            }
            vals[mName] = val
        }
    }
    return vals
}

func fetchProbeConfig(service topo.Service, reDao RuleEngineDao)[]ProbeConfig{
    fmt.Println("fetching probeConfigs for service: ", service.Id, "class:", service.Class)
    var confs []ProbeConfig
    for _,c := range service.Class{
        classConf,err:= reDao.GetProbeConfigs(c)
        if err !=nil{
            fmt.Println("unable to fetch ProbeConfig for class: "+c)
            fmt.Println(err)
        }else{
            confs = append(confs, classConf...)
        }
    }
    return confs
}


