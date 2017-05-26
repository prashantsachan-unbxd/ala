package probe

import (
    //"encoding/json"
    )
//ProbeConfig has configuration related to Probing & collecting metric for a service
// it basically 'answers how to probe a service?'
type ProbeConfig struct{
    // unique Id to identify
    Id string       `json:"id"`
    //ProbeType tells which ProbeClient to use
    //ProbeClientFactory decides the probeClient by matching this
    ProbeType string `json:"probeType"`
    //ProbeData contains the configuration of the HttpClient of specified type
    //This would be passed to the ProbeClient of specified ProbeType at the time of initialization
    //hence, the entries in it depend upon the type of ProbeClient defined by ProbeType
    ProbeData map[string]interface{} `json:"probeData"`
    //Metrics lists out the metrics to be computed after ProbeClient's executes
    //executorUtil expects each of these metrics(maps) to contain following keys: 
    //      defaultMetricValue (number) => defult value to be passed in case of any error
    //      domain             (string) => used for rule resolution []
    //      subdomain          (string) => used for rule resoultion
    //      metricName         (string) => used for rule resolution
    Metrics []map[string]interface{} `json:"metrics"`

}
//an example of ProbeConfig is as follows: 
// {
//      "id":"http-tomcat"
//      "probeType":"HTTP",
//      "probeData":{
//          "connTimeout":"3",
//          "method":"GET",
//          "path":"/tomcat.gif",
//          "readTimeout":"5"
//      },
//      "metrics":[
//          {
//              "defaultMetricValue":0,
//              "domain":"platform.monitoring",
//              "metricName":"HTTPstatus200",
//              "subdomain":"metricCollect"
//          }
//      ]
// }
