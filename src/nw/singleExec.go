package nw
import(
    "net/http"
    )
type SimpleApiExec struct{
    Client http.Client
}

func (e SimpleApiExec) Fire(req *http.Request) (*http.Response, error){
    return e.Client.Do(req)
}
