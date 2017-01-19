package nw

import (
    "net/http"
)

type ApiExec interface{
    Fire(req *http.Request) http.Response
}

func GetSimpleClient()  http.Client{
    DefaultClient := http.Client{}
    return  DefaultClient
}
