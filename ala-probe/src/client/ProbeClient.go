//Package client provides Implementations for different service protocols
package client
import (
    resp "response"
    topo "topology"
    )
//ProbeClient interface for all service protocols
//ProbeClient implementation contains the logic to connect and
// retrieve information from a service. 
//examples of ProbeClient implementation could be HTTP, mongo, zookeeper etc.
type ProbeClient interface{
    //isEmpty checks whether the instance is empty or not (equivalent to null check)
    isEmpty()bool
    // New returns a new Instance of this type, initialized with the supplied config
    New(config map[string]interface{}, service topo.Service) (ProbeClient,error)
    // Execute runs the probe request as per config & return a ProbeResponse
    Execute()(resp.ProbeResponse, error)
}