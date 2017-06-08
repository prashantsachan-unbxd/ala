package execute

import (
        log "github.com/Sirupsen/logrus"
        "sync"
        topo "topology"
        "client"
        "response"
        "time"
        "result"
        "execute/probe"

    )

const FIELD_METRIC_NAME = "metricName"
const FIELD_DEFAULT_VALUE = "defaultMetricValue"
const FIELD_METRICS = "metrics"
//fetchProbeConfigs takes a set of services & retrieves ProbeConfigs for them without repetition
func fetchProbeConfigs(pcDao probe.ProbeConfigDao, services []topo.Service)map[string][]probe.ProbeConfig{
    classes := uniqueClasses(services)
    log.WithFields(log.Fields{"module":"executor","stage":"fetch probeConfig",
    "serviceClasses": classes}).Error("ready to fetch probeConfigs")
    configs:= make(map[string][]probe.ProbeConfig)
    var wg = sync.WaitGroup{}
    for _,cls := range classes{
        wg.Add(1)
        go func(class string){
            defer wg.Done()
            pConfs,err := pcDao.GetAllProbeConfs(class)
            if err !=nil {
                log.WithFields(log.Fields{"module":"executor","stage":"probeConfig",
                 "error":err, "serviceClass": class}).Error(" unable to fetch ProbeConfig")
            }else{
                log.Debug("fetching ProbeConfs for serviceClass"+class)
                configs[class] = pConfs
            }
        }(cls)
    }
    wg.Wait()
    return configs
}
func uniqueClasses(services []topo.Service)[]string{
    classes := make([]string,0)
    for _,s:= range services{
        for _,c := range s.Class{
            classes = append(classes, c)
        }
        
    }
    return unique(classes)
}

func unique(original[]string)[]string{
    m:= make(map[string]bool)
    for _,v := range original{
        if ! m[v]{
            m[v] = true
        }
    }
    keys := make([]string, 0, len(m))
    for v,_ := range m{
        keys = append(keys,v)
    }
    return keys
}

//fetchMetrics computes all the metrics for all the services & send an event for each of them to the out channel
func fetchMetrics(reDao RuleEngineDao, pcDao probe.ProbeConfigDao, services []topo.Service, out chan result.Event){
    timestamp:= time.Now().UTC().UnixNano() / 1000000
    probeConfMap := fetchProbeConfigs(pcDao, services)
    log.WithFields(log.Fields{"module":"executor","stage":"probeConfig", 
        "value":probeConfMap}).Debug("fetched ProbeConfigs")
    //fetch metrics for each service
    for _,serv := range services{
        for _,servClass := range serv.Class{
            confs,ok := probeConfMap[servClass]
            if !ok{
                log.WithFields(log.Fields{"module":"executor","serviceId":serv.Id,
                    "serviceClass":servClass}).Debug("no ProbeConf for serviceClass")
            }else{
                for _,c:= range confs {
                    go func(pc probe.ProbeConfig, s topo.Service, ts int64){
                        //for each probeConfig, create a client & send probeRequest 
                        log.WithFields(log.Fields{"module":"executor", "probeConf":pc}).Debug("intantiating Client")
                        client, cErr:= client.GetClient(pc.ProbeType, pc.ProbeData, s)
                        if(cErr!=nil){
                            log.WithFields(log.Fields{"module":"executor", "serviceId":s.Id,"clientType":pc.ProbeType,
                                "clientData":pc.ProbeData, "error":cErr}).Error("error instantiating client")    
                            // forward the default valued event 
                            collectAndSendMetrics(reDao, nil, pc.Metrics,s, ts,out)
                            return
                        }
                        pResp, pErr := client.Execute()
                        if(pErr !=nil){
                            log.WithFields(log.Fields{"module":"executor", "serviceId":s.Id, 
                                "clientType":pc.ProbeType, "error":pErr}).Error("error in probing")
                            // forward the default valued event
                            pResp = nil
                        }
                        collectAndSendMetrics(reDao, pResp, pc.Metrics,s,ts,out)
                    }(c, serv, timestamp)
                }
            }
        }
    }
}

//collectAndSendMetrics computes set of metrics from a probeResponse & sends an event for each of them to channel
func collectAndSendMetrics(reDao RuleEngineDao, pResp response.ProbeResponse, 
    metricConfs[]map[string]interface{}, service topo.Service, timestamp int64, out chan result.Event){
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
func fetchProbeConfig(service topo.Service, reDao RuleEngineDao)[]probe.ProbeConfig{
    log.WithFields(log.Fields{"module":"executor", "serviceId":service.Id, "class":service.Class}).Debug(
        "fetching probeConfigs for service: ")
    var confs []probe.ProbeConfig
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


