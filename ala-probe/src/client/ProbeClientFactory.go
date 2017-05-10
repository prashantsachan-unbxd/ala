package client

import(
    "errors"
    topo "topology"
    )

const CLIENT_TYPE_HTTP ="HTTP"
var typeMap = map[string]ProbeClient{
    CLIENT_TYPE_HTTP: &HttpClient{},
}
//GetClient returns a ProbeClient matching with type
// it initializes & returns a new ProbeClient
//parameters 'data' & 'service' are used for initialization
func GetClient(probeType string, data map[string]interface{}, service topo.Service) (ProbeClient, error){
    dummy, ok:= typeMap[probeType]
    if !ok{
        return nil, errors.New("invalid probe type: "+probeType)
    }
    return dummy.New(data, service)
}