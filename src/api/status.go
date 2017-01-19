package api

import(
    )


type ApiStatus int
const (
    GREEN ApiStatus = iota
    YELLOW ApiStatus = iota
    RED ApiStatus = iota
)
func (s ApiStatus ) String () string{
    switch s{
        case GREEN : return "GREEN"
        case YELLOW : return "YELLOW"
        case RED : return "RED"
        default: return "UNKNOWN"
    }
}
