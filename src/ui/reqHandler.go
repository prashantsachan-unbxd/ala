package ui
import (
    mux "github.com/gorilla/mux"
    )
type ReqHandler interface {
    Register(r *mux.Router)
}
