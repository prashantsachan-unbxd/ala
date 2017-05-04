package client

import(
    "errors"
    topo "topology"
    )

const CLIENT_TYPE_HTTP ="HTTP"
var typeMap = map[string]ProbeClient{
    CLIENT_TYPE_HTTP: &HttpClient{},
}
type ProbeClientFactory struct{
       
}
func GetClient(valType string, data map[string]interface{}, service topo.Service) (ProbeClient, error){
    dummy, ok:= typeMap[valType]
    if !ok{
        return nil, errors.New("invalid validator type: "+valType)
    }
    return dummy.New(data, service)
}