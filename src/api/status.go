package api

import(
    )


type ApiStatus int
const (
    STATUS_GREEN ApiStatus = iota
    STATUS_YELLOW 
    STATUS_RED 
)
func (s ApiStatus ) String () string{
    switch s{
        case STATUS_GREEN : return "GREEN"
        case STATUS_YELLOW : return "YELLOW"
        case STATUS_RED : return "RED"
        default: return "UNKNOWN"
    }
}
