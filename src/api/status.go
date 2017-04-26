package api

import(
    "encoding/json"
    )


type ApiStatus int
const (
    STATUS_RED ApiStatus = iota
//    STATUS_YELLOW 
    STATUS_GREEN 
)
func (s ApiStatus ) String () string{
    switch s{
        case STATUS_GREEN : return "GREEN"
//        case STATUS_YELLOW : return "YELLOW"
        case STATUS_RED : return "RED"
        default: return "UNKNOWN"
    }
}
func (s ApiStatus) MarshalText() ([]byte, error) {
         return []byte(s.String()), nil
}

func (s ApiStatus) MarshalJSON() ([]byte, error) {
         return json.Marshal(s.String())
}
