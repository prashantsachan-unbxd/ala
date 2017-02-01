package ui
import(
    "encoding/json"
    "result"
    "net/http"
    "html/template"
    "fmt"
    mux "github.com/gorilla/mux"
    )

type StateController struct{
    Sm result.StateManager
}
func (h * StateController) Register (r *mux.Router){
    r.HandleFunc("/state/json", h.stateJson) 
    r.HandleFunc("/state/html", h.stateHtml)
}
func (h *StateController)stateJson(w http.ResponseWriter, r *http.Request){
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

func(h *StateController) stateHtml(w http.ResponseWriter, r *http.Request){
        t,err := template.ParseFiles("./resource/state_template.html")
        cnf:=struct{
            Url string
            RefreshTimeMs int    
        }{
            "http://localhost:8080/state/json",
            3000,
        }
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            fmt.Println("error in parsing template file for API state")
            return
        }
        err = t.Execute(w, cnf)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            fmt.Println("error in rendering template for apiState")
        }
}
