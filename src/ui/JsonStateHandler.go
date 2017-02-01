package ui
import(
    "encoding/json"
    "result"
    "net/http"
    mux "github.com/gorilla/mux"
    )

type JsonStateHandler struct{
    Sm result.StateManager
}
func (h * JsonStateHandler) Register (r *mux.Router){
    r.HandleFunc("/state", h.handleStateJson) 
}
func (h *JsonStateHandler)handleStateJson(w http.ResponseWriter, r *http.Request){
        state:= h.Sm.GetState()
        stateJson :=statusToApiList(state)
        js, err := json.Marshal(stateJson)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Write(js)
    }

