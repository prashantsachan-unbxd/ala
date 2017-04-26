package ui
import(
    "conf"
    "net/http"
    mux "github.com/gorilla/mux"
    "encoding/json"
    "io/ioutil"
    "github.com/satori/go.uuid"
     )

type ConfController struct{
    Mgr conf.ConfManager
}

func (h *ConfController) Register(r *mux.Router){
    r.Methods("GET").Path("/conf/api").HandlerFunc(h.getConfs)
    r.Methods("PUT").Path("/conf/api").HandlerFunc(h.addConf)
    r.Methods("DELETE").Path("/conf/api/{confId}").HandlerFunc(h.deleteConf)
}   
func (h *ConfController)deleteConf(w http.ResponseWriter, r* http.Request){
    pVars:= mux.Vars(r)
    confId,ok := pVars["confId"] 
    if !ok{
        http.Error(w, "apiConf Id passed is empty", http.StatusBadRequest)
        return
    }else if err:=h.delConf(confId);err !=nil{ 
        http.Error(w, "unable to delete apiConf with id: "+ confId+"\n"+
            err.Error(), http.StatusInternalServerError)
        return
    }else{
        w.Write([]byte("deleted the apiConf with Id: "+confId))
        return
    }
}
func (h *ConfController)delConf(confId string) error{
    id,err:= uuid.FromString(confId)
    if err !=nil{
        return err
    }
    return h.Mgr.DelteConf(id)
}
func (h *ConfController)addConf(w http.ResponseWriter, r *http.Request){
    data, err := ioutil.ReadAll(r.Body)
    if err!= nil{
        http.Error(w, "unable to read data:\n"+err.Error(), http.StatusBadRequest)
        return
    }
    newConf,err := conf.FromJson(string(data))
    if err != nil{
        http.Error(w, "unable to create apiConf from data\n"+ err.Error(), http.StatusBadRequest)
        return
    }
    id, err1 := h.Mgr.AddConf(newConf)
    if err1!=nil{
        http.Error(w,"unable to persist new apiConf\n"+ err1.Error(), http.StatusInternalServerError)
        return
    }
    w.Write([]byte("successfully added new api to tracking list with id:"+id.String()))
}
func (h *ConfController)getConfs(w http.ResponseWriter, r *http.Request){
    tags, ok:= r.URL.Query()["tags"]
    var filtered []conf.ApiConf
    if ok && len(tags)>0{
       filtered = h.Mgr.FilterConfs( tags)
    }else{
        filtered = h.Mgr.GetConfs()
    }
    js,err:= json.Marshal(filtered)
    if err != nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
func filter1(confs []conf.ApiConf, tags []string)[]conf.ApiConf{
    var filtered [] conf.ApiConf
    for _, c := range confs{
        found := false
        for _, t := range c.Tags{
            if(found) {break}
            for _,rt := range tags{
                if(t == rt){
                    found = true
                    filtered  = append(filtered, c)
                    break
                }
            }
        }
    }
    return filtered
}
