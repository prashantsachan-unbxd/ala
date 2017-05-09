package client
import (
    resp "response"
    topo "topology"
    )
type ProbeClient interface{
    isEmpty()bool
    // Returns a new Instance of this type using the supplied configuration
    New(config map[string]interface{}, service topo.Service) (ProbeClient,error)
    // Executes the probe request as per config & return a result
    Execute()(resp.ProbeResponse, error)
}