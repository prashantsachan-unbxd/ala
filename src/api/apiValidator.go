package api
import(
    "net/http"
    )

type ApiValidator interface{                                                                                        
GetStatus(resp http.Response, err error) ApiStatus                                                             
}
