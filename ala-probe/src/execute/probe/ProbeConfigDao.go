package probe

type ProbeConfigDao interface{
	//Init initializes the object
    Init()
    //AddService Adds a probeConfig to the existing list
    AddProbeConf(s ProbeConfig, serviceClass string)error
    
    //DeleteProbeConf deletes a probeConfig
    DeleteProbeConf(serviceClass string, id string) error

    //GetAllClasses returns list of all serviceClasses
    GetAllClasses()([]string, error)
    
    //GetAllServices returns list of Services
    GetAllProbeConfs(serviceClass string)([]ProbeConfig,error)

    // Retrieves a service by id
    // should return non-nil error if id passed doesn't match with any service
   	GetProbeConf(serviceClass, id string)(ProbeConfig,error)
}