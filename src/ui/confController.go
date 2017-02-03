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
    r.Methods("GET").Path("/conf/api").HandlerFunc(h.getAllApis)
}   

func (h *ConfController)getAllApis(w http.ResponseWriter, r *http.Request){
    js,err:= json.Marshal(h.Configs)
    if err!=nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
