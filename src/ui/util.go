package ui
import(
    "api"
    )


func statusToApiList(in map[api.Api]api.ApiStatus) map[string][]api.Api{
    out:= make(map[string][]api.Api )
    for a,s := range in{
        out[s.String()]= append(out[s.String()], a)
    }
    return out
}
