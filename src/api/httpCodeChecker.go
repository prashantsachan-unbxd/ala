package api
import(
    "fmt"
    "net/http"
)

type HttpCodeChecker struct{}

func (c HttpCodeChecker) GetStatus (resp http.Response, err error) ApiStatus{
    if err !=nil{
        fmt.Println(err)
        return STATUS_RED
    }else if resp.StatusCode >= 200 && resp.StatusCode<300{
        return STATUS_GREEN
    }else{
        return STATUS_RED
    }
    
}
