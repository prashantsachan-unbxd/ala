package ui
import(
    "conf"
    "net/http"
    mux "github.com/gorilla/mux"
    "encoding/json"
     )

type ConfController struct{
    Configs *[]conf.ApiConf
    ConfWriter conf.ConfWriter
}

func (h *ConfController) Register(r *mux.Router){
    r.Methods("GET").Path("/conf/api").HandlerFunc(h.getConfs)
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
