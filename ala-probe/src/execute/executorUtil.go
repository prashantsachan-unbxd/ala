package execute

import (
        log "github.com/Sirupsen/logrus"
        topo "topology"
        "client"
        "response"
        "time"
        "result"
    )

const FIELD_METRIC_NAME = "metricName"
const FIELD_DEFAULT_VALUE = "defaultMetricValue"
const FIELD_METRICS = "metrics"
//fetchMetrics computes all the metrics for all the services & send an event for each of them to the out channel
func fetchMetrics(reDao RuleEngineDao, services []topo.Service, out chan result.Event){
    //fetch metrics for each service
    for _,s := range services{
        // get the probeConfigs from RuleEngine
        confs:= fetchProbeConfig(s,reDao)
        log.WithFields(log.Fields{"module":"executor","serviceId":s.Id,"stage":"probeConfig"}).Debug(
            "fetched ProbeConfig")
        for _,c := range confs{
            //for each probeConfig, create a client & send probeRequest 
            log.WithFields(log.Fields{"module":"executor", "probeConfig":c}).Debug()
            client, cErr:= client.GetClient(c.ProbeType, c.ProbeData, s)
            if(cErr!=nil){
                log.WithFields(log.Fields{"module":"executor", "serviceId":s.Id,"clientType":c.ProbeType,
                    "clientData":c.ProbeData, "error":cErr}).Error("error instantiating client")    
                // forward the default valued event 
                collectAndSendMetrics(reDao, nil, c.Metrics,s,out)
                return
            }
            pResp, pErr := client.Execute()
            if(pErr !=nil){
                log.WithFields(log.Fields{"module":"executor", "serviceId":s.Id, 
                    "clientType":c.ProbeType, "error":pErr}).Error("error in probing")
                // forward the default valued event
                pResp = nil
            }
            collectAndSendMetrics(reDao, pResp, c.Metrics,s,out)
        }
    }
}

//collectAndSendMetrics computes set of metrics from a probeResponse & sends an event for each of them to channel
func collectAndSendMetrics(reDao RuleEngineDao, pResp response.ProbeResponse, 
    metricConfs[]map[string]interface{}, service topo.Service, out chan result.Event){
    timestamp:= time.Now()
    metrics := getMetricValues(reDao, pResp, metricConfs)
    for k,v := range metrics{
        log.WithFields(log.Fields{"module":"executor", "serviceId":service.Id, "metric":k,"value":v}).Debug(
            )
        out <- result.Event{service,timestamp, k,v.(float64)}
    }
}
//getMetricValues interacts with RuleEngineDao & computes metrics for a ProbeResponse
func getMetricValues(reDao RuleEngineDao, resp response.ProbeResponse, metrics[]map[string]interface{})     (map[string]interface{}){
    vals := make(map[string]interface{})
    //send default values if ProbeRespose is nil
    if resp == nil {
        for _,m1 := range metrics{
            vals[m1[FIELD_METRIC_NAME].(string)] = m1[FIELD_DEFAULT_VALUE]
        }
    }else{
        respMap:= resp.AsMap()
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
            val, reErr:= reDao.GetMetricVal(respMap, segment, defaultVal)
            if reErr !=nil {
                log.WithFields(log.Fields{"module":"executor", "metric":mName, "error": reErr}).Error(
                    "error retrieving metric from ProbeResponse")
                vals[mName] = defaultVal
            }else{
                vals[mName] = val
            }
        }
    }
    return vals
}
//getMetricValues interacts with RuleEngine & return list of Probeconfigs for a particular service
func fetchProbeConfig(service topo.Service, reDao RuleEngineDao)[]ProbeConfig{
    log.WithFields(log.Fields{"module":"executor", "serviceId":service.Id, "class":service.Class}).Debug(
        "fetching probeConfigs for service: ")
    var confs []ProbeConfig
    // send one request for each of the serviceClass
    for _,c := range service.Class{
        classConf,err:= reDao.GetProbeConfigs(c)
        if err !=nil{
            log.WithFields(log.Fields{"module":"executor", "class": service.Class, 
                "error":err}).Error("error fetching ProbeConfig")
        }else{
            confs = append(confs, classConf...)
        }
    }
    return confs
}


