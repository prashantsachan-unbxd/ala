package topology

//Service is a model for a Service Configuration
// This contains information regarding a service
type Service struct{
    // Id of service
    Id string `json:"id"`
    //Host at which the service is running
    Host string `json:"host"`
    //Port to connect to the server at
    Port int `json:"port"`
    //Class : serviceClass Names to which a service belongs
    //This value will be used to resolve ProbeConfigs for this service
    //a service can belong to multiple classes, each class defining one or more
    // metrics to be collected for it
    Class []string `json:"class"`
    //Metadata : other properties for this service, which are specific to service type
    //These could be used by the specific ProbeClients while sending probe requests to it
    Metadata map[string]interface{} `json:"metadata"`
}


