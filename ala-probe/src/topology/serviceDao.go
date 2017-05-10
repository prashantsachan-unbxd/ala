package topology

import (
    
    )
//Service Dao interface for Accessing service configurations
type ServiceDao  interface{
    //Init initializes the object
    Init()
    //GetAllServices returns list of Services
    GetAllServices()([]Service,error)
    //AddService Adds a service to the existing list
    AddService(s Service)error
    //DeleteService deletes a service from the existing list
    DeleteService(id string) error
}

