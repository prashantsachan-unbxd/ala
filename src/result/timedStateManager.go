package result
import (
    "ex"
    "api"
    "time"
    )

type TimedStateManager struct{
    Timeout time.Duration
    GreenTime map[api.Api]time.Time
}

func(sm *TimedStateManager) UpdateState(e ex.Event){
    var tnew time.Time
    if e.Status == api.STATUS_GREEN{
        tnew = e.Timestamp
    }else{
        tnew = e.Timestamp.Add(-sm.Timeout)
    }
    last,ok := sm.GreenTime[e.Api]
    if !ok{
        sm.GreenTime[e.Api] = tnew
    }else if last.Before(tnew){
        sm.GreenTime[e.Api]= tnew
    }
}
func(sm *TimedStateManager)GetState()map[api.Api]api.ApiStatus{
    thresh:= time.Now().Add(-sm.Timeout)
    state:= make(map[api.Api]api.ApiStatus)
    for a,t:= range sm.GreenTime{
        if t.After(thresh){
            state[a]= api.STATUS_GREEN
        }else{
            state[a] = api.STATUS_RED
        }
    }
    return state
}
