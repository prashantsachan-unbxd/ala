package topology

import (
    
    )

type ServiceDao  interface{
    Init()
    GetAllServices()([]Service,error)
    AddService(s Service)error
    DeleteService(id string) error
}

