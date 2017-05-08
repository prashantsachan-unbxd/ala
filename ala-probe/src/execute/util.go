package execute

import (
        "fmt"
        topo "topology"
        "encoding/json"
    )

func fetchMetrics(reDao RuleEngineDao, services []topo.Service, out chan Event){
    for _,s := range services{
        confs:= fetchProbeConfig(s,reDao);
        fmt.Println(" ProbeConfigs for service: "+s.Id)
        for _,c := range confs{
            str, _:= json.Marshal(c)
            fmt.Println(string(str))
        }
    }
}
// call ruleEngine to get Probe Info list, parse each into PRobeConfig model

        // instantiate one probeClient per ProbeConfig

        // Send request as per ProbeClient

        // call ruleEngine for each of response & each metricFilter

        // create Event for each combination

        // send it to

//Service
    // Id string `json:"id"`
    // Host string `json:"host"`
    // Port int `json:"port"`
    // Class []string `json:"class"`
    // Metadata map[string]interface{} `json:"metadata"`

//ProbeConfig
    // ProbeType string `json:"probeType"`
    // ProbeData map[string]interface{} `json:"probeData"`
    // Metrics []map[string]string `json:"metrics"`

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


