package result
import (
    "ex"
    "api"
    "time"
    "fmt"
    )

var lastCleanupTime time.Time
var latestStatus map[api.Api]api.ApiStatus
var latestTime map[api.Api]time.Time

type TimedStateManager struct{
    Timeout time.Duration
    CleanDur time.Duration
}
func init(){
    lastCleanupTime = time.Now()
    latestStatus = make(map[api.Api]api.ApiStatus)
    latestTime = make(map[api.Api]time.Time)
}
func (sm *TimedStateManager)cleanIfReq(){
    cleanNeeded :=lastCleanupTime.Add(sm.CleanDur).Before(time.Now())
    if(!cleanNeeded){
        return
    }
    threshTime := time.Now().Add(- sm.CleanDur)
    for a,t := range latestTime {
        if t.Before(threshTime){
            delete(latestTime,a )
            delete(latestStatus, a)
            fmt.Println("deleting api ", a, " because of timeOut")
        }
    }
    lastCleanupTime = time.Now()
}
func (sm *TimedStateManager)UpdateState(e ex.Event){
    t,ok:= latestTime[e.Api]
    if !ok || t.Before(e.Timestamp) {
        latestTime[e.Api] = e.Timestamp
        latestStatus[e.Api]= e.Status
    }
}
func(sm *TimedStateManager)GetState()map[api.Api]api.ApiStatus{
    sm.cleanIfReq()
    thresh := time.Now().Add(-sm.Timeout)
    state:= make(map[api.Api]api.ApiStatus)
    for a,t:= range latestTime{
       if t.After(thresh) && latestStatus[a]== api.STATUS_GREEN{
            state[a] = api.STATUS_GREEN
        }else{
            state[a] = api.STATUS_RED
        }
    }
    return state
}

