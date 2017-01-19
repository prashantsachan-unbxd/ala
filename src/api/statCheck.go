package api
import(
    "net/http"
    )

type RespChecker interface{                                                                                        
GetStatus(resp http.Response, err error) ApiStatus                                                             
}
