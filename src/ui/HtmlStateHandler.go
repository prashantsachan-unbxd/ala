package ui
import(
    "net/http"
    "result"
    "html/template"
    "fmt"
    mux "github.com/gorilla/mux"
    )

type HtmlStateHandler struct{
    Sm result.StateManager
}
func (h * HtmlStateHandler)Register(r *mux.Router){
    r.HandleFunc("/state/html",h.handleStateHtml)
}
func(h *HtmlStateHandler) handleStateHtml(w http.ResponseWriter, r *http.Request){
        t,err := template.ParseFiles("./resource/state_template.html")
        cnf:=struct{
            Url string
            RefreshTimeMs int    
        }{
            "http://localhost:8080/state",
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
