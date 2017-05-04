package topology

import (
    
    )

type ServiceDao  interface{
    GetAllServices()([]Service,error)
    AddService(s Service)error
    DeleteService(id string) error
}

