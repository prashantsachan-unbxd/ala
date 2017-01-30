package ui
import(
    "net/http"
    "result"
    "html/template"
    "fmt"
    )

type HtmlStateHandler struct{
    Sm result.StateManager
}

func(h *HtmlStateHandler) HandleFunc()func(w http.ResponseWriter, r *http.Request){
    return func(w http.ResponseWriter, r *http.Request){

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
}
