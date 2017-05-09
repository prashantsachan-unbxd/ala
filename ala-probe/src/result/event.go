package result

import (
    "time"
    topo "topology"
    )


type Event struct{
    Srvc topo.Service `json:"service"`
    Timestamp time.Time `json:"timestamp"`
    MetricName string `json:"metricName"`
    MetricVal float64 `json:"value"`
}
