
package ui
import(
    "net/http"
    mux "github.com/gorilla/mux"
    "encoding/json"
    "io/ioutil"
    topo "topology"
)
type ServiceController struct{
    ServiceDao topo.ServiceDao
}

func (this *ServiceController) Register(r *mux.Router){
    r.Methods("GET").Path("/service/{id}").HandlerFunc(this.getService)
    r.Methods("GET").Path("/service").HandlerFunc(this.getAllServices)
    r.Methods("PUT").Path("/service").HandlerFunc(this.addService)
    r.Methods("DELETE").Path("/service/{id}").HandlerFunc(this.deleteService)
}   
func (this *ServiceController)deleteService(w http.ResponseWriter, r* http.Request){
    pVars:= mux.Vars(r)
    id,ok := pVars["id"] 
    if !ok{
        http.Error(w, "service Id passed is empty", http.StatusBadRequest)
        return
    }else if err:=this.ServiceDao.DeleteService(id);err !=nil{ 
        http.Error(w, "unable to delete service with id: "+ id+"\n"+
            err.Error(), http.StatusInternalServerError)
        return
    }else{
        w.Write([]byte("deleted the service with Id: "+id))
        return
    }
}
func (this *ServiceController)addService(w http.ResponseWriter, r *http.Request){
    data, err := ioutil.ReadAll(r.Body)
    if err!= nil{
        http.Error(w, "unable to read data:\n"+err.Error(), http.StatusBadRequest)
        return
    }
    var s topo.Service 
    jErr := json.Unmarshal(data, &s)
    if jErr != nil{
        http.Error(w, "unable to parse service from data\n"+ jErr.Error(), http.StatusBadRequest)
        return
    }else if !s.IsValid(){
    	http.Error(w, "insufficient service details\n", http.StatusBadRequest)
        return
    }else{
    	daoErr:= this.ServiceDao.AddService(s)
    	if daoErr !=nil{
    		http.Error(w,"unable to persist new service\n"+ daoErr.Error(), http.StatusInternalServerError)
        	return	
    	}else{
    		w.Write([]byte("successfully added new service"))
    	}
    }
}
func (this *ServiceController)getService(w http.ResponseWriter, r *http.Request){
	pVars:= mux.Vars(r)
    id,ok := pVars["id"]
    if !ok{
        http.Error(w, "service Id passed is empty", http.StatusBadRequest)
        return
    }
    s,daoErr := this.ServiceDao.GetService(id)
    if daoErr != nil{
        http.Error(w, "unable to fetch Service with give id\n"+daoErr.Error(), http.StatusInternalServerError)
        return
    }else{
    	js,jErr:= json.Marshal(s)
    	if jErr !=nil{
    		http.Error(w,"error writing service to output\n"+ jErr.Error(), 
    			http.StatusInternalServerError)
    		return
    	}else{
    		w.Header().Set("Content-Type", "application/json")
    		w.Write(js)			
    	}
    }
}
func (this *ServiceController) getAllServices(w http.ResponseWriter, r *http.Request){
	services,daoErr:= this.ServiceDao.GetAllServices()
	if daoErr !=nil{
		http.Error(w,"error fetching services\n"+ daoErr.Error(), 
    		http.StatusInternalServerError)
		return
	}else{
		js,jErr:= json.Marshal(services)
    	if jErr !=nil{
    		http.Error(w,"error writing services to output\n"+ jErr.Error(), 
    			http.StatusInternalServerError)
    		return
    	}else{
    		w.Header().Set("Content-Type", "application/json")
    		w.Write(js)			
    	}	
	}
}
