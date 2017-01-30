package ui
import(
    "encoding/json"
    "result"
    "net/http"
    )

type JsonStateHandler struct{
    Sm result.StateManager
}
func (h * JsonStateHandler) HandleFunc() func(w http.ResponseWriter, r *http.Request){
    return func(w http.ResponseWriter, r * http.Request){
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
}

