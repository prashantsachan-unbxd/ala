package ui
import (
    "net/http"
    )
type ReqHandler interface {
    HandleFunc() func(w http.ResponseWriter, r *http.Request)
}
