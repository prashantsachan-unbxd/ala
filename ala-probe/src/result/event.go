package result

import (
    topo "topology"
    )

//Event represents computation of a metric for a specific service
//It is sent to kafka for as a message.
type Event struct{
    //Srvc : service for which the execution was done
    Srvc topo.Service `json:"service"`
    //UTC Timestamp at which the execution started (milliseconds since epoch)
    Timestamp int64 `json:"timestamp"`
    //MetricName which is computed
    MetricName string `json:"metricName"`
    //MetricVal : value of the metric
    MetricVal float64 `json:"value"`
}
