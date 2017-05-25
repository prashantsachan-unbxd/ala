package ui
import (
	"github.com/gorilla/mux"
)

type ReqController interface{
	Register(r *mux.Router)
}