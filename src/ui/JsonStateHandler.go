package ui
import(
    "encoding/json"
    "result"
    "net/http"
    "api"
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
func statusToApiList(in map[api.Api]api.ApiStatus) map[string][]api.Api{
    out:= make(map[string][]api.Api )
    for a,s := range in{
        out[s.String()]= append(out[s.String()], a)
    }
    return out
}
