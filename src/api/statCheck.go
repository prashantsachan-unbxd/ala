package api
import(
    "net/http"
    )

type RespCheck interface{                                                                                        
GetStatus(resp http.Response, err error) ApiStatus                                                             
}
