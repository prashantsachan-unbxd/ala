package ui
import(
    "encoding/json"
    "result"
    "net/http"
    "fmt"
    )

type JsonStateHandler struct{
    Sm result.StateManager
}
func (h * JsonStateHandler) HandleFunc() func(w http.ResponseWriter, r *http.Request){
    return func(w http.ResponseWriter, r * http.Request){
        state:= h.Sm.GetState()
        stateStr:=make(map[string]interface{})
        for k,v:= range state{
            kStr,err:= json.Marshal(k)
            if err ==nil{
                stateStr[string(kStr)] = v
            }else{
                fmt.Println("can't convert",k," to json for StateHandler")
            }
        }
        js, err := json.Marshal(stateStr)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(js)
    }
}
