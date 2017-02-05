package ui
import(
    "conf"
    "net/http"
    mux "github.com/gorilla/mux"
    "encoding/json"
    "io/ioutil"
     )

type ConfController struct{
    Configs *[]conf.ApiConf
    ConfWriter conf.ConfWriter
}

func (h *ConfController) Register(r *mux.Router){
    r.Methods("GET").Path("/conf/api").HandlerFunc(h.getConfs)
    r.Methods("PUT").Path("/conf/api").HandlerFunc(h.addConf)
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
    newConfigSet := append(*h.Configs, newConf)
    err= h.ConfWriter.WriteApiConf(newConfigSet)
    if err!=nil{
        http.Error(w,"unable to persist new apiConf\n"+ err.Error(), http.StatusInternalServerError)
        return
    }
    *h.Configs = newConfigSet
    w.Write([]byte("successfully added new api to tracking list"))
}
func (h *ConfController)getConfs(w http.ResponseWriter, r *http.Request){
    tags, ok:= r.URL.Query()["tags"]
    filtered  := *h.Configs
    if ok && len(filtered)>0{
       filtered = filter(filtered, tags)
    }
    js,err:= json.Marshal(filtered)
    if err != nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
func filter(confs []conf.ApiConf, tags []string)[]conf.ApiConf{
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
