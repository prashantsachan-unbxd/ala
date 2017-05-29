package ui

import(
    "net/http"
    mux "github.com/gorilla/mux"
    "encoding/json"
    "io/ioutil"
    "execute/probe"
)


type ProbeConfController struct{
    ProbeConfDao probe.ProbeConfigDao
}

func (this *ProbeConfController) Register(r *mux.Router){
    r.Methods("GET").Path("/probeconfig/class").HandlerFunc(this.getServiceClasses)
    r.Methods("GET").Path("/probeconfig/{serviceClass}/{id}").HandlerFunc(this.getProbeConf)
    r.Methods("GET").Path("/probeconfig/{serviceClass}").HandlerFunc(this.getProbeConfsForService)
    r.Methods("PUT").Path("/probeconfig/{serviceClass}").HandlerFunc(this.addProbeConf)
    r.Methods("DELETE").Path("/probeconfig/{serviceClass}/{id}").HandlerFunc(this.deleteProbeConf)
}   
func (this *ProbeConfController)deleteProbeConf(w http.ResponseWriter, r* http.Request){
    pVars:= mux.Vars(r)
    id,ok1 := pVars["id"] 
    serviceClass,ok2 := pVars["serviceClass"]
    if !ok1{
        http.Error(w, "probeConfgId passed is empty", http.StatusBadRequest)
        return
    }else if !ok2{
        http.Error(w, "serviceClass passed is empty", http.StatusBadRequest)
        return
    }else if err:=this.ProbeConfDao.DeleteProbeConf(serviceClass,id);err !=nil{ 
        http.Error(w, "unable to delete probeConf with id: "+ id+"\n"+
            err.Error(), http.StatusInternalServerError)
        return
    }else{
        w.Write([]byte("deleted the probeConfig with Id: "+id))
        return
    }
}
func (this *ProbeConfController)addProbeConf(w http.ResponseWriter, r *http.Request){
	pVars:= mux.Vars(r) 
    serviceClass,ok := pVars["serviceClass"]
    if !ok{
        http.Error(w, "serviceClass passed is empty", http.StatusBadRequest)
        return
    }
    data, err := ioutil.ReadAll(r.Body)
    if err!= nil{
        http.Error(w, "unable to read data:\n"+err.Error(), http.StatusBadRequest)
        return
    }
    var p probe.ProbeConfig
    jErr := json.Unmarshal(data, &p)
    if jErr != nil{
        http.Error(w, "unable to parse probeConf from data\n"+ jErr.Error(), http.StatusBadRequest)
        return
    }else if !p.IsValid(){
    	http.Error(w, "insufficient ProbeConfig details\n", http.StatusBadRequest)
        return
    }else{
    	daoErr:= this.ProbeConfDao.AddProbeConf(p,serviceClass)
    	if daoErr !=nil{
    		http.Error(w,"unable to persist new service\n"+ daoErr.Error(), http.StatusInternalServerError)
        	return	
    	}else{
    		w.Write([]byte("successfully added new service"))
    	}
    }
}
func (this *ProbeConfController)getProbeConf(w http.ResponseWriter, r *http.Request){
	pVars:= mux.Vars(r)
    id,ok1 := pVars["id"]
    serviceClass,ok2 := pVars["serviceClass"]
    if !ok1{
        http.Error(w, "probeConf Id passed is empty", http.StatusBadRequest)
        return
    }else if !ok2{
    	http.Error(w, "serviceClass passed is empty", http.StatusBadRequest)
        return
    }
    p,daoErr := this.ProbeConfDao.GetProbeConf(serviceClass,id)
    if daoErr != nil{
        http.Error(w, "unable to fetch probeConf with give id\n"+daoErr.Error(), http.StatusInternalServerError)
        return
    }else{
    	js,jErr:= json.Marshal(p)
    	if jErr !=nil{
    		http.Error(w,"error writing probeConfservice to output\n"+ jErr.Error(), 
    			http.StatusInternalServerError)
    		return
    	}else{
    		w.Header().Set("Content-Type", "application/json")
    		w.Write(js)			
    	}
    }
}
func (this *ProbeConfController) getProbeConfsForService(w http.ResponseWriter, r *http.Request){
	pVars:= mux.Vars(r)
    serviceClass,ok := pVars["serviceClass"]
    if !ok{
        http.Error(w, "serviceClass passed is empty", http.StatusBadRequest)
        return
    }
	probeConfs,daoErr:= this.ProbeConfDao.GetAllProbeConfs(serviceClass)
	if daoErr !=nil{
		http.Error(w,"error fetching probeConfs for class"+serviceClass+"\n"+ daoErr.Error(), 
    		http.StatusInternalServerError)
		return
	}else{
		js,jErr:= json.Marshal(probeConfs)
    	if jErr !=nil{
    		http.Error(w,"error writing probeConfs to output\n"+ jErr.Error(), 
    			http.StatusInternalServerError)
    		return
    	}else{
    		w.Header().Set("Content-Type", "application/json")
    		w.Write(js)			
    	}	
	}
}

func (this * ProbeConfController)getServiceClasses(w http.ResponseWriter, r *http.Request){
	classes, daoErr:= this.ProbeConfDao.GetAllClasses()
	if daoErr !=nil{
		http.Error(w,"error fetching list of classes\n"+ daoErr.Error(), 
    		http.StatusInternalServerError)
		return
	}else{
		js,jErr:= json.Marshal(classes)
    	if jErr !=nil{
    		http.Error(w,"error writing classNames to output\n"+ jErr.Error(), 
    			http.StatusInternalServerError)
    		return
    	}else{
    		w.Header().Set("Content-Type", "application/json")
    		w.Write(js)			
    	}	
	}
}

