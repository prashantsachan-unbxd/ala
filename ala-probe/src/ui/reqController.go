package ui
import (
    mux "github.com/gorilla/mux"
    )
type ReqController interface {
    Register(r *mux.Router)
}
