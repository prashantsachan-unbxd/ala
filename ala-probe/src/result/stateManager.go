package result 
import(
    "ex"
    "api"
    )

type StateManager interface {
    UpdateState(e ex.Event)
    GetState()map[api.Api]api.ApiStatus
}
